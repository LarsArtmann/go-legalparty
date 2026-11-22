# Changelog

All notable changes. Dates are 2026 unless stated.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [v0.2.0] - 2026-11-22

### Added

- `Regon` type with `ParseRegon`: the Polish business registry identifier in both forms — 9-digit entity and 14-digit local unit (checksum weights 8-9-2-3-4-5-6-7 mod 11 over the parent prefix). Remainder 10 encodes check digit 0, the opposite regulation edge from NIP; the 14-digit serial is deliberately not checked against the documented 13-weight scheme (KRS-published branch records falsify it). Vectors pin algorytm.org's official numbers, real KRS odpis records, and the remainder-10 edge in both directions.
- `TestParseTaxNumber` (18 vectors): all three `tax_number.*` rejection codes, the DE 10-13 digit bounds, generic-jurisdiction 5-16 bounds, separator stripping, and the VAT-shape refusal — including the lowercase-prefix and country-independent cases. `TestNewRegisterRef`: both `register.*` rejection codes, constructor trimming, and the zero-value `Validate` ordering. Every shipped parse path now carries vectors; FEATURES shows zero `PARTIALLY_FUNCTIONAL` rows.
- CI workflow (`.github/workflows/ci.yml`): race-enabled tests, go vet, govulncheck, golangci-lint, a consumer-import simulation, dprint docs-format check, and the fleet cue workflow-policy vet — every push and PR now gated; SECURITY.md's scanning claim is true.
- Fuzz targets (`fuzz_test.go`) for the checksum parsers (NIP, Regon, VATID, LEI, TaxNumber), seed-corpus mode in CI; closes the last fleet-pattern gap (go-error-family carries 16 fuzz targets, this repo had zero).
- `.config/**` excluded from dprint (machine-managed project-meta metadata, not project docs — `dprint check` passes bare, which CI relies on).
- Consumer verification sweep (2026-11-22, post-v0.1.1): all three fleet consumers (invoices, cqrs-htmx/legalparty-model, company-master-data) built and race-tested against this HEAD — the v0.1.1 tightenings (NL strict shape, AllowsLEI deny-on-invalid) break nothing; regression fixtures pinned in invoices.

### Fixed

- README carried an MIT license badge while LICENSE is proprietary — badge removed.

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
