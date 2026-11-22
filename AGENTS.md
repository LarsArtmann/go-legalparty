# AGENTS.md — go-legalparty

Pure Go value library: legal-party identifier types with real checksums
(NIP, Regon, VATID, LEI, TaxNumber, RegisterRef, CountryCode,
LegalFormClass) for the DE/PL/EU billing domain. Extracted from
`cqrs-htmx/legalparty-model` per ADR-0001 (history in CHANGELOG). Single
flat package `github.com/larsartmann/go-legalparty`, root-level files per
concept (`identifiers.go`, `taxnumber.go`, `nip.go`, `regon.go`,
`country.go`, `legalform.go`).

## Scope boundary (the most important rule)

This library owns **types, parsing, and checksums only** — no IO, no CQRS, no
persistence. The aggregate (decider/fold/events) stays in
`cqrs-htmx/legalparty-model`, which consumes this module. A GLEIF/registry
HTTP client does NOT belong here until a second consumer exists
(STACK.md revisit trigger; `gleif/` stays in legalparty-model). Adding a
dependency requires a stronger reason than "convenient" — the one-dependency
floor (go-error-family) is a deliberate design decision.

## Commands

No `flake.nix`, no Makefile, no justfile — plain Go module on purpose.

- CI (`.github/workflows/ci.yml`, every push and PR): race tests, vet,
  govulncheck, golangci-lint, a consumer-import simulation, dprint check,
  fleet workflow-policy vet. Local equivalents: the commands below.
- Tests: `go test ./...` (add `-race` per CONTRIBUTING.md)
- Fuzz: `fuzz_test.go` carries seed-corpus targets (NIP, Regon, VATID, LEI,
  TaxNumber); CI runs seeds only — local `-fuzz` runs are on-demand.
- Vet: `go vet ./...`
- Quality gate: `buildflow` — this repo is covered (the `.gitignore` block is
  buildflow-managed) but has no `.buildflow.yml`, so defaults apply. Use
  `buildflow --fix` for repairs, `buildflow -s <tool> --format finding` to
  drill in. Do not run `golangci-lint --fix` by hand (bypasses verify).
- Formatting: dprint owns markdown/json/yaml (`dprint.json`); Go formatting
  goes through buildflow.

Exit-code nuance: a buildflow run can be green per-step and still exit 1 —
the findings gate fails when error-severity findings remain. Read the
summary's findings section, not just the exit code.

## Tool gotchas

- `.golangci.yml` is buildflow-generated (`-s golangci-lint-auto-configure
  --fix` expands it). Don't hand-curate the linter list; it differs from
  sibling repos by design (per-project tailoring).
- buildflow's result cache can go stale after NEW files appear — a step
  that reports "0 files scanned" right after you add files is lying.
  Re-run with `BUILDFLOW_NO_RESULT_CACHE=1`.
- `dprint.json` fetches its four plugins over the network on first use —
  a cold run can hang waiting on that fetch. Options: point the config at
  local `.wasm` paths, delete it and let buildflow regenerate, or accept
  the fetch (cached in `~/.cache/dprint`).
- `docs-gates run .` (the doc battery: freshness, annotate, routing,
  cadence) also runs inside buildflow; it FATALs if TODO_LIST.md is missing.
- go-structure-linter suggests "use internal/ for non-public code" —
  deliberate non-fix: a pure value library is all-public by design.

## Error contract (public API, treat as frozen)

Every rejection is `errorfamily.NewRejection` with a stable machine code
`legalparty.<subject>.<cause>` (e.g. `legalparty.nip.checksum_failed`,
`legalparty.lei.invalid`). Dashboards key on these codes; consumers switch on
them. Changing a code is a breaking release. Messages are human-facing prose
and carry the offending input. Tests pin exact codes via the `errorCode`
helper — a new rejection path needs a `wantCode` vector.

## Validation semantics that look wrong but are correct

These encode regulation edges; do not "fix" them without a regulatory source:

- **NIP**: weights 6,5,7,2,3,4,5,6,7 mod 11; **remainder 10 is invalid**
  (checksum space excludes it — the mod-11 result must equal the check digit
  as-is). Hand-rolled copies that map remainder 10 to check digit 0 accept
  invalid NIPs — that bug class is why this library exists.
- **Regon**: weights 8,9,2,3,4,5,6,7 mod 11; **remainder 10 encodes check
  digit 0** — the opposite edge from NIP. The 14-digit local-unit form
  checks only the 9-digit parent prefix; the documented 13-weight serial
  check is deliberately NOT enforced (KRS-published branch REGONs falsify
  it — real master data fails that scheme).
- **DE VAT**: ISO 7064 Mod 11,10 over all 9 digits, **seed 5**, valid iff
  final check == 1. Faithful port of python-stdnum `mod_11_10`; verified
  vectors live in the `validDEVATChecksum` doc comment (BMF example
  DE136695976 passes, DE136695978 fails).
- **LEI**: ISO 17442 — map A=10..Z=35, whole string mod 97 must equal 1.
- **CountryCode**: officially-assigned ISO 3166-1 alpha-2 **only**. XK, EU,
  ZZ are deliberately absent from `assignedCountryCodes`; a billing
  jurisdiction must resolve to an assigned country. The table is
  hand-maintained — additions need an ISO source, not a vibe.
- **RegisterRef.Registry** is opaque by design: the domain never inspects its
  semantics (jurisdiction codes are owned by whoever resolves the reference).

## Type separation doctrine

`VATID`, `TaxNumber`, and `NIP` are deliberately distinct string types:

- A VAT ID starts with a country prefix; a national tax number never does.
  `ParseTaxNumber` rejects VAT-shaped input (`legalparty.tax_number.vat_shape`)
  instead of silently accepting it.
- Unprefixed domestic forms (bare NIP) must be normalized to prefixed form by
  the caller — `ParseVATID` will not guess a country.
- Normalization convention (fleet-wide): uppercase, trim, strip separators
  (space, `-`, `.`, `/`) via `normalizeTaxIdentifier`. Stored forms are
  canonical: `VATID` keeps its prefix, `NIP` and `Regon` store bare digits
  (PL prefix accepted on input, stripped on store).

## Consumers (blast radius of semantic changes)

1. `cqrs-htmx/legalparty-model` — the CQRS aggregate layer over these types
   (`aliases.go` bridges; its duplicate value sources are gone).
2. `invoices` — delegates all checksum validation to this library via
   `internal/models/tax_id_validation.go`.
3. `company-master-data` — direct dependency; keeps a local REGON checksum
   (the one identifier family not yet covered here — see TODO_LIST).

Loosening a validator here can silently un-fix a bug in a consumer; tighten
only with a regulatory reference and update the three consumers' tests.

## Documentation model

`README` sells, `STACK.md` records dependency decisions, `CHANGELOG.md` is
append-only history, `TODO_LIST.md` holds open bounded work only (done rows
are REMOVED and recorded in CHANGELOG), `FEATURES.md` is the honest status
inventory, `ROADMAP.md` holds demand-gated ideas, `docs/DOMAIN_LANGUAGE.md`
is the glossary (this library is vocabulary — its types ARE the fleet's
canonical identifier terms). Verify doc claims against code before trusting
them; fix drift on sight.

Snapshot convention: `docs/status/` holds point-in-time session reports —
immutable once written; superseding work gets a NEW report, and open items
in old reports are resolved by inline annotation (done-at verdicts), never
by rewriting history. The review-artifact dirs (`docs/architecture-understanding/`,
`docs/reviews/`, `docs/modularization/`) are the same class of artifact:
append-only, indexed by `docs/architecture-understanding/README.md`, and
tracked by docs-gates for annotation (never freshness-regressed by new
commits).

## Testing pattern

Table-driven, `t.Parallel()` on every test and subtest. `wantCode == ""`
means the input must be accepted AND normalize to `want` exactly; otherwise
the input must reject with exactly that machine code. Vectors are pinned
regulatory examples — the Polish MoF NIP (5260000152), the BMF DE VAT
example, live GLEIF records carried as constants — so tests need no network.
Add vectors for any new jurisdiction rule.

## Repo conventions

- Go 1.27; modern idioms are expected (`for i := range 9` int-range loops).
- Branch mainline is `master` (git-town config, GitHub API connector); PRs
  gate on CI — cut the PR branch BEFORE committing anything meant for the
  PR (a PR merged from an older branch tip silently strands later commits;
  that is how the five-skill review artifacts missed PR #1 and needed a
  cherry-pick).
- Dependabot: weekly gomod + github-actions updates, minor+patch grouped.
- Releases: CHANGELOG entry + annotated tag.
- `reports/` is a gitignored local artifact dir (coverage output).
