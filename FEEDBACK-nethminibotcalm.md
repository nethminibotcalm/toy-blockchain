# Feedback on Your Toy Blockchain Submission — nethminibotcalm

**Reviewer:** Dasun · **Date:** 2026-07-10

Your update is a real step forward: the CLI, JSON persistence, README, research report, and the mining-difficulty test are all new and all count. However, the core correctness issues flagged in the first review are still in the code, and the new persistence work introduced a serious new one. Please focus your next revision on the weaknesses below.

## Weaknesses

**Blockchain core (highest weight in the rubric)**

- The genesis block still uses the current time, so a fresh clone produces a different chain on every run. Determinism is a MUST requirement.
- The block hash is still built by concatenating fields with no separators, so different blocks can produce identical hash input.
- Money is still `float64`. Amounts must be integers.
- Chain validation still only checks hashes and links. It does not check proof-of-work, block order/timestamps, or replay balances — so a tampered-and-rehashed block, or a chain with negative balances, passes as valid. The genesis block is never validated, and validation doesn't report which block failed.
- Mining difficulty is still hardcoded at the call site, and mining reports no attempt count or elapsed time.
- Double-spending from the pending pool is still possible: two transactions that together exceed the sender's balance are both accepted and mined.

**New issue introduced with persistence**

- Every time the program restarts, account balances reset to the hardcoded 100 each, regardless of what the loaded chain says. The saved chain and the ledger disagree, so someone can spend their full balance again in every new session. Balances need to come from the chain itself.
- The chain loaded from `chain.json` is never validated on load, so a hand-edited file is accepted silently.

**Documentation accuracy**

- The README and report both claim that validation checks proof-of-work difficulty — the code does not do this. Never document behaviour that isn't implemented; it reads worse than an honest omission.
- The README's project structure lists `cli/` and `storage/` folders that don't exist.
- The report's difficulty-vs-effort table has no measured numbers (attempts, time) — the spec asks for real data.
- `go.mod` declares `go 1.26.4` while the README says Go 1.22+.

**Code quality**

- Every `.go` file fails `gofmt` — inconsistent indentation and spacing throughout. This is a one-command fix and it is expected to be clean.
- Two of your tests can never fail (hashing the same struct twice, validating a chain with a validator that checks almost nothing). There are no tests for double-spending, persistence, or tampering with a re-computed hash.

## What to improve first

1. Make balances derive from the chain, and fix the double-spend and restart-reset problems together.
2. Complete chain validation (all five invariants from the spec) and validate the chain when loading from file.
3. Fix determinism: constant genesis timestamp, separators in the hash input, integer amounts.
4. Make difficulty configurable, have mining report attempts/time, and put real measured numbers into the report's table.
5. Run `gofmt` on everything and correct the README so it only describes what the code actually does.

You are much closer than the first submission — but note that most of the items above were in the previous review verbatim. Fixing flagged issues before adding new features is what we're looking for. Be ready to walk through your validation and mining flow unaided at the review call.
