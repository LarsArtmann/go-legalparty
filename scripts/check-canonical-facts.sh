#!/usr/bin/env bash
# check-canonical-facts.sh — the canonical-facts drift gate.
#
# Pins the README/FEATURES validation-semantics rows to the code that owns
# them. The README table drifted for a full session once (FI VAT length,
# branch/trust classes, PL TaxNumber rule) while claiming "verified"; this
# gate makes that class of drift exit non-zero. Each fact is checked on
# BOTH sides: the doc must carry the claim AND the code must carry the
# anchor that makes it true — a doc edit alone, or a code edit alone,
# trips the gate.
#
# Usage: scripts/check-canonical-facts.sh [repo-root]
# Exit:  0 all canonical facts hold; 1 any drift found.

set -uo pipefail

ROOT="${1:-.}"
cd "$ROOT" || exit 2

failures=0
total=0

ok()   { printf '[ok]   %s\n' "$1"; total=$((total + 1)); }
drift() { printf '[DRIFT] %s\n' "$1"; failures=$((failures + 1)); total=$((total + 1)); }

# check <label> <file> <grep-extended-pattern>
check() {
	local label=$1 file=$2 pattern=$3
	if grep -qE -- "$pattern" "$file" 2>/dev/null; then
		ok "$label"
	else
		drift "$label — pattern not found in $file: $pattern"
	fi
}

# absent <label> <file> <grep-extended-pattern>
absent() {
	local label=$1 file=$2 pattern=$3
	if grep -qE -- "$pattern" "$file" 2>/dev/null; then
		drift "$label — forbidden pattern present in $file: $pattern"
	else
		ok "$label"
	fi
}

# --- Presence: every documented rule has its code anchor -----------------

check "NIP weights row (README) and NIP weights (nip.go)" \
	README.md 'remainder 10 is invalid'
check "NIP weights constant" \
	nip.go '6, 5, 7, 2, 3, 4, 5, 6, 7'

check "REGON remainder-10-encodes-0 row (FEATURES)" \
	FEATURES.md 'Remainder-10-encodes-0'
check "REGON weights constant" \
	regon.go 'weights := \[8\]int\{8, 9, 2, 3, 4, 5, 6, 7\}'
check "REGON serial-unchecked rationale (FEATURES)" \
	FEATURES.md 'serial deliberately unchecked'

check "DE VAT ISO 7064 claim (README)" \
	README.md 'ISO 7064 Mod 11,10'
check "DE VAT checksum function" \
	identifiers.go 'func validDEVATChecksum'

check "AT/PL fixed-length claim (README)" \
	README.md 'AT/PL fixed lengths'
check "AT fixed length 8 (want = 8)" \
	identifiers.go 'want = 8'
check "PL fixed length 10" \
	identifiers.go 'want = 10'

check "NL BTW shape claim (README)" \
	README.md 'NL BTW shape'
check "NL strict-shape rejection message" \
	identifiers.go '9 digits, B, and 2 digits'

check "five legal-form classes row (README)" \
	README.md 'natural_person / legal_entity / partnership / public_body / non_profit'
check "legal-form class constant list" \
	legalform.go 'LegalFormNonProfit LegalFormClass = "non_profit"'

check "CountryCode assigned-only claim (README)" \
	README.md 'no XK/EU/ZZ'
check "assigned-country table" \
	country.go 'var assignedCountryCodes'

check "TaxNumber DE-only jurisdiction rule (README)" \
	README.md 'Steuernummer'
check "TaxNumber per-jurisdiction body validator" \
	taxnumber.go 'func validateTaxNumberBody'

# --- Absence: the three drifts that burned a session, and status honesty --

absent "no branch/trust class naming (README)" \
	README.md 'branch/trust'
absent "no FI fixed-length VAT row (README)" \
	README.md '`FI`.*fixed|fixed.*`FI`'
absent "no PL NIP-linked TaxNumber rule claim (README)" \
	README.md 'PL NIP-linked'
# Status honesty: the legend (lines 13-16) DEFINES the four status words;
# the checks below forbid them as STATUS ROWS (a row carries a feature-name
# cell before the marker cell, the legend has the marker first).
absent "FEATURES has zero PARTIALLY_FUNCTIONAL rows" \
	FEATURES.md '\|[^|]+\|\s*🟡'
absent "FEATURES has zero BROKEN rows" \
	FEATURES.md '\|[^|]+\|\s*🔴'
absent "FEATURES has zero DISABLED rows" \
	FEATURES.md '\|[^|]+\|\s*🔵'
absent "no MIT license claim (README)" \
	README.md '\bMIT\b'

# --- License truth --------------------------------------------------------

if [ -f LICENSE ]; then
	ok "LICENSE file present"
else
	drift "LICENSE file missing"
fi

# --- Verdict ---------------------------------------------------------------

printf '\n'
if [ "$failures" -eq 0 ]; then
	printf 'canonical-facts: CANONICAL (%d checks green)\n' "$total"
	exit 0
fi
printf 'canonical-facts: DRIFT (%d of %d checks red) — docs and code disagree; fix the docs or the code, not the gate\n' "$failures" "$total"
exit 1
