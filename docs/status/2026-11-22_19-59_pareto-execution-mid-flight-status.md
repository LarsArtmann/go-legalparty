# Status Report — Pareto Execution, Mid-Flight (19:59 CET)

**Session:** 2026-11-22 ~19:10–19:59 CET · agent glm-5.3, Crush · executing the
user's "WHOLE TODO LIST" command over
`docs/planning/2026-11-22_19-05_pareto-execution-plan.md` (v1), realigned to
`docs/planning/2026-11-22_19-38_pareto-execution-plan-v2-verified-backlog.md`
when a parallel session superseded v1 mid-flight.

## Scope line

**Verified here (fresh commands at 19:55–19:59 unless timestamped):**
`go test -count=1 ./...` ok (go-legalparty); `golangci-lint run --timeout=5m`
→ "0 issues"; `dprint check` → **FAILS** on
`docs/planning/2026-11-22_19-38_pareto-execution-plan-v2-verified-backlog.md`
(unformatted table — reproduction below); master CI run `45057349035`
(push, merge commit `f402ee1`) = **failure, docs-format job, exit 20**; PR-branch
CI run `45055747829` = all 5 jobs green; legalparty-model `go test ./...`
3 pkgs ok and `TestLegalFormClass_Predicate` PASS with `-count=1 -run` (19:58);
invoices `go test ./...` 47 pkgs ok (19:34, pre-commit gate also green);
fuzz: 5 targets × 15 s PASS (19:26–19:33); `docs-gates run .` 5/5 PASS (19:21);
proxy: `go list -m -versions` shows `v0.1.0 v0.1.1`; scratch-module
`go get @v0.1.1` + 7-parse smoke PASS (19:17).

**NOT verified here:** company-master-data's suite against library HEAD (not
run — its adoption is gated behind the v0.2.0 release); pkg.go.dev rendering
of v0.1.1 (license-suppressed by design, `@v0.1.1` fetch returned 404 while
re-indexing); master CI for any state after run `45057349035`; the
parallel session's current intent (I read its committed artifacts only).

---

## a) FULLY DONE

1. **v0.1.1 released and consumable** — CHANGELOG cut `7a95a2a`, annotated
   tag pushed, proxy lists it, scratch consumer resolves + compiles + all 7
   parse paths smoke-tested. Ships the corrected SECURITY.md advisory URL,
   truthful README rows, and (from the parallel session, adopted in
   `0afeb02`) the AllowsLEI + Dutch-VAT rule fixes.
2. **TaxNumber + RegisterRef + rule vectors** — `0e57147`:
   `TestParseTaxNumber` (18 vectors: all three `tax_number.*` codes, DE 10/13
   bounds, generic 5/16 bounds, separators, VAT-shape refusal incl.
   lowercase-prefix + country-independent), `TestNewRegisterRef` (both
   `register.*` codes, trim, zero-value Validate ordering). FEATURES.md now
   shows **0 `PARTIALLY_FUNCTIONAL` rows** (8 test functions).
3. **REGON type + checksum + vectors** — `32a46b8`: `regon.go` (9-digit
   weights 8-9-2-3-4-5-6-7 mod 11 over the parent prefix; remainder 10
   encodes check digit 0 — the anti-NIP edge; 14-digit serial deliberately
   unchecked, KRS-falsification evidence carried in the doc comment) +
   `TestParseRegon` (15 vectors: algorytm.org official, real KRS odpis
   records, remainder-10 both directions, 14-digit parent checks). Passed
   first run, zero expectation fixes.
4. **REGON docs (library half)** — `17b60c5`: README example + table row,
   DOMAIN_LANGUAGE entry + canonical-form note, AGENTS regulation edge +
   file inventory, FEATURES row. Also removed README's false MIT badge
   (LICENSE is proprietary).
5. **CI exists and gates PRs green** — merged via PR #1 (`f402ee1`):
   `.github/workflows/ci.yml` with test (race, vet, govulncheck), lint
   (golangci-lint v2.15.0), consumer-import simulation, docs-format (pinned
   dprint 0.64.3), cue workflow-policy vet — all actions SHA-pinned,
   actionlint-clean, policy-vet green. First-run triage executed:
   govulncheck v1.1.4 panicked on Go 1.27 syntax → bumped to v1.10.0
   (`3f4bd18`, version verified via `go list -m -versions` + clean local
   scan). SECURITY.md's "CI scanning planned" claim is now true and
   specific; version row bumped to v0.1.1.
6. **Lint is CI-grade zero** — `f3cc1c1`: checksum spec constants named
   (`leiLength`, `leiModulus`, `iso7064*`, `nipLength`, `regonModulus`,
   `regonZeroRemainder`, `minVATPrefixedLength`, `deVATLength` — consumer
   checksum.go naming precedent), country-table globals carry explained
   `//nolint:gochecknoglobals` directives. `golangci-lint run` exits 0 with
   zero findings (verified 19:59).
7. **Fuzz targets** — shipped inside `f3cc1c1`: 5 targets (NIP, VATID, LEI,
   TaxNumber, Regon) pinning crash-safety, re-parse idempotence of accepted
   values, and the frozen `legalparty.*` rejection-code namespace on every
   rejection. 5 × 15 s runs clean; no panics, no idempotence breaks.
8. **Consumer hygiene (v1-M11)** — invoices: bumped to v0.1.1, stale
   `// indirect` marker gone, vendor resynced (`go mod vendor`), 47 test
   packages green, committed `44ac61f3` (amended over a daemon race commit,
   provenance noted). legalparty-model: bumped to v0.1.1, 3 packages green,
   committed `1af150a4` (same race handling).
9. **Leftover-test verdict (F46)** — legalparty-model's
   `identifiers_test.go` is **KEEP**: it pins the alias bridge surface and
   carries vectors the library lacks (IT codice fiscale, ELSTER 13-digit
   form). Re-verified green post-bump.
10. **v2-M1 consumer verification (2 of 3 repos + verdicts)** — invoices and
    legalparty-model suites green against v0.1.1; invalid-class rule vectors
    added to legalparty-model (`TestLegalFormClass_Predicate` extended,
    PASS, committed by daemon as `46d061b4`). invoices NL-fixture task
    verified **N/A**: invoices has no NL VAT path (PL NIP + DE VAT only,
    `tax_id_validation.go:61-81`), so the Dutch tightening is unreachable
    there today.
11. **Parallel-session adoption** — its refactor/fix work committed with
    attribution (`0afeb02`), its dependabot github-actions ecosystem adopted
    (`10f55df`), its plan v2 read in full and used to re-rank this session.

## b) PARTIALLY DONE

1. **v0.2.0 release (v2-M2)** — everything pre-release is green locally
   (REGON in `[Unreleased]`, tests/race/lint/docs-gates). **Blocked by red
   master** (see d/1) and not yet stamped/tagged. Effort to finish: S
   (~30 min incl. verification).
2. **company-master-data adoption (v2-M3 / v1-M6)** — not started; depends
   on v0.2.0 being fetchable. The swap plan is fully researched (local copy
   at `internal/domain/checksum.go`, vectors in its checksum_test.go,
   property tests reference its helpers). Effort: M.
3. **v1-M5 consumer-doc half** — README provenance line, AGENTS consumer
   item 3, and the TODO_LIST REGON-row removal are deliberately deferred to
   the adoption commit so every commit stays truthful (library docs landed;
   consumer-migration sentences have not).
4. **TODO_LIST truth (v2-M7 half)** — my three row deletions shipped
   (vectors, release-decision, CI); the REGON and fuzz rows remain stale
   (plan v2's own errata flags them; its "one deliberate edit" rule awaits
   that task).

## c) NOT STARTED

All from plan v2, none blocked except as noted: v2-M5 (AT `U`-strip doc
comment; RegisterRef trim-semantics decision + vectors); v2-M6
(canonical-facts drift gate → CI); v2-M8 (docs honesty batch: report-1
§b/§d annotations, series index for the 4 HTML reports, AGENTS reconcile
with CI/regon/fuzz + `docs/status/` convention line); v2-M9 (file:line
citation refresh across the 4 HTML reports); v2-M10 IsValid surface (gate
D1); v2-M12 buildflow `--build-mode full`; v2-M13 dprint local-plugin
decision; v2-M15 host hygiene (1.10 GB buildflow DB VACUUM, binary rebuild
`be700ae`→HEAD, PATH precedence check); v2-M19 Dependabot/govulncheck waves
(2 Dependabot runs were already queuing on master at 18:50 — the new
github-actions ecosystem is live).

## d) TOTALLY FUCKED UP

1. **MASTER CI IS RED.** Run `45057349035` (push on merge `f402ee1`):
   docs-format job exit 20. Root cause reproduced locally: the parallel
   session committed plan-v2 markdown (`75eda51`) straight to master with
   unformatted table alignment — its commit bypassed the formatting lane my
   PR had passed. Severity: **blocks the v0.2.0 release** (go-release rule:
   never tag while master CI is red — a red tag-run is permanent).
   Mitigation ready and trivial: `buildflow -s dprint-format --fix` +
   commit + push (~2 min). Not applied: this report-and-wait order arrived
   first.
2. **Three auto-commit daemon races** — the daemon committed my staged
   changes mid-pre-commit-hook in invoices (`6938bce0`) and legalparty-model
   (`ddebf09b`, `46d061b4`), twice producing a failed `git commit` (HEAD
   lock) and mislabeled history. Recovered by amending the daemon commits
   (they were local-only and carried my content). Residual risk: any
   concurrent session can hit this at any time; nothing structural stops it.
3. **Sed-edited Go code (my own mistake)** — the first fuzz draft was
   mangled by a `sed` import insertion plus a wrong prefix slice
   (`[:12]` vs the 11-char `"legalparty."`), caught by the local test run
   before any commit. Should have used the edit tool; the failure was
   cheap only because tests ran first.
4. **govulncheck pinned from memory** — v1.1.4 (stale recall) panicked in
   CI on Go 1.27 syntax; fixed to v1.10.0 after querying versions. The
   verify-before-pin discipline I apply to external claims, I skipped for
   my own workflow pin.
5. **Local toolchain carries a known vuln** — Go 1.27.3 locally is affected
   by GO-2026-6859 (net/http MaxBytesReader race, fixed in 1.27.4); CI
   runs 1.27.4 so it is clean there. Local `govulncheck` exits 3. Host
   toolchain bump pending (v2-M15 adjacent).

## e) WHAT WE SHOULD IMPROVE

1. **Direct-to-master pushes bypass every local gate** — plan-v2's author
   committed without buildflow/pre-commit. Now that CI exists, enable
   branch protection on master (require PR + green CI), or make
   plain-commits-to-master a stated no-go in AGENTS.md.
2. **Multi-session + daemon concurrency has a real error bill** — 3 races,
   1 red master, several re-read-before-edit saves this session alone. The
   plan-errata coordination convention worked; consider also a "claiming
   task X" line in the plan file before starting a medium task.
3. **The dprint/CI defense-in-depth worked exactly as designed** — master
   red within minutes of the unformatted merge. Keep docs-format in CI
   forever; it paid for itself on run #1.
4. **Pin tool versions from query output, not recall** — the govulncheck
   panic cost a CI round-trip. Workflow edits should cite the version-list
   command in the commit message (I did, after the fact).
5. **LICENSE-vs-badge class of drift** — an MIT badge sat over a
   proprietary LICENSE until today. The planned canonical-facts gate (v2-M6)
   should include the LICENSE/badge pair in its pinned facts.
6. **pkg.go.dev shows nothing for this vocabulary library** (license-gated
   docs). Deliberate fleet pattern, but for a library whose product IS its
   public types, a rendered-docs mirror (or MIT for this one repo) is worth
   a conscious decision rather than drift-by-default.

## f) Next tasks (ranked; feeds HARVEST)

| #  | Task                                                                                                                                                                                                                   | Impact   | Effort | Category      |
| -- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- | ------ | ------------- |
| 1  | Fix master: `buildflow -s dprint-format --fix`, commit, push; verify CI green (run for the fix)                                                                                                                        | Critical | S      | Bug           |
| 2  | Release v0.2.0: CHANGELOG stamp, green battery (`go test -race`, vet, buildflow dev, docs-gates), annotated tag, push, proxy + scratch `go get` + parse smoke, pkg.go.dev fetch                                        | Critical | S      | Release       |
| 3  | GitHub Releases for v0.1.1 + v0.2.0 (`--prerelease`, go-release Phase 7)                                                                                                                                               | Low      | S      | Release       |
| 4  | company-master-data: bump @v0.2.0, delegate ValidREGON to `legalparty.ParseRegon`, delete the local REGON checksum + dead helpers, run suite, update its AGENTS/STACK/CHANGELOG                                        | High     | M      | Feature       |
| 5  | Consumer-doc half of REGON: README provenance line, AGENTS consumer item 3, TODO REGON row removal → CHANGELOG (rides the adoption commit)                                                                             | Med      | S      | Documentation |
| 6  | Canonical-facts drift gate: `scripts/check-canonical-facts.sh` pinning README/FEATURES rows ↔ code symbols (weights, remainder edges, class list, AT/PL/NL lengths, LICENSE/badge pair), CI job, seeded-drift red demo | Med      | M      | Quality       |
| 7  | TODO_LIST: remove stale REGON + fuzz rows; record in CHANGELOG (v2-M7 half)                                                                                                                                            | Med      | S      | Cleanup       |
| 8  | AGENTS.md reconcile: CI workflow section, fuzz pattern line, `docs/status/` convention line                                                                                                                            | Med      | S      | Documentation |
| 9  | AT `U`-strip doc comment in ParseVATID with the AT UID regulation source (v2-M5a)                                                                                                                                      | Med      | S      | Documentation |
| 10 | RegisterRef trim-semantics decision (NewRegisterRef trims, raw-struct Validate doesn't — pin both) + vectors (v2-M5b)                                                                                                  | Med      | S      | Feature       |
| 11 | ROADMAP routing batch: jurisdiction-rule seam, modularize-verdict revisit, PL VATID checksum, GLEIF vector cadence, census item (v2-M7b)                                                                               | Med      | S      | Documentation |
| 12 | Annotate report 1 (`2026-11-22_16-24`) resolved §b/§d items inline (v2-M8a)                                                                                                                                            | Low      | S      | Documentation |
| 13 | Series index for the 4 HTML reports + 2 D2/SVG pairs under docs/architecture-understanding/ (v2-M8b) — DONE 2026-11-22: `docs/architecture-understanding/README.md`                                                    | Low      | S      | Documentation |
| 14 | Citation refresh: re-verify file:line refs in the 4 HTML reports against HEAD (v2-M9)                                                                                                                                  | Low      | M      | Documentation |
| 15 | IsValid surface (gate D1): `LEI/VATID/NIP/Regon.IsValid` + `TaxNumber.IsValidFor(cc)` + vectors + README/DOMAIN_LANGUAGE/FEATURES (v2-M10)                                                                             | Med      | M      | Feature       |
| 16 | buildflow `--build-mode full` + triage findings or record deliberate non-fixes (v2-M12)                                                                                                                                | Med      | M      | Quality       |
| 17 | dprint plugin decision: `BUILDFLOW_DPRINT_PLUGINS` local wasm vs accept fetch; apply or document (v2-M13)                                                                                                              | Low      | S      | Quality       |
| 18 | Host hygiene: `sqlite3 …/buildflow.db VACUUM` (1.10 GB), rebuild buildflow (`be700ae` → HEAD), PATH check; bump local Go toolchain past GO-2026-6859 (v2-M15)                                                          | Low      | S      | Cleanup       |
| 19 | Dependabot wave triage: gomod + the new github-actions PRs (v2-M19) — DONE 2026-11-22: zero open PRs, zero open security alerts                                                                                        | Med      | S      | Hygiene       |
| 20 | legalparty-model: re-run aggregate suite against v0.2.0 after release; push its master (2 local commits incl. the daemon's fixture commit)                                                                             | Med      | S      | Hygiene       |
| 21 | errata in plan v2: F02 N/A verdict (no NL path in invoices), master-red incident, this report's pointer                                                                                                                | Low      | S      | Documentation |
| 22 | Decide NL-VAT delegation for invoices (route NL through legalparty or record the N/A permanently)                                                                                                                      | Low      | S      | Feature       |
| 23 | Fuzz CI posture decision: keep seed-corpus-only in CI (current) vs add 30 s `-fuzztime` smoke job                                                                                                                      | Low      | S      | Quality       |
| 24 | Branch protection / AGENTS rule: no direct-to-master pushes now that CI gates exist                                                                                                                                    | Med      | S      | Process       |
| 25 | Rendered-docs decision for pkg.go.dev (license change vs docs mirror vs accept suppression)                                                                                                                            | Low      | S      | Process       |

## g) Questions I cannot answer myself

1. **Red-master ownership:** the failing file belongs to the parallel
   session (plan v2). I have the two-minute fix ready (`dprint --fix` +
   push). Fix it myself as part of "get shit done", or leave it for that
   session to own its file?
2. **Push policy for consumer repos:** invoices' master is already pushed
   (in sync); cqrs-htmx carries 2 local commits (my v0.1.1 bump + the
   daemon's fixture commit). May I push cqrs-htmx's master, or do you
   handle fleet-repo pushes yourself?
3. **Gate D1 (IsValid surface):** implement it now under the "whole list"
   command, or hold it for the deepening-brief design pass plan v2 gated
   it on? Both are defensible; it is your call.

---

**Provenance:** written by the executing session at 19:59 CET per the
status-report skill; no commit made (harness contract — the auto-commit
daemon picks this file up). HARVEST note: section (f) is routing-ready;
items 1–5 are already TODO_LIST-grade, 6+ split between TODO_LIST and
ROADMAP per the demand-gating already written into plan v2.
