# Roadmap

> Long-term direction and raw ideas NOT yet refined into actionable tasks.
> When an idea becomes bounded and actionable, it graduates to `TODO_LIST.md`.

## Themes

### Registry IO (demand-gated)

A GLEIF/registry lookup client does not belong in this library until a
**second consumer** exists (decision record: `STACK.md` "Deliberately
excluded"; the `gleif/` code lives on in `cqrs-htmx/legalparty-model`).
Trigger: any second repo needing registry resolution — then extract, with its
own STACK.md row.

### More jurisdiction checksums

- **REGON shipped in v0.2.0** (9-digit and 14-digit forms, remainder-10-means-0
  edge). The remaining half — deleting company-master-data's local copy once
  it adopts `ParseRegon` — is `TODO_LIST.md` work, not roadmap.
- **PL VATID checksum**: `ParseVATID` accepts PL national parts on shape only
  (10 digits); no NIP checksum runs inside the EU-prefixed form. Closing that
  last known message-vs-check gap is demand-gated (a consumer billing
  EU-prefixed PL VAT IDs) and would reuse the NIP weights with consumer-repo
  test updates.
- VAT national checksums beyond DE (FR/IT/ES/... currently fall to the
  generic 8-12 shape check). Add per-jurisdiction rules only when a consumer
  bills that jurisdiction — each rule needs pinned official vectors.
- TaxNumber jurisdiction rules beyond DE's 10-13 digits (UK UTR, PL NIP-linked
  form) — same demand gate.

### GLEIF vector cadence (report 19-26 §f/38)

The LEI vectors pin live GLEIF records as constants so tests need no network;
nothing re-validates them when those records change or lapse upstream. A
scripted refresh (re-fetch pinned LEIs, regenerate vectors, diff) turns that
cadence manual-now, automatic-later. Trigger: first time a pinned record's
status changes, or a second LEI-carrying consumer appears — whichever first.

### Jurisdiction rule seam (deepening candidate 2 — speculative)

If a third or fourth jurisdiction grows real per-country rules, the
country-switch inside `ParseVATID`/`ParseTaxNumber` deserves a deliberate
rule-table seam instead of accreting switch cases. Until that demand exists
the seam is speculative scaffolding — do not build it. Source:
`docs/architecture-understanding/2026-11-22_19-10_value-library-deepening.html`.

### Modularization revisit triggers

The single-module verdict (`docs/modularization/2026-11-22_SINGLE_MODULE_VERDICT.html`)
holds while all of these stay false. Revisit the split if any fires:

- a second independent consumer appears for a strict subset (same trigger as
  registry IO extraction),
- a jurisdiction explosion makes the identifier files unwieldy (rule seam
  above lands first),
- any IO dependency (registry client, cache) forces a physical boundary.

### Explicitly out of scope

- Fleet domain-language census (report 16-24 §f/23): dropped for this repo —
  the fleet-wide census is owned by the `census-digest` skill at `~/projects`
  scope; this repo's vocabulary lives in `docs/DOMAIN_LANGUAGE.md` and needs
  no separate census row.

## Open questions

- Should `RegisterRef.Registry` ever grow a controlled vocabulary, or does
  opacity remain the contract? (Current decision: opaque by design —
  `identifiers.go:86`. Revisit only if two consumers disagree on registry
  codes.)
