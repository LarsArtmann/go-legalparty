# Status Report — Pareto Plan v2 Execution (mid-flight, tier 1–20% shipped)

**Written:** 2026-11-23 01:12 CET · execution session (agent, glm-5.3-flash, Crush)
**Scope:** executing `docs/planning/2026-11-22_19-38_pareto-execution-plan-v2-verified-backlog.md`
(owner order: "NOW GET SHIT DONE — the WHOLE TODO LIST") from a clean tree @
`f402ee1` (PR #1 merge, which stranded the five-skill review artifacts).
**Position at writing:** 8 of 13 medium tasks fully shipped, v0.2.0 released,
M10 mid-implementation, **3 commits on local master are unpushed**, M10 work
is **uncommitted**, and the file-size gate just forced a vatid.go extraction
that invalidates part of the M9 citation refresh just completed.

---

## a) FULLY DONE (shipped, verified, committed)

1. **Five-skill review artifacts rescued to master** (`98d9207`, cherry-pick).
   Discovery: PR #1 merged from an older branch tip, so `78ea3d7` (type-model
   review, modularization verdict, five-skill status report) never reached
   master — verified via merge-parent inspection (`f402ee1`'s second parent is
   `3f4bd18`, not `1a37ed2`). The stray plan duplicate `1a37ed2` was
   deliberately skipped (content already on master via `75eda51`). The
   AGENTS.md PR-branch-ordering lesson is now written into repo conventions.
2. **M1 consumer verification sweep** — all three fleet consumers
   race-tested against this HEAD via isolated `go.work` / GOWORK files
   (no consumer go.mod was touched): invoices GREEN, cqrs-htmx/legalparty-model
   GREEN (aliases compile), company-master-data GREEN. The v0.1.1 tightenings
   (NL strict shape, AllowsLEI deny-on-invalid) break nothing.
3. **F02 invoices NL fixtures** — new `TestDutchVATIDShapeValidation` (4
   vectors incl. letters-in-digits rejection), landed in invoices via daemon
   commit then message-amended, **pushed** (`51177112`).
4. **F04 legalparty-model invalid-class predicates** — already covered by a
   concurrent session's `TestLegalFormClass_Predicate` (both zero-value and
   unknown-class vectors); verified, no action needed.
5. **M2 release v0.2.0** — CHANGELOG stamped (`ee39135`), code battery green,
   annotated tag pushed, **scratch-module `go get @v0.2.0` verified** (NIP +
   Regon parse green through the proxy). pkg.go.dev render check: page NOT
   yet indexed at check time (proxy is faster than the crawler; re-check
   pending). v0.1.1's page verified live — with the honest finding that godoc
   is suppressed by pkg.go.dev license policy (proprietary LICENSE = no
   redistributable docs). Recorded as annotation, not a packaging bug.
6. **M3 company-master-data REGON adoption** (`e61ec9f`, local, unpushed) —
   bumped to v0.2.0, `ValidREGON` now delegates to `ParseRegon` keeping the
   bare-digits form contract, local modulo-11 implementation deleted, full
   race suite green (property tests' independent oracle confirms semantic
   equivalence), CHANGELOG entry added. ADR-0001 single-source adoption is
   now complete fleet-wide: no identifier checksum remains local anywhere.
7. **M4 invoices hygiene** — `go mod tidy` was a NO-OP: the concurrent
   `44ac61f3` bump already removed the stale `// indirect` marker. My NL
   fixture test is the session's invoices payload.
8. **M5 contract pinning** (`c253fb0`) — AT `U`-strip regulation comment
   (ATU + 8 digits, U is display format), RegisterRef split contract stated
   and pinned by a new vector: constructor trims, Validate is presence-only
   and never mutates.
9. **M6 canonical-facts drift gate** (`c1d9381`) —
   `scripts/check-canonical-facts.sh`: 26 checks pinning README/FEATURES
   validation rows to code anchors (NIP/REGON weights + remainder edges, DE
   ISO 7064, AT/PL lengths, NL shape, five classes, assigned-only countries,
   DE-only TaxNumber rule) plus absence checks for the three burned drifts
   and false status rows. Wired as its own CI job (actionlint green).
   **Proven both ways**: green on clean tree, red on a seeded FI claim,
   green after revert. AGENTS.md Commands section documents it.
10. **M7 backlog truth** (in `ee39135`) — TODO_LIST REGON row narrowed to
    the adoption half (then shipped by M3; row removal pending final
    reconciliation), fuzz row removed + recorded in CHANGELOG, ROADMAP
    gained PL-VATID / GLEIF-cadence / rule-seam / modularize-revisit themes
    and an explicit out-of-scope census drop.
11. **M8 docs honesty** (in `ee39135`) — AGENTS.md snapshot convention +
    cadence decision for the three review-artifact dirs, CI/fuzz
    reconciliation, series index `docs/architecture-understanding/README.md`
    cross-linking 4 reports + 2 diagram pairs, inline annotations on three
    status reports (9 audit-all items resolved with done-at verdicts).
12. **M19 Dependabot triage** — zero open PRs, zero open security alerts.
13. **Splice-script persistence** (19-26 §f/28) — original /tmp script was
    already lost; recreated as `html-report-kit/assets/splice_report.py`
    (anchor-span, fails loudly on non-unique anchors), smoke-tested.
14. **F54 parked triggers** — M16/M17/M18 demand triggers all routed to
    ROADMAP (verified present in `ee39135`).

## b) PARTIALLY DONE (in flight, honest state)

1. **M10 IsValid surface — code complete, gate-blocked, UNCOMMITTED.** All
   five methods shipped in the working tree: `LEI.IsValid`, `VATID.IsValid`,
   `NIP.IsValid`, `Regon.IsValid`, `TaxNumber.IsValidFor(cc)` (parse-and-
   discard, zero value never valid, TaxNumber deliberately country-dependent).
   Vector tests added (`TestIsValidSurface`, `TestRegonIsValid`,
   `TestTaxNumberIsValidFor`) and race-green. Docs rows added (README idiom
   note, FEATURES row, DOMAIN_LANGUAGE concept, CHANGELOG `[Unreleased]`).
   **Blocked on:** the 350-line file-size gate — `identifiers.go` is 365 and
   `identifiers_test.go` is 367. The vatid.go extraction (block spans
   lines 206–end + `TestParseVATID`) was scoped (symbol map taken) but NOT
   executed when this report was written.
2. **M2's pkg.go.dev render check** — `go get` half verified; the crawler
   half returns 404 at check time. Expected indexing lag; re-check queued
   in the final phase, plus the two audit-all annotations that depend on it.
3. **M9 citation refresh** (`26ca5ed`) — 15 file:line refs re-verified
   against HEAD and fixed across 3 HTML reports + FEATURES. **Immediately
   half-obsolete**: the forced vatid.go extraction will move the VATID block
   again (~6 of the 15 refs). The refresh was correct for the tree it ran
   on; it will need a second pass post-extraction. Sequencing miss (see §d).
4. **buildflow --fix formatting wave** — the sanctioned repair pass
   re-aligned the markdown tables my M7/M10 doc edits had widened across
   TODO_LIST, planning file, status reports, series index, script header.
   These formatting deltas are uncommitted and mixed into the M10 working
   tree; the tree is green except the file-size gate (2 error findings).
5. **M11's aggregate half** — legalparty-model already race-tested against
   HEAD in M1; the leftover `identifiers_test.go` verdict (aliases vs dead
   code) and the formal v0.2.0-bump test pass are NOT done.

## c) NOT STARTED

1. **vatid.go extraction** (M10 completion prerequisite) + the follow-on
   AGENTS.md file-list row, drift-gate anchor moves (4 VAT checks point at
   identifiers.go), and the citation re-pass (§b3).
2. **M11** leftover-test verdict in legalparty-model (F44–F45).
3. **M12** `buildflow --build-mode full` + triage (F46–F47).
4. **M13** dprint plugin decision (F48–F49) — the preflight WARN still fires
   on every run.
5. **M14** lint leftover policy (F50–F51) — gochecknoglobals/wsl per D3 rec.
6. **M15** host hygiene (F52–F53) — state DB 1.07 GB, binary staleness
   warning now live (`769b3a4` vs `4faa6ad` in the BuildFlow repo).
7. **Final reconciliation** — M10 commit, TODO_LIST REGON-row removal
   (M3 shipped it), CHANGELOG v0.2.0-era entries, full battery, **push of
   the 3 unpushed master commits** (`c253fb0`, `c1d9381`, `26ca5ed`) + the
   M10 commit, CI-green verification on push.
8. **Foreign-file disposition** — the 19-59 concurrent-session status report
   remains untracked on disk (annotated with 2 done-at verdicts, never
   staged; its owner session should claim it).

## d) TOTALLY FUCKED UP (no excuses)

1. **Sequencing: M9 before M10.** I refreshed 15 line-number citations
   against a tree I was about to grow by ~40 lines in the same session —
   the vatid.go extraction now forced by the size gate re-drifts ~6 of them.
   The work is honest but half wasted; a citation refresh belongs AFTER the
   code moves it references, and I knew M10 was next in tier order.
2. **The size gate should have been predicted.** Turn 1's session hit the
   exact same gate (taxnumber.go was born from it). Adding ~40 lines to a
   325-line file was foreseeable; `wc -l` before writing new code into
   already-large files is now a personal rule.
3. **Three doc-comment indentation mistakes in one session** — the edit
   tool inserted tab-indented continuation lines into column-0 doc comments
   twice (RegisterRef, then both IsValid methods), plus a changelog bullet
   mislaid in the wrong version section during the v0.2.0 stamp. All caught
   by my own re-reads, but the same trap three times is a habit, not luck.
4. **A factually wrong test vector.** I asserted `IsValidFor("FR")` must
   reject an 11-digit number because "DE bounds must not leak" — wrong: the
   generic 5-16 rule legitimately accepts it. The test suite caught me
   (that is its job); I replaced it with a vector that actually demonstrates
   country dependence (9 digits: DE rejects, FR accepts).
5. **The drift gate's first draft had a false positive** — forbade
   `PARTIALLY_FUNCTIONAL` everywhere, tripping on the legend that DEFINES
   the vocabulary. Caught on first run of the clean tree; fixed to
   row-level patterns (a status row carries a feature cell before the
   marker). A gate that fails on its own repo's clean tree proves nothing.
6. **The seeded-drift revert left a stray blank line** in README (sed
   removed the text, not the added newline). Caught by `git diff --stat`,
   fixed; the demo's honesty depended on the revert being byte-clean.
7. **Invoices go.work**: my sweep created a workspace pinning invoices to
   local go-legalparty HEAD, then I deleted it via trash after use — correct
   ending, but I should have used a /tmp GOWORK file there from the start
   (as I did for cqrs-htmx) instead of creating then cleaning up a file in
   a foreign repo.

## e) WHAT WE SHOULD IMPROVE (process, going forward)

1. **`wc -l` before writing into large files** — the size gate is a
   first-class constraint; check it before, not after.
2. **Citation-touching work goes last in any code-touching session** —
   line-number references are derived artifacts; refresh them only after
   the tree settles (or anchor citations to symbols, not line numbers).
3. **Isolated GOWORK files everywhere** — never create a workspace file in
   a foreign repo, even temporarily, even gitignored.
4. **Doc-comment edits: paste the full comment block including column-0
   `//` prefixes** — never splice continuation lines; that is where all
   three indentation mistakes came from.
5. **Verify rule semantics before writing the assertion** — one minute
   reading `validateTaxNumberBody` would have saved the wrong FR vector.
6. **Run the repo's own gates on a new gate's first draft** — the
   seeded-drift demo did this well; do the same for absence-check patterns
   against legend/comment noise (grep before you forbid).

## f) WHAT TO DO NEXT (in order, resuming exactly here)

1. Extract `vatid.go` (VATID block: `identifiers.go` lines 206–365) +
   `vatid_test.go` (move `TestParseVATID`); verify: `go test -race`,
   size gate green on both files.
2. Re-point drift-gate anchors (4 VAT checks: validDEVATChecksum,
   want=8/10, NL shape message → `vatid.go`) and re-verify 26/26 green.
3. Second citation pass: update the ~6 moved refs (modularity 287,
   type-model 287/298/208/260-324, FEATURES 200) to `vatid.go` positions;
   AGENTS.md file list gains `vatid.go`.
4. Commit M10 (code + tests + docs + formatting wave) with a proper message.
5. Push master (3 queued commits + M10) — awaiting owner push authorization
   (§g1) — then verify CI green incl. the new canonical-facts job.
6. Re-check pkg.go.dev v0.2.0 render; annotate audit-all items 14/20
   (16-24 §f/20 and 18-44 §f/14) with the verified render verdict.
7. M11: read legalparty-model's leftover `identifiers_test.go`, rule
   aliases-vs-dead, act; run its aggregate suite pinned to v0.2.0.
8. M12: `BUILDFLOW_NO_RESULT_CACHE=1 buildflow --build-mode full`; triage
   findings; record baseline in AGENTS.md gotchas.
9. M13: try `BUILDFLOW_DPRINT_PLUGINS` local-wasm route; apply or document;
   kill the recurring preflight WARN.
10. M14: encode D3 policy (suppress gochecknoglobals with the "data tables"
    justification; keep wsl) via the buildflow-native config route.
11. M15: `sqlite3 ~/.local/state/buildflow/buildflow.db VACUUM` (1.07 GB);
    rebuild buildflow (`nix build . && nix run .#reinstall`); PATH check.
12. Final: TODO_LIST REGON-row removal + CHANGELOG purge entry; full
    battery (`go test -race`, vet, buildflow, docs-gates, drift gate);
    reconcile the 19-59 foreign file disposition with its owner session.
13. Post-v0.2.0 wave: M19 recurrence note (next Dependabot wave), monitor
    first CI runs of the new canonical-facts job.
14. Parked (demand-gated, untouched by design): M16 PL VATID checksum,
    M17 FR/IT/ES + TaxNumber rules, M18 registry IO client.

## g) QUESTIONS FOR THE OWNER (blocking decisions)

1. **Push authorization:** master is 3 commits ahead of origin (M5 contracts
   `c253fb0`, drift gate `c1d9381`, citation refresh `26ca5ed`) and M10 will
   add a fourth. The plan's release steps authorized push for v0.2.0 (done);
   these follow-ups were "commit" only in the fine tasks. Push now, or batch
   push after M10 lands?
2. **company-master-data `e61ec9f` is local-only** (REGON delegation, suite
   green). Push it, or leave it to that repo's own session/daemon rhythm?
3. **vatid.go extraction approved?** The size gate demands it (365/367 vs
   350). The alternative — raising the limit or nolint — contradicts the
   gate discipline that created taxnumber.go. Recommendation: extract, as
   scoped in §f1–f3.

---

**Verdict:** tier-1% and tier-4% of plan v2 are fully shipped and verified
(consumers green, v0.2.0 consumable, fleet checksum split brain closed);
tier-20% is 6/6 shipped. The rest-tier (M10–M15) is 0.8/6 — M10 is written,
green, and documented but gate-blocked on the extraction; M11–M15 are
untouched. The tree is honest about its own state: 3 unpushed commits,
18 modified files, and a gate demanding one more refactor before the
IsValid surface can land.
