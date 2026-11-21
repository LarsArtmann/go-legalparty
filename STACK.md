# STACK.md — go-legalparty

Decision record per the stack-selection skill.

## Profile

| Axis | Answer |
| --- | --- |
| Shape | Pure Go library (no CLI/service) |
| Domain | Legal-party identifier validation (DE/PL/EU) |
| Persistence | None |
| Interfaces | Go API only |
| Constraints | Minimal deps; extraction per ADR-0001 with migration rows |

## Chosen foundation

| Concern | Choice | Why this, not that | Owning skill | Revisit trigger |
| --- | --- | --- | --- | --- |
| Error taxonomy | go-error-family v0.12 | Fleet standard; stable machine codes are the point (dashboards key on them) | how-to-golang | — |
| Everything else | stdlib | checksums + tables need nothing else | how-to-golang | a GLEIF/registry IO client grows here (gleif/ stays in legalparty-model until a second consumer) |

## Deliberately excluded

| Candidate | Reason |
| --- | --- |
| full extraction of legalparty-model | the decider/fold/events layers are cqrs-htmx concerns; extracting them couples a value library to go-cqrs-lite |
| golang "validators" libs | generic; no NIP/HRB/KRS knowledge, no error-family taxonomy |

## Wiring order

1. `meta set` registration (2026-11-21)
2. Migration rows: cqrs-htmx/legalparty-model consumes this; invoices replaces tax_id_validation.go (fixing its remainder-10 NIP bug); company-master-data adopts (ADR-0001 rule 5)
