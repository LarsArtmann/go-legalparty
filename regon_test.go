package legalparty

import "testing"

// The REGON vectors pin the weighted checksum with three independent
// sources: algorytm.org's official example numbers, KRS odpis records of
// real Artmann entities, and the remainder-10 edge — where REGON encodes
// check digit 0 while NIP declares the same remainder invalid (the exact
// confusion this extraction ended).
func TestParseRegon(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		in       string
		want     string
		wantCode string
	}{
		{"algorytm.org official 9-digit", "732065814", "732065814", ""},
		{"KRS odpis holding", "527355695", "527355695", ""},
		{"KRS odpis technologies", "527265621", "527265621", ""},
		{"remainder 10 encodes check digit 0", "400000000", "400000000", ""},
		{"algorytm.org official 14-digit", "23511332857188", "23511332857188", ""},
		{"KRS 14-digit local unit", "52735569500000", "52735569500000", ""},
		{"PL prefix and separators strip", "PL 527-355-695", "527355695", ""},
		{"tampered check digit", "527355694", "", "legalparty.regon.checksum_failed"},
		{"tampered body", "527265620", "", "legalparty.regon.checksum_failed"},
		{"remainder-10 edge with nonzero check digit", "400000001", "", "legalparty.regon.checksum_failed"},
		{"14-digit with checksum-invalid parent", "52735569400000", "", "legalparty.regon.checksum_failed"},
		{"too short", "52735569", "", "legalparty.regon.invalid"},
		{"too long", "527355695123456", "", "legalparty.regon.invalid"},
		{"letters", "52735569X", "", "legalparty.regon.invalid"},
		{"empty", "", "", "legalparty.regon.invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ParseRegon(tt.in)
			if tt.wantCode == "" {
				if err != nil {
					t.Fatalf("ParseRegon(%q) must accept: %v", tt.in, err)
				}

				if string(got) != tt.want {
					t.Fatalf("ParseRegon(%q) = %q, want %q", tt.in, got, tt.want)
				}

				return
			}

			if err == nil {
				t.Fatalf("ParseRegon(%q) must reject, got %q", tt.in, got)
			}

			if code := errorCode(t, err); code != tt.wantCode {
				t.Fatalf("error code: got %q, want %q", code, tt.wantCode)
			}
		})
	}
}

func TestRegonIsValid(t *testing.T) {
	t.Parallel()

	if !Regon("732065814").IsValid() {
		t.Fatal("algorytm.org official 9-digit REGON must be valid")
	}

	if !Regon("23511332857188").IsValid() {
		t.Fatal("official 14-digit local-unit REGON must be valid")
	}

	if Regon("52735569400000").IsValid() {
		t.Fatal("checksum-invalid parent must be invalid")
	}

	if Regon("").IsValid() || Regon("73206581").IsValid() {
		t.Fatal("zero value and short input must be invalid")
	}
}
