package legalparty

import (
	"strings"

	errorfamily "github.com/larsartmann/go-error-family"
)

// NIP is a Polish Numer Identyfikacji Podatkowej: 10 digits whose last
// digit is a weighted checksum (weights 6,5,7,2,3,4,5,6,7 mod 11; a
// remainder of 10 is invalid by regulation, not zero — the checksum
// space deliberately excludes it). Carried unprefixed or as PL+digits.
type NIP string

// ParseNIP validates and normalizes a Polish NIP: separators stripped
// (the fleet's identifier-normalization convention), an optional PL
// prefix accepted, checksum verified.
func ParseNIP(s string) (NIP, error) {
	v := normalizeTaxIdentifier(strings.ToUpper(s))
	v = strings.TrimPrefix(v, "PL")

	if len(v) != 10 {
		return "", errorfamily.NewRejection("legalparty.nip.invalid",
			"a Polish NIP is 10 digits (optionally PL-prefixed), got "+s)
	}

	weights := [9]int{6, 5, 7, 2, 3, 4, 5, 6, 7}

	if !isAllDigits(v) {
		return "", errorfamily.NewRejection("legalparty.nip.invalid",
			"a Polish NIP carries digits only, got "+s)
	}

	sum := 0

	for i := range 9 {
		sum += int(v[i]-'0') * weights[i]
	}

	// Remainder 10 is INVALID by Polish regulation (the checksum space
	// excludes it) — the mod-11 result must equal the check digit as-is.
	if sum%11 != int(v[9]-'0') {
		return "", errorfamily.NewRejection("legalparty.nip.checksum_failed",
			"NIP "+v+" fails the weighted checksum (weights 657234567, mod 11; remainder 10 is invalid)")
	}

	return NIP(v), nil
}
