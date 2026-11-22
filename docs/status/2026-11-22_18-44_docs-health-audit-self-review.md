# Status Report: docs-health AUDIT Phase + Brutal Self-Review

**Created:** 2026-11-22 18:44 CET · agent (glm-5.3, Crush) · session leg: the full docs-health AUDIT commanded after report 1 (`docs/status/2026-11-22_16-24_agents-md-doc-corpus-gate-green.md`)

**Scope line.** VERIFIED this leg (with commands): all seven living docs + SECURITY.md read and checked against source (`legalform.go`, `identifiers.go`, `country.go`, `nip.go`); every TODO_LIST evidence ref spot-checked (`identifiers.go:63/74/265`, `legalform.go:58`); README code examples EXECUTED via throwaway test (7/7 pass, exact canonical outputs, test deleted after); `go test ./...` pass; `docs-gates run .` 5/5 PASS; `buildflow --build-mode dev` exit 0; AGENTS.md temporal-pollution grep (guide pattern) clean; AGENTS.md size 6,940 B (budget: sweet spot). NOT verified: `go vet` not run standalone this leg (buildflow's golangci covers govet); "live GLEIF records" still inherited-from-comments; pkg.go.dev rendering (no release cut); the health report's as-found scores are my own classification (no prior baseline exists — first audit).

## What did you forget? / What could you have done better? (the honest answers)

1. **Report 1's scope line overclaimed verification.** It said README "table claims match code (verified against source earlier)" — FALSE. The FI fixed-length, branch/trust legal-form classes, and "PL NIP-linked" rows were never checked row-by-row; the audit found all three wrong. A scope line that claims verification that didn't happen is the worst failure available in this skill's grammar — it converts unknown risk into false assurance.
2. **Shipped README examples without executing them.** I mentally walked `ParseTaxNumber("DE", "217/8150/8173")` and `NewRegisterRef(...)` and moved on. The verify-checklist's own bar ("walk through each step mentally") was met — and the mental model was right (executed 7/7 this leg, after the fact) — but an executable check was 30 seconds away and I declined it. Fixed now; the pattern (execute, don't simulate) is the lesson.
3. **A careless multiedit construct failed wholesale** on SECURITY.md (fourth edit carried a bad old_string shape; atomic abort). Recovered with four single edits — but the failure was manufactured by me rushing the batch.
4. **Report 1 is now partially resolved and carries no annotations.** Its §d items (AGENTS pollution, split brain, skill-order) all have verdicts now; my own footer says "when superseded, annotate — never rewrite." I neither annotated nor explicitly decided not to.

## a) FULLY DONE

1. **Full docs-health AUDIT executed per skill**: verify-checklist, agents-quality-guide (rubric + pruning), health-report-format all read BEFORE acting (the process failure of the first leg, corrected).
2. **SECURITY.md Critical fix**: vulnerability-report URL pointed at `go-must` — misdirecting security reports for this library. Now go-legalparty. Plus: v0.1.2→v0.1.0, false CI-scanning claim made honest, footer updated.
3. **README factual drift fixed (3 rows, pre-existing)**: LegalFormClass "branch/trust" → public_body/non_profit; VATID "AT/FI" → AT/PL; TaxNumber "PL NIP-linked" → DE-only rules, honest generic.
4. **AGENTS.md hardened per rubric**: temporal qualifiers stripped (grep clean), dprint remote-plugin gotcha + internal/ deliberate non-fix added, consumers section now current truth. 6.9 KB.
5. **Migration narrative single-homed**: README/STACK/AGENTS state current truth; CHANGELOG owns the history.
6. **README examples extended and now EXECUTED**: TaxNumber + RegisterRef added; 7/7 verified with exact canonical outputs.
7. **5 TODO rows shipped to CHANGELOG `[Unreleased]`** (AGENTS hardening, narrative single-homing, README examples + drift fixes, CHANGELOG link, SECURITY/AUTHORS verified).
8. **Health report delivered inline with visible math** (Accuracy 6.25 / Fitness 7.75, as-found, first audit — no invented baseline).
9. End state: `go test` pass, docs-gates 5/5, buildflow dev exit 0, 9 open TODO rows, all evidence refs valid.

## b) PARTIALLY DONE

1. **Report 1 annotation**: resolved §d/§b items still read as open in the timestamped file (see #4 above). The gates accept it; my own footer grammar doesn't.
2. **AGENTS.md documentation-model section** lists the seven docs but omits the `docs/status/` snapshot convention — one line missing.

## c) NOT STARTED

Unchanged from report 1, all still rowed: REGON, test vectors, CI wiring, fuzz targets, `--build-mode full`, dprint local-plugin decision, invoices `// indirect` (repo: invoices), legalparty-model leftover test (repo: cqrs-htmx), release decision. Plus new: canonical-facts drift gate (see f), census routing (see f).

## d) TOTALLY FUCKED UP

1. **The scope-line overclaim (see honest answers #1)** — false assurance shipped in a report whose entire job is calibrated assurance.
2. **Examples shipped unexecuted (see #2)** — caught and closed this leg, but only because the self-review demanded it.
3. The manufactured multiedit failure (see #3).

## e) WHAT WE SHOULD IMPROVE

1. **Never claim verification in a scope line that a command didn't produce.** The FI/branch-trust/PL drift survived one full session because a scope line said it was checked.
2. **Execute, don't simulate** — any doc-shipped code example gets run, not walked.
3. **Annotate resolved items in prior reports** in the same change that resolves them (or record the deliberate non-annotation).
4. **A standing drift gate** for README-vs-code facts (the fleet's `check-canonical-facts.sh` pattern) would have caught the FI/branch-trust class mechanically. Manual per-audit verification is how it survived this long.
5. Report 1's f/23 (domain-language census) is routed nowhere — it lives only in the timestamped file. Fleet-level items need a fleet-level home or an explicit drop decision.

## f) Next things

| #  | Task                                                                                             | Effort | Status in TODO_LIST  |
| -- | ------------------------------------------------------------------------------------------------ | ------ | -------------------- |
| 1  | REGON 9+14 digit + GUS vectors; then delete company-master-data's local copy                     | M      | rowed                |
| 2  | Vectors: ParseTaxNumber / RegisterRef / RequiresRegistryAnchor / AllowsLEI; flip 3 FEATURES rows | S      | rowed                |
| 3  | CI decision + wiring (zero workflows; buildflow gate candidate)                                  | M      | rowed                |
| 4  | Fix invoices' stale `// indirect` (repo: invoices)                                               | S      | rowed                |
| 5  | Release decision: v0.1.1 (SECURITY fix argues for prompt shipping) vs v0.2.0 with REGON+vectors  | S      | rowed (question g/1) |
| 6  | Fuzz targets for the four parsers                                                                | M      | rowed                |
| 7  | `buildflow --build-mode full` once                                                               | M      | rowed                |
| 8  | dprint local-plugin decision                                                                     | S      | rowed                |
| 9  | legalparty-model identifiers_test.go: aliases or dead? (repo: cqrs-htmx)                         | S      | rowed                |
| 10 | Annotate report 1's resolved §b/§d items inline (same-session verdicts available)                | S      | NEW — rowing now     |
| 11 | Add `docs/status/` convention line to AGENTS.md documentation model                              | S      | NEW — rowing now     |
| 12 | Canonical-facts drift gate (pin README table rows to code symbols; fleet script pattern)         | M      | NEW — rowing now     |
| 13 | Route the census item: ROADMAP open question or explicit drop                                    | S      | NEW — rowing now     |
| 14 | pkg.go.dev render check after next release                                                       | S      | rides #5             |
| 15 | Triage first Dependabot wave when it arrives                                                     | S      | recurring, unrowed   |

## g) Questions I cannot answer myself

1. **Release:** the SECURITY.md fix (misdirected advisory URL) argues for shipping v0.1.1 promptly rather than holding for v0.2.0/REGON — cut v0.1.1 now, or hold?
2. **CI:** wire buildflow as the push/PR gate — copy a fleet-canonical workflow file from a sibling (which?), or compose fresh?
3. **Drift gate:** do you want the standing canonical-facts script (README-table-rows ↔ code symbols, fails CI on drift), or is per-audit manual verification the chosen tradeoff for a library this small?
