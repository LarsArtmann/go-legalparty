package legalparty

import (
	errorfamily "github.com/larsartmann/go-error-family"
)

// TaxNumber is a national direct-tax number (the German Steuernummer, UK
// UTR, ...). Deliberately a different type from [VATID]: tax numbers carry
// no country prefix, so a prefixed VAT-shaped value is rejected here rather
// than silently accepted. Germany validates the 10-13 digit national forms;
// other jurisdictions are format-checked.
type TaxNumber string

// ParseTaxNumber validates and constructs a TaxNumber for the given
// jurisdiction.
func ParseTaxNumber(country CountryCode, s string) (TaxNumber, error) {
	v := normalizeTaxIdentifier(s)
	if v == "" {
		return "", errorfamily.NewRejection("legalparty.tax_number.required",
			"a tax number is required")
	}

	if isPrefixedVATShape(v) {
		return "", errorfamily.NewRejection("legalparty.tax_number.vat_shape",
			"a tax number carries no country prefix; parse "+v+" as a VAT ID instead")
	}

	if err := validateTaxNumberBody(country, v); err != nil {
		return "", err
	}

	return TaxNumber(v), nil
}

// validateTaxNumberBody enforces the per-jurisdiction shape after separator
// removal: Germany takes the 10-13 digit national forms, everyone else is
// format-checked.
func validateTaxNumberBody(country CountryCode, v string) error {
	if country == "DE" {
		if len(v) < 10 || len(v) > 13 {
			return errorfamily.NewRejection("legalparty.tax_number.national_invalid",
				"a German Steuernummer has 10 to 13 digits (legacy or ELSTER form)")
		}
	} else if len(v) < 5 || len(v) > 16 {
		return errorfamily.NewRejection("legalparty.tax_number.national_invalid",
			"a "+string(country)+" tax number is 5 to 16 characters after separator removal")
	}

	if !isAllUpperAlphaNumeric(v) {
		return errorfamily.NewRejection("legalparty.tax_number.national_invalid",
			"a tax number carries only letters and digits after separator removal")
	}

	return nil
}

// minVATPrefixedLength is the shortest string the VAT shape can be:
// a two-letter country prefix plus at least one digit.
const minVATPrefixedLength = 3

// isPrefixedVATShape reports whether the value is a two-letter country
// prefix followed by digits only, the canonical VAT ID shape.
func isPrefixedVATShape(v string) bool {
	if len(v) < minVATPrefixedLength {
		return false
	}

	return isAllUpperAlpha(v[:2]) && isAllDigits(v[2:])
}
