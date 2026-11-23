# Status Report — Pareto v2 Closure, v0.3.0 Release, Master Healed

**Written:** 2026-11-23 02:57 CET · session agent: glm-5.3, Crush
**Scope:** the closure leg of the pareto-execution plan v2 — healing red
master, verifying the concurrent session's shipped work, closing
M11–M15, and cutting v0.3.0. This report covers THIS session's run only;
the companion reports `2026-11-23_01-12` and `2026-11-23_02-22` cover the
concurrent legs.

**Verified here (fresh commands, this session):** `go test -race -count=1
./...` ok · `go vet ./...` clean · `golangci-lint run` 0 issues ·
`bash scripts/check-canonical-facts.sh .` 26/26 · `dprint check` exit 0 ·
`docs-gates run .` 5/5 PASS · `BUILDFLOW_NO_RESULT_CACHE=1 buildflow
--build-mode full` **EXIT=0** · master CI read back green on `a57b54f`,
`7c3f552`, `dc4e640`, `21e1160`, `fc52cb2`, `2058260` (run `45095778152`),
final tip `df5df27` (run `45096232056`) — the one red (`e2c2359`, run
`45090971301`) was the daemon's unformatted sweep, superseded green within
5 minutes · isolated-worktree dprint+canon verification of `1b46ae3`
before anyone pushed it · scratch module `go get go-legalparty@v0.3.0` +
smoke: all five IsValid surfaces agree; proxy version list shows v0.3.0 ·
`gh release list`: all four releases live · buildflow DB `wal_checkpoint
(TRUNCATE)` → WAL 0 bytes; binary `769b3a4` → `f129f43` via `nix build . &&
nix run .#reinstall`, doctor green · legalparty-model `GOWORK=off go test
-race ./...` 3/3 green against v0.2.0 · company-master-data delegation
grepped: `ValidREGON` → `legalparty.ParseRegon`, no local checksum symbols
· push state: `git rev-list origin/master..master --count` = 0 on all four
repos at close.

**NOT verified here:** the invoices and company-master-data full suites
(trusted from their sessions' reports after verifying push state and the
delegation grep — not re-run) · the `e2c2359` red run's `--log-failed`
(cause predicted from the daemon pattern; the 02-22 report's a5 confirmed
it) · pkg.go.dev's v0.3.0 page render (crawl lag; license suppression
already accepted) · godoc content anywhere (license-blocked).

## a) FULLY DONE

| #   | What                                                                                                                                                                                                                                                                         | Evidence                                                                                                     |
| --- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------ |
| a1  | **Master healed and kept green** — formatted the daemon-swept annotation tables (`a57b54f`), every subsequent push gated green through final tip `df5df27`                                                                                                                   | runs 45091521690 / 45092027201 / 45092999112 / 45093387807 / 45093446758 / 45095778152 / 45096232056 success |
| a2  | **v0.3.0 released end-to-end** — CHANGELOG stamp, SECURITY.md version row (was stale at v0.1.1 through two releases), CI-verified-green BEFORE tagging (the d4 lesson, applied), annotated tag, proxy + scratch-module smoke, GitHub Release `--prerelease`                  | `2058260` + tag `v0.3.0`; `go list -m -versions` shows it; smoke main.go all-green                           |
| a3  | **All four GitHub Releases backfilled** — none existed at session start; v0.1.0/v0.1.1/v0.2.0 notes cut from their CHANGELOG sections                                                                                                                                        | `gh release list`: 4 pre-releases                                                                            |
| a4  | **M12 full-mode battery, verified by running it** — the 7c3f552 backlog-drain CLAIMED it but the buildflow DB had no go-legalparty run row; direct execution: 53 steps, 0 failed, 29 warning-severity findings triaged non-error                                             | `EXIT=0` output captured; DB `resume_runs` queried (no row)                                                  |
| a5  | **M15 host hygiene** — DB WAL checkpointed + VACUUMed (1.1 GB is live data, not bloat), buildflow rebuilt from BuildFlow HEAD `f129f43` and reinstalled once that tree settled; doctor green; TODO row dropped                                                               | `buildflow --version` before/after; `df5df27`                                                                |
| a6  | **M11 closed** — verdict: legalparty-model's `identifiers_test.go` is ALIVE (it pins the aliases bridge against stub wiring and hosts the GLEIF fixtures), suite race-green on v0.2.0; go.mod already required v0.2.0 (verified, not assumed)                                | 3/3 packages ok; aliases.go + test read in full                                                              |
| a7  | **M13/M14 verified, not just claimed** — local dprint-plugin route is an empty stub (find: no .wasm), 198 MB plugin cache exists, CI portability argument holds; `.golangci.yml` encodes the policy (test-file exclusions at :249, justified nolints on the two spec tables) | grep/read of the actual config and dirs                                                                      |
| a8  | **Plan-v2 errata closed** — five rows appended: F02 shipped-as-fixtures (not N/A), F44/45 resolved-keep, medium-task closure, both master docs-format incidents with fixes                                                                                                   | `dc4e640`                                                                                                    |
| a9  | **AGENTS gotcha: the daemon bypasses dprint** — two same-day reds from auto-synced unformatted tables; guard is `dprint check` after every daemon sweep; verified complementary (not duplicative) vs the M13 entry                                                           | `21e1160`; AGENTS.md:81-84                                                                                   |
| a10 | **02-22 foreign report hand-committed with an honest message** instead of leaving it for the daemon sweep (their own d2 lesson, applied by me)                                                                                                                               | `c0ddf8f`                                                                                                    |

## b) PARTIALLY DONE

| #  | Item                     | Works now                                                                 | Open remainder                                                                                                       | Blocker             | Effort |
| -- | ------------------------ | ------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | ------------------- | ------ |
| b1 | Consumer v0.3.0 adoption | Released, proxy-served, additive API (IsValid) — nothing breaks on v0.2.0 | All three consumers still pin v0.2.0; no one consumes IsValid yet                                                    | None — cadence call | S each |
| b2 | Errata accuracy          | The verdict rows are right                                                | `dc4e640`'s closure row says the rebuild was "DNS-blocked" — wrong cause (see d1); needs one appended clarifying row | None                | S      |
| b3 | 02-22 §f harvest         | `docs-gates` green today (freshness window not yet triggered)             | The 50 next-things live ONLY in the report — TODO_LIST is at zero rows; entombment risk the skill warns about        | None — HARVEST pass | M      |
| b4 | pkg.go.dev v0.3.0        | Proxy + scratch verified; v0.1.1/v0.2.0 pages render per the 02-22 report | v0.3.0 page crawl not re-checked (accepted lag, license suppression unchanged)                                       | Crawler             | S      |

## c) NOT STARTED

| #  | Item                                                                                                                           | Why not started                                            | Priority      |
| -- | ------------------------------------------------------------------------------------------------------------------------------ | ---------------------------------------------------------- | ------------- |
| c1 | pkg.go.dev godoc decision (license redistributability vs docs mirror vs accept)                                                | Owner decision — g1, inherited from 02-22                  | Owner gate    |
| c2 | RegisterRef field branding (`RegistryCode`/`RegisterNumber`) vs opaque freeze                                                  | Owner decision — g3; public-API change, no consumer demand | Owner gate    |
| c3 | ROADMAP demand-gated jurisdiction work (PL/NL/FR real VAT checksums, AT TaxNumber, rule seam)                                  | Gated by design; no consumer pull                          | Medium, gated |
| c4 | Foreign-repo red gates (cqrs-htmx: docs-gates audit-all/cadence + 97 file-size errors; invoices: file-size/docs-gates/cue-vet) | Their backlogs — this session's scope was go-legalparty    | High THERE    |
| c5 | Rejection-code freeze test (enumerate every `legalparty.*` code, pin each to a vector)                                         | Strong idea surfaced in 02-22 §f7; unrowed, unplanned      | High          |
| c6 | example_test.go / coverage floor / table-type fuzz targets                                                                     | Nice-to-haves from 02-22 §f19/20/21; unrowed               | Low           |

## d) TOTALLY FUCKED UP

| #  | What broke / what I got wrong                                                                                                                                                                                                                                                                                                                                                                                                             | Severity | Root cause                                                               | Mitigation                                                                   |
| -- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- | ------------------------------------------------------------------------ | ---------------------------------------------------------------------------- |
| d1 | **Committed a wrong CAUSE into the plan errata.** `dc4e640` says the buildflow rebuild was "DNS-blocked" — I diagnosed from the tail of ONE failed `nix build` (a DNS timeout line was visible) and never read the full log. The real blocker was the `-dirty` go-modules hash from a concurrent session editing BuildFlow (02-22 report d7). The verdict row itself was right; the cause is misinformation in a semi-immutable artifact. | Medium   | Tail-of-log pattern-matching; wrote the cause before the retry proved it | Append one clarifying errata row (f1); rule: full log before permanent prose |
| d2 | **VACUUM hygiene theater.** `du` showed 2.2 GB post-VACUUM and I briefly misread it; after checkpoint the WAL went to 0 and the main DB stayed 1.1 GB — freelist was ~0, the DB was never bloated. The 01-12 report never claimed bloat; I ran the expensive operation without measuring first.                                                                                                                                           | Low      | Ran prescribed hygiene without `PRAGMA freelist_count` pre-check         | Measure bloat before vacuuming; the wasted work was ~2 min                   |
| d3 | **Two stale-mtime edit rejections in a row** (AGENTS.md, CHANGELOG.md) — I edited off `sed`/`grep` output instead of a registered `view`. Both recovered, but this is the recurring trap.                                                                                                                                                                                                                                                 | Low      | Grep-read feels like a read; the edit tool disagrees — rightly           | Registered `view` immediately before every `edit`                            |
| d4 | **Arbitrary-sleep CI polling.** `sleep 30/40` then one-shot status checks instead of `gh run watch` from push time; one background watch (job 03F) died to a connection reset and sat unnoticed for minutes.                                                                                                                                                                                                                              | Low      | Improvised polling instead of the purpose-built blocking command         | `gh run watch <id> --exit-status` in background at push time                 |
| d5 | **Todo list went stale mid-flight** — updated at start and end; in between it claimed "verify master CI" in_progress while four other workstreams ran.                                                                                                                                                                                                                                                                                    | Low      | Treated the tracker as bookkeeping, not steering                         | Update on every phase transition                                             |
| d6 | **Inverted verify-trust posture on M12.** I almost accepted 7c3f552's "first full-mode battery" claim at face value; only the missing DB row triggered the direct run. With three concurrent writers, report claims are hypotheses until re-run.                                                                                                                                                                                          | Low      | Concurrent-session reports read as authority                             | Fresh command beats report citation, every time                              |

## e) WHAT WE SHOULD IMPROVE

| #  | Improvement                                                                                                                                             | Impact |
| -- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------ |
| e1 | **Full error log before permanent prose** — one `nix log <drv>` would have prevented d1; causes written into append-only artifacts must be retry-proven | High   |
| e2 | **Registered view before every edit** — `grep`/`sed` reads do not satisfy the edit tool and cost a failed round-trip each time (d3)                     | Medium |
| e3 | **`gh run watch --exit-status` at push time, in background** — replaces sleep-polling and dead-watch blind spots (d4)                                   | Medium |
| e4 | **Re-run concurrent sessions' claims** — with a daemon and sibling agents writing to the same repos, a report line is a hypothesis (d6)                 | High   |
| e5 | **Batch independent doc commits** — four sequential docs pushes each paid a (cached, ~50 s) CI run; dependency-free doc edits can share a commit        | Low    |
| e6 | **Measure before hygiene** — `PRAGMA freelist_count` before VACUUM, `wc -l` before size-gate writes, `du -sh` before/after (d2)                         | Medium |
| e7 | **Fold the SECURITY.md version row into the release checklist** — I caught it ad hoc at v0.3.0; v0.2.0 shipped with a two-versions-stale row            | Medium |
| e8 | **Update the todo tracker at every phase transition** — a stale in_progress row is a lie with extra steps (d5)                                          | Low    |

## f) Next things (25 specific — the extended backlog is 02-22 §f's 50 rows, routed by reference)

**Harvest note:** rows 1-3 are THIS session's debts; 4+ are the actionable
subset of 02-22 §f that still has no TODO_LIST/ROADMAP row. High-impact
rows belong in TODO_LIST, idea-class in ROADMAP, foreign-repo rows in
those repos' backlogs.

| #  | Task                                                                                                                            | Impact | Effort | Category      | Lands in                  |
| -- | ------------------------------------------------------------------------------------------------------------------------------- | ------ | ------ | ------------- | ------------------------- |
| 1  | Append the clarifying errata row: rebuild blocker was the dirty BuildFlow tree, not DNS (d1)                                    | Medium | S      | Documentation | this repo                 |
| 2  | HARVEST pass: row 02-22 §f's unrowed high-impact items into TODO_LIST/ROADMAP before docs-gates freshness turns                 | High   | M      | Documentation | this repo                 |
| 3  | Row g1/g3 owner questions durably (ROADMAP decisions section) so they outlive the reports                                       | Medium | S      | Documentation | this repo                 |
| 4  | Add "CI green on tag commit" + "SECURITY.md version row" to AGENTS Releases conventions (e7; 02-22 e4)                          | High   | S      | Process       | this repo                 |
| 5  | Fleet bump to v0.3.0 (cqrs-htmx/legalparty-model, invoices, company-master-data) — or decide Dependabot-lazy                    | Medium | S×3    | Maintenance   | consumers                 |
| 6  | Rejection-code freeze test: enumerate every `legalparty.*` code in source, pin each to a vector (02-22 f7)                      | High   | M      | Quality       | this repo                 |
| 7  | Prove branching-flow's nip/rego index warnings structurally impossible (fixed-length loops) and document, or add bounds asserts | Medium | S      | Quality       | this repo                 |
| 8  | Rename `wrap2` in legalparty-model's alias test to an honest name (02-22 e9)                                                    | Low    | S      | Cleanup       | repo: cqrs-htmx           |
| 9  | Check company-master-data's CHANGELOG carries the REGON-adoption entry post-daemon-sweeps (02-22 f45)                           | Low    | S      | Documentation | repo: company-master-data |
| 10 | Re-check pkg.go.dev v0.3.0 module page after crawl lag                                                                          | Low    | S      | Release       | this repo                 |
| 11 | Restart LSP; confirm the stale gofumpt/gocognit warnings clear (02-22 e7)                                                       | Low    | S      | Cleanup       | host                      |
| 12 | Widen the AGENTS cache gotcha: also bites single-step drill-downs (02-22 e6/f8)                                                 | Low    | S      | Documentation | this repo                 |
| 13 | `example_test.go` with runnable godoc examples (NIP/VATID/LEI) ready for when license allows rendering (02-22 f19)              | Low    | S      | Documentation | this repo                 |
| 14 | Coverage floor decision + CI surfacing (full mode already computes it; 02-22 f21)                                               | Medium | S      | Quality       | this repo                 |
| 15 | Fuzz CountryCode/LegalFormClass/normalizeTaxIdentifier (idempotence, round-trip; 02-22 f20/f37)                                 | Low    | S      | Quality       | this repo                 |
| 16 | Fleet sweep: any consumer still below v0.2.0? (`go list -m all` across repos; 02-22 f36)                                        | Medium | S      | Maintenance   | fleet                     |
| 17 | Re-verify SECURITY.md's govulncheck claim matches ci.yml after the canonical-facts job landed (02-22 f27)                       | Low    | S      | Documentation | this repo                 |
| 18 | Provenance note in architecture-understanding README: d2/svg pairs pre-date the vatid.go extraction (02-22 f30)                 | Low    | S      | Documentation | this repo                 |
| 19 | cqrs-htmx red gates: docs-gates audit-all/cadence + 97 file-size errors (02-24/25) — foreign, high for that repo                | High   | M/L    | Quality       | repo: cqrs-htmx           |
| 20 | invoices red gates: file-size, docs-gates, cue-vet permissions (02-22 c4)                                                       | Medium | L      | Quality       | repo: invoices            |
| 21 | Decide `coverage.out` in legalparty-model (committed build artifact; 02-22 f32)                                                 | Low    | S      | Cleanup       | repo: cqrs-htmx           |
| 22 | GOWORK-isolation pattern → crush-config lessons.md (it prevented three consumer-repo writes; 02-22 f46)                         | Low    | S      | Process       | crush-config              |
| 23 | Version-marker on check-canonical-facts.sh before it gets copied fleet-wide (02-22 f43)                                         | Low    | S      | Maintenance   | this repo                 |
| 24 | Survey ParseLegalFormClass lowercase-normalization asymmetry vs CountryCode; deliberate verdict (02-22 f33)                     | Low    | S      | Feature       | this repo                 |
| 25 | Verify the five-skill review's open items all landed post-drain (audit-all says yes; spot-check one; 02-22 f48)                 | Low    | S      | Documentation | this repo                 |

## g) Questions I cannot answer myself

1. **pkg.go.dev godoc (inherited g1):** keep the LICENSE proprietary and
   accept that pkg.go.dev never renders the API docs, make it
   redistributable, or build a docs mirror? This is a legal/branding call
   only you can make, and it gates `example_test.go` value and any
   future public-facing docs.

2. **RegisterRef field branding (inherited g3):** branching-flow flags the
   plain-string `(registry, number)` pair as swap-prone. Branding the
   fields (`RegistryCode`, `RegisterNumber`) makes the swap
   unrepresentable but is a public-API change against a documented
   opaque-by-design contract. Type safety or frozen opacity?

3. **v0.3.0 fleet cadence:** bump all three consumers to v0.3.0 now
   (the IsValid surface is consumer-visible), or let the weekly Dependabot
   wave drive adoption lazily? Both are defensible; it is your cadence
   preference, not a technical fact.
