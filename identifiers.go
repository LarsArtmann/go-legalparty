package legalparty

import (
	"strconv"
	"strings"

	errorfamily "github.com/larsartmann/go-error-family"
)

// Spec constants for the identifier checksums: named so the regulation
// numbers they encode stay greppable next to their use sites.
const (
	leiLength  = 20 // ISO 17442: a LEI is exactly 20 characters
	leiModulus = 97 // ISO 17442: mod-97 check over the mapped string
	letterBase = 10 // ISO 17442: letters map A=10 .. Z=35
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
	if len(v) != leiLength {
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

// IsValid reports whether the value parses as a LEI, discarding the
// parsed result: the one-call validity check for callers that need no
// normalized form. The zero value is not valid.
func (l LEI) IsValid() bool {
	_, err := ParseLEI(string(l))

	return err == nil
}

// validLEIChecksum applies the ISO 17442 check: map every character to its
// numeric form (digits as-is, A=10 .. Z=35), interpret the whole 20-character
// string as one decimal number, and require number mod 97 == 1.
func validLEIChecksum(code string) bool {
	remainder := 0

	for _, r := range code {
		switch {
		case r >= '0' && r <= '9':
			remainder = (remainder*10 + int(r-'0')) % leiModulus
		case r >= 'A' && r <= 'Z':
			v := int(r-'A') + letterBase
			remainder = (remainder*10 + v/10) % leiModulus
			remainder = (remainder*10 + v%10) % leiModulus
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

// NewRegisterRef validates and constructs a RegisterRef. It is the
// normalization point: both halves are trimmed of surrounding
// whitespace before the presence check, so the stored pair is canonical
// (interior whitespace is kept — registry entry numbers carry it).
func NewRegisterRef(registry, number string) (RegisterRef, error) {
	r := RegisterRef{Registry: strings.TrimSpace(registry), Number: strings.TrimSpace(number)}
	if err := r.Validate(); err != nil {
		return RegisterRef{}, err
	}

	return r, nil
}

// Validate enforces that both the registry code and the number are
// present. It is a presence check, not a normalization point: values with
// surrounding whitespace pass (a raw struct predating the constructor is
// not made invalid by it), and Validate never mutates. Use
// NewRegisterRef to store canonical, trimmed pairs.
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
