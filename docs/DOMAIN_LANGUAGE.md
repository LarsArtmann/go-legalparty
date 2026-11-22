# Domain Language

A **Unified Language** for `go-legalparty` — shared across Customer, Product
Owner, Developer, and AI. Inspired by Domain-Driven Design (DDD) Ubiquitous
Language. Every term below should mean the **same thing** to everyone who
reads it.

This library IS vocabulary: its types are the fleet's canonical names for
legal-party identifiers. Provenance: extracted 2026-11-21 from
`cqrs-htmx/legalparty-model` per ADR-0001, ending the checksum split brain
(invoices hand-rolled its own NIP/VAT validators, with a real bug).

## Value Objects

| Term               | Definition                                                                                                                       | Context / Invariants                                                                                                                                                                                                                  |
| ------------------ | -------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Legal Party**    | Any entity that can appear on a business document: natural person, company, partnership, public body, non-profit                 | The aggregate lives in `cqrs-htmx/legalparty-model`; this library owns the identifier value types                                                                                                                                     |
| **NIP**            | Polish Numer Identyfikacji Podatkowej: 10 digits, weighted checksum (6-5-7-2-3-4-5-6-7 mod 11)                                   | Remainder 10 is INVALID by regulation (not mapped to 0) — the edge the old hand-rolled copies got wrong; stored bare, PL prefix accepted on input                                                                                     |
| **Regon**          | Polish Rejestr Gospodarki Narodowej identifier: 9 digits (entity) or 14 (local unit), weighted checksum (8-9-2-3-4-5-6-7 mod 11) | Remainder 10 encodes check digit 0 — the OPPOSITE edge from NIP; 14-digit form = valid 9-digit parent + 5-digit serial (serial unchecked: KRS-published records falsify the documented 13-weight scheme); PL prefix accepted on input |
| **VATID**          | Value-added tax ID in canonical EU-prefixed form ("DE136695976")                                                                 | Different type from TaxNumber by design: a VAT ID carries a country prefix, a national tax number never does. DE validates ISO 7064 Mod 11,10 (seed 5); NL carries the literal B; AT 8 / PL 10 digits; others generic 8-12            |
| **TaxNumber**      | National direct-tax number (German Steuernummer, UK UTR, ...)                                                                    | No country prefix; VAT-shaped input is rejected and routed to ParseVATID. DE: 10-13 digits; others: generic 5-16                                                                                                                      |
| **LEI**            | Legal Entity Identifier (ISO 17442): 20 chars [A-Z0-9], mod-97 checksum == 1                                                     | Issued under GLEIF governance; natural persons cannot hold one (`LegalFormClass.AllowsLEI`)                                                                                                                                           |
| **RegisterRef**    | A registry + entry number pair (HRB, KRS, Companies House...)                                                                    | `Registry` is opaque by design — semantics owned by whoever resolves the reference; the domain never inspects it                                                                                                                      |
| **CountryCode**    | ISO 3166-1 alpha-2 code of the governing jurisdiction                                                                            | Officially-assigned set only — XK (Kosovo), EU, ZZ deliberately rejected; a billing jurisdiction must resolve                                                                                                                         |
| **LegalFormClass** | One of five peer classes: natural_person, legal_entity, partnership, public_body, non_profit                                     | Peers, not a hierarchy. legal_entity + partnership must anchor to a registry (`RequiresRegistryAnchor`); sole traders are natural persons                                                                                             |

## Concepts

| Term                        | Definition                                                                               | Context                                                                                                  |
| --------------------------- | ---------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| **Separator normalization** | Fleet-wide input convention: uppercase, trim, strip spaces/`-`/`.`/`/` before validation | `normalizeTaxIdentifier`; canonical forms are what the types store                                       |
| **Rejection**               | A `go-error-family` Rejection with a stable machine code `legalparty.<subject>.<cause>`  | Codes are the public API (dashboards key on them); messages are human prose carrying the offending input |
| **Canonical form**          | The normalized stored representation of an identifier                                    | VATID keeps its prefix; NIP and Regon drop the PL prefix; LEI is uppercased                              |

## Anti-language (terms deliberately NOT used)

| Avoid                   | Use instead                   | Why                                                                                           |
| ----------------------- | ----------------------------- | --------------------------------------------------------------------------------------------- |
| "VAT number" (bare)     | VATID or TaxNumber            | Ambiguous — the two are distinct types and conflating them rejects valid values on both sides |
| "Tax ID"                | Name the concrete type        | Same ambiguity; the invoices repo's historic tax_id_validation.go is the cautionary tale      |
| Registry code semantics | opaque `RegisterRef.Registry` | The domain must not inspect registry codes; vocabulary belongs to the resolver                |
