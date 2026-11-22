# Changelog

All notable changes. Dates are 2026 unless stated.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [Unreleased]

### Added

- `TestParseTaxNumber` (18 vectors): all three `tax_number.*` rejection codes, the DE 10-13 digit bounds, generic-jurisdiction 5-16 bounds, separator stripping, and the VAT-shape refusal — including the lowercase-prefix and country-independent cases. `TestNewRegisterRef`: both `register.*` rejection codes, constructor trimming, and the zero-value `Validate` ordering. Every shipped parse path now carries vectors; FEATURES shows zero `PARTIALLY_FUNCTIONAL` rows.
- Released v0.1.1 (2026-11-22): corrected SECURITY.md advisory URL, truthful README table, legal-form and Dutch-VAT rule fixes.

## [v0.1.1] - 2026-11-22

### Added

- Repo doc corpus for agent and docs-gates operation: `AGENTS.md` (session context: scope boundary, frozen rejection-code contract, regulation edges, consumer blast radius), `TODO_LIST.md` (REGON support; test vectors for the untested parse paths), `FEATURES.md` (honest inventory — TaxNumber, RegisterRef, and the LegalFormClass rule methods are `PARTIALLY_FUNCTIONAL` pending vectors), `ROADMAP.md` (demand-gated: registry IO, more jurisdiction checksums), `docs/DOMAIN_LANGUAGE.md` (glossary + anti-language).
- `.golangci.yml` seeded and expanded via buildflow's `golangci-lint-auto-configure` (the repo previously carried no lint config).
- README installation section and example coverage for `TaxNumber` + `RegisterRef`.

### Fixed

- README validation table drifted from code: `LegalFormClass` listed "branch/trust" (actual classes: public_body, non_profit), `VATID` claimed an FI fixed length (only AT/PL exist), `TaxNumber` claimed a PL NIP-linked rule (only DE has jurisdiction rules). All rows now match `legalform.go`/`identifiers.go`.
- SECURITY.md carried template drift from another repo: supported-version table claimed v0.1.2 (v0.1.0 is current), the vulnerability-report advisory URL pointed at `go-must` instead of `go-legalparty`, and the practices list claimed CI scanning with no CI present.
- `LegalFormClass.AllowsLEI` returned true for values that never passed `ParseLegalFormClass` (an unknown class allowed a LEI while `RequiresRegistryAnchor` denied the same input — the two rules disagreed). Invalid classes now allow nothing; pinned by `TestLegalFormRules`.
- `ParseVATID` accepted Dutch national parts with letters in the nine digit positions (e.g. `NLA23456789B01`) even though its own rejection message promises 9 digits, B, 2 digits. The NL check now enforces that exact shape.

### Changed

- Migration narrative single-homed: README/STACK/AGENTS state current consumer truth, this CHANGELOG owns the history — cqrs-htmx/legalparty-model consumes this library with its duplicate value sources removed, invoices delegates its checksums here (remainder-10 NIP bug fixed), company-master-data adopts directly with only REGON still local.
- AGENTS.md hardened per the docs-health agents-quality guide: temporal qualifiers stripped, dprint remote-plugin and stale-cache gotchas added, the go-structure-linter "use internal/" advisory recorded as a deliberate non-fix.
- Character-class checks (digits, uppercase letters, alphanumerics) single-homed into `isAllDigits`/`isAllUpperAlpha`/`isAllUpperAlphaNumeric`; the five hand-rolled rune loops now share one implementation. `validateFixedLengthVAT` drops its per-call map allocation for a plain switch.
- `TaxNumber` extracted from `identifiers.go` into `taxnumber.go` (file-size gate; keeps the root-level file-per-concept convention, mirroring `nip.go`).

## [v0.1.0] - 2026-11-21

### Added

- Extracted from `cqrs-htmx/legalparty-model`: `CountryCode` (assigned ISO 3166-1 alpha-2 only), `LegalFormClass`, `LEI` (ISO 17442 mod-97), `VATID` (DE ISO 7064 Mod 11,10 checksum; NL BTW; fixed-length; generic EU), `TaxNumber`, `RegisterRef` — separator normalization and go-error-family rejection codes throughout.
- New `NIP` with the regulation-correct weighted checksum (weights 6-5-7-2-3-4-5-6-7 mod 11, remainder 10 invalid) — fixing the class of bug in hand-rolled copies that mapped remainder 10 to check digit 0.
- Checksum-vector tests pinning live GLEIF records, the BMF DE example, and the MoF PL example.
