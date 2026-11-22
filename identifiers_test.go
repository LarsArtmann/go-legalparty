package legalparty

import (
	"testing"

	errorfamily "github.com/larsartmann/go-error-family"
)

// errorCode digs the rejection code out, so vector tables can pin the
// exact machine-readable failure.
func errorCode(t *testing.T, err error) string {
	t.Helper()

	code := errorfamily.Code(err)
	if code == "" {
		t.Fatalf("error %v carries no rejection code", err)
	}

	return code
}

// leiLiveDE and leiLiveNL are live GLEIF records (fetched 2026-11-20 by
// the cqrs-htmx legalparty module; carried here so the vectors stay
// pinned without network access).
const (
	leiLiveDE = "984500663D3E060B3G67"
	leiLiveNL = "7245007CS2LXYD2JCT52"
)

func TestParseLEI(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		in       string
		want     string
		wantCode string
	}{
		{"live German record", leiLiveDE, leiLiveDE, ""},
		{"live Dutch record", leiLiveNL, leiLiveNL, ""},
		{"lowercase normalizes", "7245007cs2lxyd2jct52", leiLiveNL, ""},
		{"checksum tampered", "7245007CS2LXYD2JCT53", "", "legalparty.lei.checksum_failed"},
		{"too short", "5493001KJTIIGC8Y1V1", "", "legalparty.lei.invalid"},
		{"bad character", "5493001KJTIIGC8Y1V1!", "", "legalparty.lei.invalid"},
		{"empty", "", "", "legalparty.lei.invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ParseLEI(tt.in)
			if tt.wantCode == "" {
				if err != nil {
					t.Fatalf("ParseLEI(%q) must accept: %v", tt.in, err)
				}

				if string(got) != tt.want {
					t.Fatalf("ParseLEI(%q) = %q, want %q", tt.in, got, tt.want)
				}

				return
			}

			if err == nil {
				t.Fatalf("ParseLEI(%q) must reject, got %q", tt.in, got)
			}

			if code := errorCode(t, err); code != tt.wantCode {
				t.Fatalf("error code: got %q, want %q", code, tt.wantCode)
			}
		})
	}
}

func TestParseVATID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		in       string
		want     string
		wantCode string
	}{
		{"BMF example passes checksum", "DE136695976", "DE136695976", ""},
		{"tampered checksum rejected", "DE136695978", "", "legalparty.vat_id.checksum_failed"},
		{"separators normalize", "de 136 695 976", "DE136695976", ""},
		{"austrian UID format", "ATU13585627", "ATU13585627", ""},
		{"polish NIP-shaped VAT", "PL1234567890", "PL1234567890", ""},
		{"dutch BTW", "NL123456789B01", "NL123456789B01", ""},
		{"dutch letters in digit positions", "NLA23456789B01", "", "legalparty.vat_id.national_invalid"},
		{"dutch B misplaced", "NL123456B78901", "", "legalparty.vat_id.national_invalid"},
		{"generic EU form", "FR12345678901", "FR12345678901", ""},
		{"no country prefix", "136695976", "", "legalparty.vat_id.prefix_invalid"},
		{"unassigned prefix", "ZZ123456789", "", "legalparty.vat_id.prefix_invalid"},
		{"german wrong length", "DE1366959", "", "legalparty.vat_id.national_invalid"},
		{"too short", "DE12", "", "legalparty.vat_id.national_invalid"},
		{"empty", "  ", "", "legalparty.vat_id.required"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ParseVATID(tt.in)
			if tt.wantCode == "" {
				if err != nil {
					t.Fatalf("ParseVATID(%q) must accept: %v", tt.in, err)
				}

				if string(got) != tt.want {
					t.Fatalf("ParseVATID(%q) = %q, want %q", tt.in, got, tt.want)
				}

				return
			}

			if err == nil {
				t.Fatalf("ParseVATID(%q) must reject, got %q", tt.in, got)
			}

			if code := errorCode(t, err); code != tt.wantCode {
				t.Fatalf("error code: got %q, want %q", code, tt.wantCode)
			}
		})
	}
}

// The NIP vectors pin the weighted checksum: 5260000152 is the Polish
// Ministry of Finance's own example number; the remainder-10 case pins
// the regulation edge invoices' hand-rolled version got wrong (it
// mapped remainder 10 to check digit 0 and accepted an invalid NIP).
func TestParseNIP(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		in       string
		want     string
		wantCode string
	}{
		{"MoF example passes checksum", "5260000152", "5260000152", ""},
		{"PL prefix accepted", "PL 526-000-01-52", "5260000152", ""},
		{"tampered check digit", "5260000153", "", "legalparty.nip.checksum_failed"},
		{"wrong length", "526000015", "", "legalparty.nip.invalid"},
		{"letters", "526000015A", "", "legalparty.nip.invalid"},
		{"empty", "", "", "legalparty.nip.invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ParseNIP(tt.in)
			if tt.wantCode == "" {
				if err != nil {
					t.Fatalf("ParseNIP(%q) must accept: %v", tt.in, err)
				}

				if string(got) != tt.want {
					t.Fatalf("ParseNIP(%q) = %q, want %q", tt.in, got, tt.want)
				}

				return
			}

			if err == nil {
				t.Fatalf("ParseNIP(%q) must reject, got %q", tt.in, got)
			}

			if code := errorCode(t, err); code != tt.wantCode {
				t.Fatalf("error code: got %q, want %q", code, tt.wantCode)
			}
		})
	}
}

func TestParseCountryCodeAndLegalForm(t *testing.T) {
	t.Parallel()

	if got, err := ParseCountryCode("de"); err != nil || got != "DE" {
		t.Fatalf("ParseCountryCode lowercases-normalizes: got %q err %v", got, err)
	}

	if _, err := ParseCountryCode("ZZ"); err == nil {
		t.Fatal("unassigned ZZ must reject")
	}

	if got, err := ParseLegalFormClass("legal_entity"); err != nil || got != LegalFormLegalEntity {
		t.Fatalf("ParseLegalFormClass: got %q err %v", got, err)
	}

	if _, err := ParseLegalFormClass("corp"); err == nil {
		t.Fatal("unknown legal form class must reject")
	}
}

// The rule vectors pin both class rules for every defined value AND for
// values that never passed ParseLegalFormClass: an invalid class must
// allow nothing, so the two rules cannot disagree on unknown input.
func TestLegalFormRules(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		class      LegalFormClass
		wantAnchor bool
		wantLEI    bool
	}{
		{"natural person", LegalFormNaturalPerson, false, false},
		{"legal entity", LegalFormLegalEntity, true, true},
		{"partnership", LegalFormPartnership, true, true},
		{"public body", LegalFormPublicBody, false, true},
		{"non-profit", LegalFormNonProfit, false, true},
		{"empty zero value", "", false, false},
		{"unknown class", LegalFormClass("corp"), false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.class.RequiresRegistryAnchor(); got != tt.wantAnchor {
				t.Fatalf("RequiresRegistryAnchor(%q) = %v, want %v", tt.class, got, tt.wantAnchor)
			}

			if got := tt.class.AllowsLEI(); got != tt.wantLEI {
				t.Fatalf("AllowsLEI(%q) = %v, want %v", tt.class, got, tt.wantLEI)
			}
		})
	}
}

// RegisterRef is the one non-checksummed constructor: the vectors pin
// presence of both halves plus the trimming, and that the same rules hold
// when Validate runs on a zero value built without the constructor.
func TestNewRegisterRef(t *testing.T) {
	t.Parallel()

	t.Run("happy path trims both halves", func(t *testing.T) {
		t.Parallel()

		got, err := NewRegisterRef(" de_munich ", " HRB 12345 ")
		if err != nil {
			t.Fatalf("NewRegisterRef must accept: %v", err)
		}

		if got.Registry != "de_munich" || got.Number != "HRB 12345" {
			t.Fatalf("NewRegisterRef trims: got %+v", got)
		}
	})

	t.Run("registry required", func(t *testing.T) {
		t.Parallel()

		if _, err := NewRegisterRef("   ", "HRB 12345"); err == nil {
			t.Fatal("blank registry must reject")
		} else if code := errorCode(t, err); code != "legalparty.register.registry_required" {
			t.Fatalf("error code: got %q, want legalparty.register.registry_required", code)
		}
	})

	t.Run("number required", func(t *testing.T) {
		t.Parallel()

		if _, err := NewRegisterRef("de_munich", ""); err == nil {
			t.Fatal("blank number must reject")
		} else if code := errorCode(t, err); code != "legalparty.register.number_required" {
			t.Fatalf("error code: got %q, want legalparty.register.number_required", code)
		}
	})

	t.Run("zero value fails validate with registry first", func(t *testing.T) {
		t.Parallel()

		err := RegisterRef{}.Validate()
		if err == nil {
			t.Fatal("zero RegisterRef must fail Validate")
		}

		if code := errorCode(t, err); code != "legalparty.register.registry_required" {
			t.Fatalf("error code: got %q, want legalparty.register.registry_required", code)
		}
	})

	t.Run("raw struct validate tolerates untrimmed input", func(t *testing.T) {
		t.Parallel()

		// Trim semantics, pinned both ways: the constructor is the
		// normalization point (trims, per the happy-path vector above);
		// Validate is presence-only — it accepts surrounding whitespace
		// and never mutates, so a raw struct predating the constructor
		// stays valid.
		r := RegisterRef{Registry: " HRB ", Number: " 12345 "}
		if err := r.Validate(); err != nil {
			t.Fatalf("Validate must accept untrimmed values: %v", err)
		}

		if r.Registry != " HRB " || r.Number != " 12345 " {
			t.Fatalf("Validate must not mutate: got %+v", r)
		}
	})
}
