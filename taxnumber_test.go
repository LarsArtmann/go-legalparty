package legalparty

import "testing"

// The Steuernummer vectors pin the shape contract: Germany takes the
// 10-13 digit national forms (format-checked — no public checksum exists),
// everyone else 5-16 characters, and a VAT-shaped value is refused rather
// than silently accepted as a tax number.
func TestParseTaxNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		country  CountryCode
		in       string
		want     string
		wantCode string
	}{
		{"german 10-digit form", "DE", "2185150872", "2185150872", ""},
		{"german 11-digit form with separators", "DE", "218/515/08723", "21851508723", ""},
		{"german 13-digit ELSTER form", "DE", "2185150872312", "2185150872312", ""},
		{"german 9 digits is too short", "DE", "218515087", "", "legalparty.tax_number.national_invalid"},
		{"german 14 digits is too long", "DE", "21851508723123", "", "legalparty.tax_number.national_invalid"},
		{"generic jurisdiction 10 digits", "FR", "1234567890", "1234567890", ""},
		{"generic jurisdiction lower bound 5", "FR", "12345", "12345", ""},
		{"generic jurisdiction upper bound 16", "FR", "1234567890123456", "1234567890123456", ""},
		{"generic jurisdiction 4 chars is too short", "FR", "1234", "", "legalparty.tax_number.national_invalid"},
		{
			"generic jurisdiction 17 chars is too long",
			"FR",
			"12345678901234567",
			"",
			"legalparty.tax_number.national_invalid",
		},
		{"letters survive in generic form", "GB", "1234567890A", "1234567890A", ""},
		{"non-separator punctuation rejects", "GB", "12345678_9", "", "legalparty.tax_number.national_invalid"},
		{"spaces and dots strip first", "DE", " 218.515.087 23 ", "21851508723", ""},
		{"vat-shaped value refused", "DE", "DE136695976", "", "legalparty.tax_number.vat_shape"},
		{"vat shape catches lowercase prefix", "PL", "pl1234567890", "", "legalparty.tax_number.vat_shape"},
		{"vat shape is country-independent", "FR", "FR123456789", "", "legalparty.tax_number.vat_shape"},
		{"empty", "DE", "", "", "legalparty.tax_number.required"},
		{"only separators is empty", "DE", "  /-. ", "", "legalparty.tax_number.required"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ParseTaxNumber(tt.country, tt.in)
			if tt.wantCode == "" {
				if err != nil {
					t.Fatalf("ParseTaxNumber(%q, %q) must accept: %v", tt.country, tt.in, err)
				}

				if string(got) != tt.want {
					t.Fatalf("ParseTaxNumber(%q, %q) = %q, want %q", tt.country, tt.in, got, tt.want)
				}

				return
			}

			if err == nil {
				t.Fatalf("ParseTaxNumber(%q, %q) must reject, got %q", tt.country, tt.in, got)
			}

			if code := errorCode(t, err); code != tt.wantCode {
				t.Fatalf("error code: got %q, want %q", code, tt.wantCode)
			}
		})
	}
}
