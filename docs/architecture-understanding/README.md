# Architecture-understanding series index

The four HTML reports and two diagram pairs produced by the 2026-11-22
five-skill review session (architecture-deepening, architecture-review,
architecture-visualization, data-model-review, go-modularize), with the one
question each answers and its standing verdict. Read the one whose question
you carry; each is self-contained.

| Artifact                                                                 | Question it answers                                                        | Standing verdict                                                                                                    |
| ------------------------------------------------------------------------ | -------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `2026-11-22_19-02-go-legalparty-structure.svg` (+ `.d2` source)          | What does the module look like right now?                                  | 8 root-level files, one flat package, one dependency                                                                |
| `2026-11-22_19-02-go-legalparty-structure-improved.svg` (+ `.d2` source) | What would a restructured version look like?                               | Differently-shaded same shape: no restructuring worth its cost                                                      |
| `2026-11-22_19-05_modularity.html`                                       | How modular is the codebase, scored?                                       | 4.3/5 across seven dimensions; maintain, do not restructure                                                         |
| `2026-11-22_19-10_value-library-deepening.html`                          | Where is the nearest architecture seam worth deepening?                    | Candidate 1 (IsValid surface) worth exploring; candidate 2 (jurisdiction rule seam) speculative — routed to ROADMAP |
| `../reviews/2026-11-22_go-legalparty-type-model.html`                    | Which invariants do the types enforce, where are the stringly-typed seams? | 8 problems, M1 = uniform IsValid surface (the deepening brief's pick)                                               |
| `../modularization/2026-11-22_SINGLE_MODULE_VERDICT.html`                | Should this split into multiple Go modules?                                | KEEP single module; every proposed seam fails gate 2 in both currencies                                             |

On-sight fixes that shipped during the same session are recorded in
`CHANGELOG.md` under v0.1.1. The executable follow-ups live in
`../planning/2026-11-22_19-38_pareto-execution-plan-v2-verified-backlog.md`
(the IsValid surface is task M10, gate D1).
