# go-legalparty

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Legal-party identifiers with real checksums: the value layer under every German/Polish business document in the fleet. Extracted 2026-11-21 from `cqrs-htmx/legalparty-model` (which keeps the CQRS aggregate; this is the types) — ending the checksum split brain the domain-language census counted (invoices hand-rolled its own NIP/VAT validators, with a real bug: it accepted NIPs whose weighted sum mod 11 is 10, which Polish regulation defines as invalid).

```go
import "github.com/larsartmann/go-legalparty"

nip, err := legalparty.ParseNIP("PL 526-000-01-52")    // checksum verified
vat, err := legalparty.ParseVATID("DE136695976")        // ISO 7064 Mod 11,10
lei, err := legalparty.ParseLEI("984500663D3E060B3G67") // ISO 17442 mod-97
tax, err := legalparty.ParseTaxNumber("DE", "217/8150/8173") // Steuernummer shape
ref, err := legalparty.NewRegisterRef("de_munich", "HRB 123456")
cc,  err := legalparty.ParseCountryCode("de")            // → DE (assigned only)
lf,  err := legalparty.ParseLegalFormClass("legal_entity")
```

## Installation

```
go get github.com/larsartmann/go-legalparty
```

## What it validates

| Type             | Rules carried                                                                                                                              |
| ---------------- | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `NIP`            | 10 digits, weights 6‑5‑7‑2‑3‑4‑5‑6‑7 mod 11; remainder 10 is invalid (regulation edge invoices got wrong)                                  |
| `VATID`          | two-letter assigned CC prefix + national part: DE full ISO 7064 Mod 11,10 checksum; NL BTW shape; AT/PL fixed lengths; generic EU fallback |
| `LEI`            | 20 chars, ISO 17442 mod‑97 checksum, case-normalized                                                                                       |
| `TaxNumber`      | national tax numbers, no country prefix (DE Steuernummer 10–13 digits; other jurisdictions format-checked 5–16)                            |
| `RegisterRef`    | registry + number (HRB, KRS, …)                                                                                                            |
| `CountryCode`    | ISO 3166‑1 alpha‑2 officially-assigned set only (no XK/EU/ZZ)                                                                              |
| `LegalFormClass` | natural_person / legal_entity / partnership / public_body / non_profit                                                                     |

Separators (space, `-`, `.`, `/`) normalize away; every rejection is a `go-error-family` Rejection with a stable machine code (`legalparty.nip.checksum_failed`, …).

## Dependencies

One: `go-error-family` (the fleet's error taxonomy). Everything else stdlib.

## Provenance and migration (ADR-0001)

- Extracted from `cqrs-htmx/legalparty-model` v4 (files: identifiers/country/legalform; the module keeps decider/fold/events and consumes this library).
- `invoices` and `company-master-data` validate through this library; `company-master-data` keeps a local REGON checksum until coverage lands here (see TODO_LIST).
- Migration history: CHANGELOG.
