package legalparty

import (
	"strings"
	"testing"

	errorfamily "github.com/larsartmann/go-error-family"
)

// The fuzz targets pin the parsers' crash-safety and one structural
// invariant each: whatever the library ACCEPTS must re-parse to itself
// (normalization is idempotent — a canonical form that fails re-parse is
// a normalization bug), and whatever it REJECTS must carry a rejection
// code from the frozen legalparty.* namespace. Seeds are the vector-table
// entries: valid, tampered, separator, and shape-confusion forms.

func FuzzParseNIP(f *testing.F) {
	f.Add("5260000152")
	f.Add("PL 526-000-01-52")
	f.Add("5260000153")
	f.Add("52600001A2")
	f.Add("")
	f.Add("pl5260000152")

	f.Fuzz(func(t *testing.T, input string) {
		got, err := ParseNIP(input)
		if err != nil {
			if code := errorfamily.Code(err); !strings.HasPrefix(code, "legalparty.") {
				t.Fatalf("ParseNIP(%q) rejected without a legalparty.* code: %v", input, err)
			}

			return
		}

		again, err := ParseNIP(string(got))
		if err != nil || again != got {
			t.Fatalf("ParseNIP not idempotent: %q accepted, re-parse gave %q, %v", input, again, err)
		}
	})
}

func FuzzParseVATID(f *testing.F) {
	f.Add("DE136695976")
	f.Add("de 136 695 976")
	f.Add("DE136695978")
	f.Add("ATU13585627")
	f.Add("NL123456789B01")
	f.Add("PL1234567890")
	f.Add("FR12345678901")
	f.Add("ZZ123456789")
	f.Add("")

	f.Fuzz(func(t *testing.T, input string) {
		got, err := ParseVATID(input)
		if err != nil {
			if code := errorfamily.Code(err); !strings.HasPrefix(code, "legalparty.") {
				t.Fatalf("ParseVATID(%q) rejected without a legalparty.* code: %v", input, err)
			}

			return
		}

		again, err := ParseVATID(string(got))
		if err != nil || again != got {
			t.Fatalf("ParseVATID not idempotent: %q accepted, re-parse gave %q, %v", input, again, err)
		}
	})
}

func FuzzParseLEI(f *testing.F) {
	f.Add("984500663D3E060B3G67")
	f.Add("7245007cs2lxyd2jct52")
	f.Add("7245007CS2LXYD2JCT53")
	f.Add("5493001KJTIIGC8Y1V1!")
	f.Add("")

	f.Fuzz(func(t *testing.T, input string) {
		got, err := ParseLEI(input)
		if err != nil {
			if code := errorfamily.Code(err); !strings.HasPrefix(code, "legalparty.") {
				t.Fatalf("ParseLEI(%q) rejected without a legalparty.* code: %v", input, err)
			}

			return
		}

		again, err := ParseLEI(string(got))
		if err != nil || again != got {
			t.Fatalf("ParseLEI not idempotent: %q accepted, re-parse gave %q, %v", input, again, err)
		}
	})
}

func FuzzParseTaxNumber(f *testing.F) {
	f.Add("DE", "218/515/08723")
	f.Add("DE", "21851508723123")
	f.Add("DE", "DE136695976")
	f.Add("FR", "1234567890")
	f.Add("GB", "1234")
	f.Add("DE", "  /-. ")

	f.Fuzz(func(t *testing.T, country, input string) {
		got, err := ParseTaxNumber(CountryCode(country), input)
		if err != nil {
			if code := errorfamily.Code(err); !strings.HasPrefix(code, "legalparty.") {
				t.Fatalf("ParseTaxNumber(%q, %q) rejected without a legalparty.* code: %v", country, input, err)
			}

			return
		}

		again, err := ParseTaxNumber(CountryCode(country), string(got))
		if err != nil || again != got {
			t.Fatalf(
				"ParseTaxNumber not idempotent: (%q, %q) accepted, re-parse gave %q, %v",
				country, input, again, err,
			)
		}
	})
}

func FuzzParseRegon(f *testing.F) {
	f.Add("527355695")
	f.Add("PL 527-355-695")
	f.Add("527355694")
	f.Add("400000000")
	f.Add("52735569500000")
	f.Add("527355695123456")
	f.Add("")

	f.Fuzz(func(t *testing.T, input string) {
		got, err := ParseRegon(input)
		if err != nil {
			if code := errorfamily.Code(err); !strings.HasPrefix(code, "legalparty.") {
				t.Fatalf("ParseRegon(%q) rejected without a legalparty.* code: %v", input, err)
			}

			return
		}

		again, err := ParseRegon(string(got))
		if err != nil || again != got {
			t.Fatalf("ParseRegon not idempotent: %q accepted, re-parse gave %q, %v", input, again, err)
		}
	})
}
