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

**Last updated:** 2026-11-23 · list drained to zero — every row shipped:
REGON adoption (company-master-data delegates to `ParseRegon`), invoices
`// indirect` fix, canonical-facts drift gate + CI leg, first full-mode
buildflow battery, dprint plugin decision (accept cached remote fetch),
legalparty-model alias-wiring pin, report annotations, `docs/status/`
convention line, census routing to ROADMAP. Open long-term ideas live in
`ROADMAP.md`.

## High Impact

_(none — the table is empty by convention: shipped rows are removed, not
ticked)_

## Medium Impact

_(none)_

## Low Impact

_(none — see the one host-hygiene row below; the work-item tables are empty
by convention: shipped rows are removed, not ticked)_

### Host hygiene (not repo work — carried here so it has a home)

| Task                                                                                                   | Status    | Impact | Effort | Evidence                                                                                                                           |
| ------------------------------------------------------------------------------------------------------ | --------- | ------ | ------ | ---------------------------------------------------------------------------------------------------------------------------------- |
| Rebuild the stale buildflow binary once BuildFlow's tree is clean (a concurrent session is editing it) | 🔴 `TODO` | Low    | M      | Doctor: binary `769b3a4` vs HEAD `38d6d91ce`; `nix build .` fails on the `-dirty` go-modules hash while their session is in flight |
