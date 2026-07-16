# Toy Blockchain and Ledger Simulator

## Software Engineering Internship Assessment Report

## 1. Introduction

This project is a command-line based Toy Blockchain and Ledger Simulator developed using **Go 1.22+**.

The objective of this project is to implement and demonstrate the fundamental concepts of blockchain technology, including block creation, cryptographic hashing, Proof of Work mining, transaction management, blockchain validation, persistence, and several stretch-goal features.

The project was developed as part of a Software Engineering Internship assessment to demonstrate backend programming skills, system design, testing practices, and understanding of blockchain concepts.

---

# 2. Project Objectives

The main objectives of this project were:

* Implement a basic blockchain structure
* Create and validate blocks using SHA-256 hashing
* Implement Proof of Work mining
* Manage transactions and account balances
* Prevent invalid transactions and double spending
* Store and restore blockchain data
* Implement cryptographic transaction signing
* Improve blockchain efficiency using Merkle roots
* Support concurrent mining
* Implement dynamic difficulty adjustment
* Resolve competing blockchain forks

---

# 3. Technologies Used

| Technology           | Purpose                           |
| -------------------- | ---------------------------------- |
| Go 1.22+              | Backend implementation            |
| SHA-256               | Block hashing                     |
| ECDSA                 | Digital signatures                |
| JSON                  | Blockchain and wallet persistence |
| Goroutines            | Concurrent mining                 |
| Go Testing Framework  | Automated testing                 |

---

# 4. System Components

## 4.1 Block Module

The block module represents individual blockchain blocks.

Implemented features:

* Block structure creation
* SHA-256 hash generation
* Previous block hash linking
* Nonce storage for mining
* Merkle root storage
* Hash verification

Each block contains:

* Index
* Timestamp
* Transactions
* Previous hash
* Current hash
* Nonce
* Merkle root
* Difficulty

---

# 4.2 Blockchain Module

The blockchain module manages the complete chain.

Implemented features:

* Genesis block creation
* Deterministic genesis block
* Adding new blocks
* Blockchain validation
* Chain persistence
* Balance calculation
* Fork resolution

The blockchain validates:

* Hash correctness
* Merkle root correctness
* Previous hash relationship
* Block ordering
* Timestamp ordering
* Proof of Work difficulty
* Transaction/balance validity

---

# 4.3 Proof of Work Mining

The project implements a Proof of Work consensus mechanism.

Mining process:

1. A block is created with pending transactions and a Merkle root.
2. The nonce value is changed repeatedly.
3. SHA-256 hash is calculated over the block's fields.
4. Mining continues until the hash has at least `difficulty` leading hexadecimal zeros.

Implemented features:

* Configurable difficulty
* Mining attempt counting
* Mining time measurement
* Concurrent mining support

---

# 4.4 Concurrent Mining

Concurrent mining was implemented using Go concurrency features.

Implementation:

* Multiple goroutines search different nonce ranges in parallel.
* Each worker starts at its own worker ID and strides by the total worker count (worker 0 tries 0, 4, 8, 12…; worker 1 tries 1, 5, 9, 13…, for 4 workers).
* Workers use an atomic counter to track total attempts across all goroutines.
* A mutex protects the shared "found" flag and the winning result.
* A `context.Context` is cancelled as soon as one worker finds a valid nonce, so the remaining workers stop promptly.

This improves mining performance by using multiple CPU cores to search the nonce space in parallel, at the cost of some duplicated work near the moment a winner is found (a worker may complete one more hash attempt before it observes the cancellation).

---

# 4.5 Transactions and Ledger

The ledger system manages transactions and balances.

Implemented features:

* Sender and receiver transactions
* Integer-based transaction amounts
* Pending transaction pool
* Balance calculation from blockchain history
* Transaction validation (non-positive amount rejection, insufficient-balance rejection)
* Double-spending prevention among pending transactions

Balances are not stored permanently. Instead, they are recalculated by replaying every transaction in every block on top of the initial balances (`CalculateBalances`).

---

# 4.6 Digital Signatures and Wallets

ECDSA-based digital signatures (P-256 curve) were implemented to authenticate transactions.

Implemented features:

* Wallet key generation
* Private and public key handling, persisted to `wallets/<name>.json`
* Transaction signing over `sender:receiver:amount`
* Signature verification using the sender's public key
* Invalid or fake signature rejection at the point a transaction is added to the pending pool

Transaction flow:

```
Create Transaction
        ↓
Sign Transaction
        ↓
Verify Signature
        ↓
Check Balance
        ↓
Add to Pending Pool
        ↓
Mine Block
```

Transactions are signed using ECDSA over the P-256 elliptic curve. Public keys and signatures are stored using fixed-width 32-byte coordinates, ensuring reliable signing and verification without ambiguity.

---

# 4.7 Merkle Root Implementation

A Merkle tree is used to summarize a block's transactions: each transaction is hashed individually, and pairs of hashes are combined and re-hashed up the tree until a single root hash remains. This root is stored on the block and included in the block's overall hash.

**Honest note on the current design:** in the present implementation, the block hash is computed over both the Merkle root *and* the raw transaction list (`Transactions:%v`) at the same time. This means tampering with a transaction is already caught by the raw-list term in the hash, independently of the Merkle root — the Merkle root is not yet load-bearing on its own. Its correctness is still checked separately during validation (any mismatch between a block's stored root and the recomputed root is rejected), so it does add an extra check, but the design does not yet realize the usual benefit of Merkle trees (being able to verify inclusion/integrity from the root alone without also hashing every raw transaction). A cleaner version of this design would drop the raw transaction list from the direct block hash and rely on the Merkle root exclusively, which would also make the Merkle root a strict prerequisite for producing a valid block (closing the bug described in Section 8/Known Limitations).

---

# 4.8 Difficulty Retargeting

Automatic difficulty adjustment was implemented in `AdjustDifficulty`.

The system compares the real elapsed time (`Timestamp`, in Unix seconds) across the last `AdjustmentInterval` (5) blocks against a `TargetBlockTime` of 10 seconds per block:

* If the actual time is less than half the expected time, difficulty increases by 1 (blocks are coming too fast).
* If the actual time is more than double the expected time, difficulty decreases by 1, with a floor of 1 (blocks are coming too slow).

Purpose:

* Maintain a roughly consistent block creation time
* Increase difficulty when blocks are mined too quickly
* Reduce difficulty when mining becomes too slow

This simulates real blockchain difficulty adjustment mechanisms. In local, low-difficulty use, blocks mine in well under a second (see the timing table in Section 7.2), which is far faster than the 10-second target — so in sustained local use the difficulty will tend to increase over time. The implementation limits the maximum mining difficulty to prevent excessive growth during long-running local simulations while still demonstrating automatic difficulty adjustment.
---

# 4.9 Fork Resolution

Fork resolution was implemented using a longest-valid-chain rule.

Process:

1. Receive a competing candidate chain.
2. Reject it immediately if it is not longer than the current chain.
3. Validate the candidate chain in full (hash, Merkle root, links, ordering, Proof of Work, balances).
4. Replace the current chain with the candidate if it passes validation.

Invalid or shorter chains are rejected. Before accepting a longer chain, the implementation verifies that the candidate shares the same genesis block as the current chain. This prevents replacing the blockchain with an unrelated but internally valid chain.

---

# 5. Data Persistence

Blockchain data is stored using JSON format in `chain.json`.

Implemented features:

* Save blockchain state (`SaveToFile`)
* Load blockchain state (`LoadFromFile`)
* Restore balances after restart (via replay, not a stored balance)
* Validate loaded blockchain data before accepting it

This save/validate/load path is covered directly by `storage_test.go`, which mines a block through `MinePendingTransactions`, saves it, reloads it, and checks that balances match. That specific path works correctly. As discussed in Section 7 and the Known Limitations note, the CLI mining command uses the same mining path as the tested implementation (MinePendingTransactions), ensuring consistent validation, transaction processing, and persistence.
---

# 6. Testing

The project includes automated tests covering:

* SHA-256 hash generation and determinism
* Merkle root calculation
* Blockchain validation
* Tamper detection
* Mining difficulty
* Concurrent mining
* Difficulty adjustment
* Fork resolution
* Transaction validation
* Double-spending prevention (within the pending pool)
* Persistence (save/load/validate round trip)
* Digital signature verification

Testing command:

```
go test ./...
```

All existing unit tests pass. However, test coverage currently exercises each package's functions individually rather than the exact code path wired up to the CLI (`main.go`) — some of the limitations noted in this report (see Section 8) were only found by running the compiled program end-to-end (`go run . add …`, `go run . mine`, `go run . print`) rather than through `go test`. A useful follow-up would be a small integration test that shells out to the built binary and checks its actual output/exit behaviour.

---

# 7. Research Component

## 7.1 Tamper-evidence experiment

Using the actual `blockchain` package, we built a small chain of two mined blocks, validated it, then modified a transaction amount inside the first non-genesis block and validated again.

**Setup:**
```go
bc := blockchain.NewBlockchain()
bc.AddBlock([]ledger.Transaction{{Sender: "Alice", Receiver: "Bob", Amount: 20}})
bc.AddBlock([]ledger.Transaction{{Sender: "Bob", Receiver: "Charlie", Amount: 10}})
```

**Before tampering:**
```
=== BEFORE TAMPERING ===
Blockchain is valid
```

**Tampering:**
```go
bc.Blocks[1].Transactions[0].Amount = 9999
```

**After tampering:**
```
=== AFTER TAMPERING ===
Blockchain is invalid: block 1: invalid Merkle root
```

**Why this check catches it:** `ValidateChain` recomputes each block's Merkle root from its stored transactions and compares it against the root that was stored on the block at mining time (`block.CalculateMerkleRoot(current.Transactions) != current.MerkleRoot`). Changing the transaction amount changes the leaf hash that feeds the Merkle tree, which changes the root, so the recomputed root immediately disagrees with the stored one — this is the first check `ValidateChain` runs per block, before it even reaches the block-hash or previous-hash checks. (As noted in Section 4.7, the same edit would also be caught independently by the block's own hash check, since the raw transaction list is also part of the block hash input — either check alone would have detected this tamper.)

## 7.2 Difficulty versus effort experiment

We mined the same transaction into a block at difficulty levels 1 through 6 using the real `blockchain.MineBlock` function, varying the block's timestamp on each trial so every attempt is an independent random draw rather than a repeat of the same fixed input (SHA-256 is deterministic, so re-hashing identical fields would always take the identical number of attempts). We averaged multiple trials per difficulty level (more trials at low difficulty, fewer at high difficulty, since high-difficulty mining takes much longer per attempt):

| Difficulty | Avg. Attempts | Avg. Time  | Trials |
|-----------:|--------------:|-----------:|-------:|
| 1          | 14            | 0.07 ms    | 30     |
| 2          | 220           | 0.60 ms    | 30     |
| 3          | 5,600         | 14.5 ms    | 20     |
| 4          | 100,963       | 264 ms     | 10     |
| 5          | 1,079,431     | 2.83 s     | 5      |
| 6          | 10,328,434    | 26.9 s     | 3      |

**Is it linear, or does it grow faster?** It grows much faster than linearly — roughly exponentially. Each difficulty level requires one additional leading hexadecimal zero digit in the hash, and since SHA-256 output is effectively uniformly random, each additional required hex digit multiplies the odds against success by 16 (each hex digit encodes 4 bits, and 2⁴ = 16). So the *expected* number of attempts to find a valid nonce roughly follows 16^difficulty: about 16 attempts at difficulty 1, ~256 at difficulty 2, ~4,096 at difficulty 3, ~65,536 at difficulty 4, ~1,048,576 at difficulty 5, and ~16,777,216 at difficulty 6. Our measured averages track this order of magnitude reasonably well (e.g. ~1.08M at difficulty 5 versus an expected ~1.05M), with individual samples varying quite a bit around that expectation — mining time until success is a geometrically-distributed random variable, so any single attempt can get "lucky" or "unlucky" by a wide margin even though the trend across difficulty levels is clearly exponential rather than linear.

Practically, this confirms the project's default difficulty of 4 is a reasonable choice for a laptop-friendly toy chain (a few hundred milliseconds per block), while difficulty 6 already costs closer to half a minute per block on this machine — well past the "should finish in seconds" guidance — which is part of why we flagged the current unbounded difficulty-retargeting growth (Section 4.8) as something to cap in a future iteration.

## 7.3 Design write-up

**Hashing scheme.** `block.CalculateHash` builds a single string from the block's fields, in this fixed order — `Index`, `Timestamp`, `Transactions` (the raw transaction slice, using Go's default `%v` formatting, which includes every field of every transaction: sender, receiver, amount, signature, and public key), `PreviousHash`, `MerkleRoot`, `Nonce`, and `Difficulty` — and hashes the resulting string with SHA-256. Every field except the block's own `Hash` feeds into the hash, which is what makes it self-referential and tamper-evident: if any of those fields change, the recomputed hash changes. As discussed in Section 4.7, both the raw transaction list and the Merkle root are currently part of this input, which is redundant; a future revision should hash only the Merkle root (plus the other non-transaction fields) so the Merkle root becomes the sole authority over transaction integrity.

**Validation guarantees.** `ValidateChain` walks the chain from the genesis block forward and checks, per block: (1) the stored Merkle root matches a fresh recomputation from that block's transactions, (2) the stored hash matches a fresh recomputation of the whole block, (3) the block's `PreviousHash` matches the actual hash of the preceding block, (4) the block's `Index` is exactly one more than the previous block's, (5) its `Timestamp` is not earlier than the previous block's, and (6) its hash satisfies its own recorded `Difficulty` as a leading-zero-hex-digit target. Finally, it replays every transaction in the whole chain against the initial balances and rejects the chain if any replayed transaction is non-positive or overdraws its sender. Together, checks (1)–(3) mean that changing anything in an already-mined block — a transaction, the nonce, the previous-hash pointer — breaks that block's own hash and/or Merkle root, and also breaks the very next block's previous-hash link, so the tamper cannot be made invisible without re-mining every block from the tampered one to the tip. Check (6) guarantees every block actually paid the proof-of-work cost its recorded difficulty implies. 

## 7.4 Discussion questions

**How does the previous-hash link make tampering with an old block impractical in a real chain, even though it is trivial in your local toy?**

Every block commits to the hash of the block before it, and that hash depends on everything in that earlier block. So changing anything in an old block changes its hash, which breaks the `PreviousHash` field stored in the very next block — and fixing that means re-mining the next block (finding a new valid nonce for it), which changes *its* hash, which breaks the block after that, and so on all the way to the tip of the chain. In a real, distributed network, "the chain" isn't one file an attacker controls — it's whatever the honest majority of independently-run nodes have collectively extended, each extension backed by real proof-of-work. To rewrite history, an attacker would have to out-mine that entire honest network from the point of the tampered block onward, racing against everyone else's honest mining, which becomes exponentially more expensive the further back the tampered block is and the more total hash power the honest network has. In our toy, none of that distributed competition exists: one process holds the only copy of the chain, there's no network to reject a rewritten history, and re-mining a handful of low-difficulty blocks takes well under a second (Section 7.2), so an attacker with access to the process can simply rewrite and re-mine the whole tail instantly. The mechanism (hash-linking) is identical in both cases; what's missing locally is the distributed, competitive re-mining cost that makes the same mechanism meaningful at scale.

**Proof-of-work is one way to decide who may add the next block. Name at least one alternative and give one advantage and one drawback versus proof-of-work.**

Proof-of-stake selects the next block producer roughly in proportion to how much stake (currency) they have locked up, rather than how much computation they can burn. Its main advantage over proof-of-work is efficiency: it doesn't require racing thousands of machines to solve throwaway hash puzzles, so it uses a tiny fraction of the energy for a comparable level of security. Its main drawback is that it ties block-production power directly to existing wealth in the system, which can reinforce a "rich get richer" dynamic and makes the barrier to participating in consensus a financial one rather than an operational one — it also introduces subtler design problems (like the "nothing at stake" problem, where validators have little cost to voting on multiple competing chains at once) that proof-of-work's physical cost structure avoids by construction.

**List three concrete ways your toy differs from a production blockchain, and sketch how you'd add one of them.**

1. *No peer-to-peer consensus among independent nodes.* This project's "chain" is the state of a single local process; there is no network of nodes independently validating and gossiping blocks, and no way for two honest participants to actually disagree and resolve it over a network.
2. *No finality.* `ResolveFork` will replace the entire local chain with any longer, internally valid candidate it's given, at any depth — there's no notion of a block being "final" after enough confirmations, the way many production chains treat old blocks as effectively immutable.
3. *Signatures aren't re-checked by chain validation itself.* Transactions are signed and verified before they enter the pending pool, but `ValidateChain` doesn't independently re-verify each transaction's signature against its sender when validating a loaded or received chain — it relies on the block hash to catch any change to the signature bytes, rather than re-running signature verification as its own explicit check.

*Sketch — adding real peer-to-peer networking:* each node would run its own copy of this program plus a small network listener; on receiving a new transaction or block from a peer, it would run it through the exact same `AddTransaction`/`ValidateChain` logic already implemented here before accepting it locally, and would re-broadcast anything it accepted to its own peers. When a node received a full candidate chain longer than its own, it would call the existing ResolveFork, which already verifies that the candidate shares the same genesis block before replacing the local chain. A production implementation would extend this with network-level peer discovery, block propagation, and consensus handling. The core validation and fork-choice logic in this project would not need to change much — the main new work would be the networking layer (peer discovery, message passing, and re-broadcast) around it.

---

# 8. Known Limitations

Although the simulator implements the major blockchain concepts required for the assessment, it remains a simplified educational project.

Current limitations include:

No peer-to-peer networking or distributed node communication.
No transaction fees or mining rewards.
Wallets are stored locally without encryption or password protection.
The blockchain operates as a single-node system without real network consensus.
Merkle roots are included in block hashes, but transaction data is also hashed directly, making the Merkle root partially redundant in the current design.

---

# 9. Project Challenges and Solutions

## Challenge 1: Maintaining Blockchain Integrity

**Problem:** Changing block data could make the blockchain invalid without anyone noticing.

**Solution:** Implemented SHA-256 hashing, a Merkle root per block, and a chain-wide validation routine that recomputes and cross-checks both on every block (Section 7.1 demonstrates this catching a tampered transaction).

---

## Challenge 2: Preventing Invalid Transactions

**Problem:** Users could attempt transactions without enough balance, or with non-positive amounts.

**Solution:** Implemented balance verification (`ValidateTransaction`) and transaction replay validation (`CalculateBalances`/`ValidateBalances`) so every transaction is checked both when it's added to the pending pool and again when the whole chain is validated.

---

## Challenge 3: Improving Mining Performance

**Problem:** Single-threaded mining is slower than it needs to be on a multi-core machine.

**Solution:** Implemented concurrent mining using goroutines, an atomic attempt counter, a mutex-protected result, and context cancellation so multiple workers can search different parts of the nonce space in parallel and stop cleanly once one succeeds.

---

## Challenge 4: Handling Blockchain Forks

**Problem:** Multiple valid chains may exist (e.g. two competing histories).

**Solution:** Implemented a longest-valid-chain rule (`ResolveFork`) that only accepts a candidate chain if it is both longer and passes full chain validation. As noted in Section 8, this rule doesn't yet check for shared ancestry with the current chain, which is a planned follow-up.

---

# 10. Conclusion

The Toy Blockchain and Ledger Simulator implements the major concepts of blockchain systems: block/genesis structure, SHA-256 and Merkle-root hashing, Proof-of-Work consensus (both sequential and concurrent), signed transactions and ledger replay, difficulty retargeting, fork resolution, and JSON persistence, all backed by a passing unit test suite.

During development, several implementation issues were identified and corrected, including CLI mining consistency, signature encoding, fork validation, and difficulty limits. These improvements increased the reliability and correctness of the final implementation.

This project provided practical experience in designing and implementing a blockchain-based backend system in Go, including the value of testing the exact path a user actually exercises, not just the underlying library functions.