package legalparty

import (
	"strconv"
	"strings"

	errorfamily "github.com/larsartmann/go-error-family"
)

// LEI is a Legal Entity Identifier (ISO 17442): a 20-character,
// globally unique identifier for legal entities that participate in
// financial transactions, issued under GLEIF governance. The final two
// characters are the mod-97 check digits.
//
// Natural persons cannot hold a LEI (GLEIF's scope covers legal entities
// only); [LegalFormClass.AllowsLEI] enforces that at the profile level.
type LEI string

// ParseLEI validates and constructs a LEI. Input is uppercased and
// trimmed, must be exactly 20 characters of [A-Z0-9], and must pass the
// ISO 17442 mod-97 checksum.
func ParseLEI(s string) (LEI, error) {
	v := strings.ToUpper(strings.TrimSpace(s))
	if len(v) != 20 {
		return "", errorfamily.NewRejection("legalparty.lei.invalid",
			"a LEI is exactly 20 characters, got "+strconv.Itoa(len(v))+" in "+s)
	}

	if !isAllUpperAlphaNumeric(v) {
		return "", errorfamily.NewRejection("legalparty.lei.invalid",
			"a LEI carries only digits and uppercase letters, got "+s)
	}

	if !validLEIChecksum(v) {
		return "", errorfamily.NewRejection("legalparty.lei.checksum_failed",
			"LEI "+v+" fails the ISO 17442 mod-97 checksum")
	}

	return LEI(v), nil
}

// validLEIChecksum applies the ISO 17442 check: map every character to its
// numeric form (digits as-is, A=10 .. Z=35), interpret the whole 20-character
// string as one decimal number, and require number mod 97 == 1.
func validLEIChecksum(code string) bool {
	remainder := 0

	for _, r := range code {
		switch {
		case r >= '0' && r <= '9':
			remainder = (remainder*10 + int(r-'0')) % 97
		case r >= 'A' && r <= 'Z':
			v := int(r-'A') + 10
			remainder = (remainder*10 + v/10) % 97
			remainder = (remainder*10 + v%10) % 97
		default:
			return false
		}
	}

	return remainder == 1
}

// RegisterRef points at an entry in a commercial or public register
// (Handelsregister HRB, Companies House number, RCS number, ...). Registry
// is an opaque registry/jurisdiction code owned by whoever resolves the
// reference (an OpenCorporates jurisdiction code such as "de_munich", a
// national register key) — the domain never inspects its semantics.
type RegisterRef struct {
	Registry string `json:"registry"`
	Number   string `json:"number"`
}

// NewRegisterRef validates and constructs a RegisterRef.
func NewRegisterRef(registry, number string) (RegisterRef, error) {
	r := RegisterRef{Registry: strings.TrimSpace(registry), Number: strings.TrimSpace(number)}
	if err := r.Validate(); err != nil {
		return RegisterRef{}, err
	}

	return r, nil
}

// Validate enforces that both the registry code and the number are present.
func (r RegisterRef) Validate() error {
	if strings.TrimSpace(r.Registry) == "" {
		return errorfamily.NewRejection("legalparty.register.registry_required",
			"a register reference names its registry (e.g. the jurisdiction code)")
	}

	if strings.TrimSpace(r.Number) == "" {
		return errorfamily.NewRejection("legalparty.register.number_required",
			"a register reference carries the registry entry number")
	}

	return nil
}

// normalizeTaxIdentifier uppercases and strips the separator characters
// national tax forms carry (spaces, slashes, dashes, dots).
func normalizeTaxIdentifier(s string) string {
	v := strings.ToUpper(strings.TrimSpace(s))

	var b strings.Builder
	b.Grow(len(v))

	for _, r := range v {
		if r == ' ' || r == '/' || r == '-' || r == '.' {
			continue
		}

		b.WriteRune(r)
	}

	return b.String()
}

// isAllDigits reports whether every rune in s is an ASCII digit. The empty
// string is not all digits.
func isAllDigits(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}

	return true
}

// isAllUpperAlpha reports whether every rune in s is an uppercase ASCII
// letter. The empty string is not all letters.
func isAllUpperAlpha(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if r < 'A' || r > 'Z' {
			return false
		}
	}

	return true
}

// isAllUpperAlphaNumeric reports whether every rune in s is an uppercase
// ASCII letter or a digit. The empty string is not a match.
func isAllUpperAlphaNumeric(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if (r < '0' || r > '9') && (r < 'A' || r > 'Z') {
			return false
		}
	}

	return true
}

// VATID is a value-added tax identifier in its canonical EU-prefixed form
// ("DE136695976", "PL1234567890"). It is deliberately a different type from
// [TaxNumber]: a VAT ID starts with a country prefix, a national tax number
// never does, and conflating the two rejects valid values on both sides.
//
// Germany validates the full ISO 7064 Mod 11,10 checksum; the other known
// forms are format-validated. Unprefixed domestic forms (e.g. a bare NIP)
// must be normalized to the prefixed form by the caller.
type VATID string

// ParseVATID validates and constructs a VATID in canonical prefixed form.
func ParseVATID(s string) (VATID, error) {
	v := normalizeTaxIdentifier(s)
	if len(v) < 3 {
		return "", errorfamily.NewRejection("legalparty.vat_id.required",
			"a VAT ID carries a country prefix and a national part, got "+s)
	}

	cc, err := ParseCountryCode(v[:2])
	if err != nil {
		return "", errorfamily.NewRejection("legalparty.vat_id.prefix_invalid",
			"a VAT ID starts with a country prefix: "+err.Error())
	}

	national := v[2:]
	if cc == "AT" && len(national) > 0 && national[0] == 'U' {
		national = national[1:]
	}

	if err := validateVATNationalPart(cc, national); err != nil {
		return "", err
	}

	return VATID(v), nil
}

func validateVATNationalPart(cc CountryCode, national string) error {
	if err := validateVATCharacters(cc, national); err != nil {
		return err
	}

	return validateVATLength(cc, national)
}

// validateVATCharacters checks the national part's alphabet: digits only,
// except the Dutch form which carries the literal B separator.
func validateVATCharacters(cc CountryCode, national string) error {
	for _, r := range national {
		switch {
		case r >= '0' && r <= '9':
			continue
		case r >= 'A' && r <= 'Z':
			if cc != "NL" {
				return errorfamily.NewRejection("legalparty.vat_id.national_invalid",
					"the national part of a "+string(cc)+" VAT ID carries only digits")
			}
		default:
			return errorfamily.NewRejection("legalparty.vat_id.national_invalid",
				"the national part of a VAT ID carries only uppercase letters and digits")
		}
	}

	return nil
}

// validateVATLength enforces the per-jurisdiction national-part shape.
func validateVATLength(cc CountryCode, national string) error {
	switch cc {
	case "DE":
		return validateDEVAT(national)
	case "NL":
		return validateNLVAT(national)
	case "AT", "PL":
		return validateFixedLengthVAT(cc, national)
	default:
		return validateGenericVAT(national)
	}
}

func validateDEVAT(national string) error {
	if len(national) != 9 {
		return errorfamily.NewRejection("legalparty.vat_id.national_invalid",
			"a German USt-IdNr. is DE followed by exactly 9 digits")
	}

	if !validDEVATChecksum(national) {
		return errorfamily.NewRejection("legalparty.vat_id.checksum_failed",
			"USt-IdNr. "+national+" fails the ISO 7064 Mod 11,10 checksum")
	}

	return nil
}

func validateNLVAT(national string) error {
	if len(national) != 12 || national[9] != 'B' || !isAllDigits(national[:9]) || !isAllDigits(national[10:]) {
		return errorfamily.NewRejection("legalparty.vat_id.national_invalid",
			"a Dutch BTW-nummer is NL followed by 9 digits, B, and 2 digits")
	}

	return nil
}

// validateFixedLengthVAT checks the single-length national forms
// (AT: 8 digits, PL: 10 digits).
func validateFixedLengthVAT(cc CountryCode, national string) error {
	var want int

	switch cc {
	case "AT":
		want = 8
	case "PL":
		want = 10
	}

	if len(national) != want {
		return errorfamily.NewRejection("legalparty.vat_id.national_invalid",
			"the national part of a "+string(cc)+" VAT ID is "+strconv.Itoa(want)+" digits")
	}

	return nil
}

// validateGenericVAT checks jurisdictions without a dedicated rule.
func validateGenericVAT(national string) error {
	if len(national) < 8 || len(national) > 12 {
		return errorfamily.NewRejection("legalparty.vat_id.national_invalid",
			"the national part of this VAT ID is 8 to 12 characters")
	}

	return nil
}

// validDEVATChecksum applies the ISO 7064 Mod 11,10 checksum to the full
// nine-digit German national part. Faithful port of the published reference
// implementation (python-stdnum, stdnum/iso7064/mod_11_10.py): the checksum
// runs over ALL digits with a seed of 5 and is valid iff it equals 1.
// Verified vectors: BMF/BZSt example DE136695976 validates, DE136695978
// fails, calc('79462') == '3', calc('00200667308') == '5'.
func validDEVATChecksum(national string) bool {
	check := 5

	for i := range len(national) {
		d := int(national[i] - '0')

		base := check
		if base == 0 {
			base = 10
		}

		check = ((base*2)%11 + d) % 10
	}

	return check == 1
}
