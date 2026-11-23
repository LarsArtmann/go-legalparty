package legalparty

import (
	"strconv"

	errorfamily "github.com/larsartmann/go-error-family"
)

// Spec constants for the VAT ID rules: named so the regulation
// numbers they encode stay greppable next to their use sites.
const (
	minVATIDLength = 3 // two-letter country prefix plus at least one digit
	deVATLength    = 9 // German USt-IdNr. national part is 9 digits

	iso7064Seed   = 5  // ISO 7064 Mod 11,10 seed (python-stdnum reference)
	iso7064Weight = 2  // ISO 7064 Mod 11,10 digit weight
	iso7064Mod11  = 11 // ISO 7064 Mod 11,10 intermediate modulus
	iso7064Mod10  = 10 // ISO 7064 Mod 11,10 final radix
)

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
	if len(v) < minVATIDLength {
		return "", errorfamily.NewRejection("legalparty.vat_id.required",
			"a VAT ID carries a country prefix and a national part, got "+s)
	}

	cc, err := ParseCountryCode(v[:2])
	if err != nil {
		return "", errorfamily.NewRejection("legalparty.vat_id.prefix_invalid",
			"a VAT ID starts with a country prefix: "+err.Error())
	}

	national := v[2:]
	// Austria's UID officially writes as ATU + 8 digits: the leading U is
	// display format, not part of the identifier. Strip it so the 8-digit
	// fixed-length shape check applies uniformly to ATU-form and bare
	// digits input; the stored canonical form keeps the U out only of the
	// shape check — VATID(v) still stores the caller's prefixed spelling.
	if cc == "AT" && len(national) > 0 && national[0] == 'U' {
		national = national[1:]
	}

	if err := validateVATNationalPart(cc, national); err != nil {
		return "", err
	}

	return VATID(v), nil
}

// IsValid reports whether the value parses as a VATID in canonical
// prefixed form, discarding the parsed result. The zero value is not
// valid; unprefixed domestic forms are not guessed.
func (v VATID) IsValid() bool {
	_, err := ParseVATID(string(v))

	return err == nil
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
	if len(national) != deVATLength {
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
	check := iso7064Seed

	for i := range len(national) {
		d := int(national[i] - '0')

		base := check
		if base == 0 {
			base = iso7064Mod10
		}

		check = ((base*iso7064Weight)%iso7064Mod11 + d) % iso7064Mod10
	}

	return check == 1
}
