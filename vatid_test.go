package legalparty

import "testing"

func TestVATIDIsValid(t *testing.T) {
	t.Parallel()

	if !VATID("DE136695976").IsValid() {
		t.Fatal("BMF example DE136695976 must be valid")
	}

	if VATID("DE136695978").IsValid() {
		t.Fatal("wrong DE checksum must be invalid")
	}

	if VATID("NLA23456789B01").IsValid() {
		t.Fatal("letters in NL digit positions must be invalid")
	}

	if VATID("").IsValid() || VATID("5260000152").IsValid() {
		t.Fatal("zero value and prefix-less input must be invalid (no country is guessed)")
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
