# Changelog

All notable changes. Dates are 2026 unless stated.

## [v0.1.0] - 2026-11-21

### Added

- Extracted from `cqrs-htmx/legalparty-model`: `CountryCode` (assigned ISO 3166-1 alpha-2 only), `LegalFormClass`, `LEI` (ISO 17442 mod-97), `VATID` (DE ISO 7064 Mod 11,10 checksum; NL BTW; fixed-length; generic EU), `TaxNumber`, `RegisterRef` — separator normalization and go-error-family rejection codes throughout.
- New `NIP` with the regulation-correct weighted checksum (weights 6-5-7-2-3-4-5-6-7 mod 11, remainder 10 invalid) — fixing the class of bug in hand-rolled copies that mapped remainder 10 to check digit 0.
- Checksum-vector tests pinning live GLEIF records, the BMF DE example, and the MoF PL example.
