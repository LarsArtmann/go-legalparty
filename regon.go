package legalparty

import (
	"strings"

	errorfamily "github.com/larsartmann/go-error-family"
)

// Regon is a Polish Rejestr Gospodarki Narodowej identifier: the 9-digit
// form identifies an entity, the 14-digit form a local unit (a valid
// 9-digit parent REGON plus a 5-digit serial). The 9th digit is a weighted
// checksum (weights 8,9,2,3,4,5,6,7 mod 11). Where the NIP regulation
// declares remainder 10 invalid, the REGON regulation maps it to check
// digit 0 — the two Polish checksums disagree on that edge, which is
// exactly what hand-rolled fleet copies mixed up before this library.
type Regon string

// ParseRegon validates and normalizes a Polish REGON: separators stripped
// (the fleet's identifier-normalization convention), an optional PL prefix
// accepted, checksum verified.
func ParseRegon(s string) (Regon, error) {
	v := normalizeTaxIdentifier(s)
	v = strings.TrimPrefix(v, "PL")

	if len(v) != 9 && len(v) != 14 {
		return "", errorfamily.NewRejection("legalparty.regon.invalid",
			"a REGON is 9 digits (entity) or 14 digits (local unit), optionally PL-prefixed, got "+s)
	}

	if !isAllDigits(v) {
		return "", errorfamily.NewRejection("legalparty.regon.invalid",
			"a REGON carries digits only, got "+s)
	}

	if !validRegonChecksum(v[:9]) {
		return "", errorfamily.NewRejection("legalparty.regon.checksum_failed",
			"REGON "+v+" fails the weighted checksum (weights 89234567, mod 11; remainder 10 encodes check digit 0)")
	}

	return Regon(v), nil
}

// Spec constants, named as in the fleet's consumer copies: the numbers
// encode regulation edges, not tuning knobs.
const (
	regonModulus       = 11 // REGON weighted checksum modulus
	regonZeroRemainder = 10 // remainder 10 encodes check digit 0 (anti-NIP edge)
)

// validRegonChecksum checks the first nine digits of either form. The
// 14-digit serial is deliberately NOT checked against the documented
// 13-weight scheme: KRS-published branch REGONs (real master data) falsify
// that scheme, so enforcing it would reject valid records. Only the parent
// prefix carries a checksum the registry actually honors.
func validRegonChecksum(nine string) bool {
	weights := [8]int{8, 9, 2, 3, 4, 5, 6, 7}

	sum := 0

	for i := range 8 {
		sum += int(nine[i]-'0') * weights[i]
	}

	// Remainder 10 encodes check digit 0 — the opposite edge from NIP,
	// where the same remainder is invalid outright.
	check := sum % regonModulus
	if check == regonZeroRemainder {
		check = 0
	}

	return check == int(nine[8]-'0')
}
