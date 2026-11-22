# TODO List

> Short-term, actionable, bounded work items, verified against the actual
> code (and, for migration rows, against the consumer repos). When an item
> ships, remove it here and record it in `CHANGELOG.md` under the version it
> shipped in. For long-term and demand-gated ideas, see `ROADMAP.md`.

**Created:** 2026-11-22 · docs-health BUILD pass (agent, glm-5.3) · backlog
basis: code sweep + consumer-repo verification (`cqrs-htmx/legalparty-model`,
`invoices`, `company-master-data`) + session report
`docs/status/2026-11-22_16-24_agents-md-doc-corpus-gate-green.md` §f (rows
marked R§f/N derive from it).

**Last updated:** 2026-11-22 · execution-session pass — REGON row narrowed
(the type shipped in v0.2.0; only the company-master-data adoption half
remains), fuzz-target row removed (fuzz_test.go shipped inside merged PR #1,
recorded in CHANGELOG v0.2.0), census item routed to ROADMAP "Explicitly out
of scope", ROADMAP gained PL-VATID / GLEIF-cadence / rule-seam /
modularize-revisit rows (plan v2 M7).

## High Impact

| Task                                                                                                                                                               | Status    | Impact | Effort | Evidence                                                                                                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | --------- | ------ | ------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Delete company-master-data's local REGON copy (the type itself shipped in v0.2.0; adoption remains) (R§f/8-10)                                                    | 🔴 `TODO` | High   | S      | `company-master-data/internal/domain/checksum.go:9` keeps REGON local "NOT yet covered by go-legalparty". Bump it to v0.2.0, delegate to `legalparty.ParseRegon`, delete the local checksum + dead helpers |

## Medium Impact

| Task                                                                                                         | Status    | Impact | Effort | Evidence                                                                                                                |
| ------------------------------------------------------------------------------------------------------------ | --------- | ------ | ------ | ----------------------------------------------------------------------------------------------------------------------- |
| Fix stale `// indirect` marker on go-legalparty in invoices' go.mod (`go mod tidy`) (R§f/11, repo: invoices) | 🔴 `TODO` | Med    | S      | `invoices/internal/models/tax_id_validation.go:8` imports it directly while go.mod marks it indirect                    |
| Canonical-facts drift gate: pin README table rows to code symbols (fleet `check-canonical-facts.sh` pattern) | 🔴 `TODO` | Med    | M      | README drifted 3 rows (FI length, branch/trust classes, PL rule) through a full session + a false "verified" scope line |

## Low Impact

| Task                                                                                                                      | Status    | Impact | Effort | Evidence                                                                   |
| ------------------------------------------------------------------------------------------------------------------------- | --------- | ------ | ------ | -------------------------------------------------------------------------- |
| Run `buildflow --build-mode full` once (only dev mode has been exercised) (R§f/13)                                        | 🔴 `TODO` | Low    | M      | Session runs used `--build-mode dev` only                                  |
| dprint local-plugin decision: accept network fetch vs local wasm paths (R§f/22)                                           | 🔴 `TODO` | Low    | S      | buildflow preflight WARN names both options on every run                   |
| Check legalparty-model's leftover identifiers_test.go: tests aliases or dead code? (R§f/24, repo: cqrs-htmx)              | 🔴 `TODO` | Low    | S      | File present while the value sources it tested were removed                |
| Annotate report 1's resolved §b/§d items inline (AGENTS pollution, split brain, skill-order — verdicts exist in report 2) | 🔴 `TODO` | Low    | S      | Report 1 §b/§d read as open; fixes shipped same session                    |
| Add `docs/status/` snapshot convention line to AGENTS.md documentation model                                              | 🔴 `TODO` | Low    | S      | AGENTS.md documents seven files but not the status-report convention       |
| Route the fleet census item: ROADMAP open question or explicit drop (report 1 §f/23 lives nowhere)                        | 🔴 `TODO` | Low    | S      | Report 1 f/23 unrowed; routing gate passed but the item has no living home |
