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

- **REGON** is the next checksum family (already `TODO_LIST.md`; the local
  copy in `company-master-data` is the demand signal).
- VAT national checksums beyond DE (FR/IT/ES/... currently fall to the
  generic 8-12 shape check). Add per-jurisdiction rules only when a consumer
  bills that jurisdiction — each rule needs pinned official vectors.
- TaxNumber jurisdiction rules beyond DE's 10-13 digits (UK UTR, PL NIP-linked
  form) — same demand gate.

## Open questions

- Should `RegisterRef.Registry` ever grow a controlled vocabulary, or does
  opacity remain the contract? (Current decision: opaque by design —
  `identifiers.go:63`. Revisit only if two consumers disagree on registry
  codes.)
