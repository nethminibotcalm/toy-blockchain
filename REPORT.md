# Toy Blockchain and Ledger Simulator – Research Report

## 1. Introduction

This project implements a command-line (CLI) Toy Blockchain and Ledger Simulator in Go. The objective is to demonstrate the fundamental concepts behind blockchain technology, including cryptographic hashing, Proof of Work (PoW), transaction validation, blockchain integrity, and persistent storage.

Unlike production blockchain systems, this implementation focuses on educational purposes by providing a simplified blockchain that runs on a single machine without networking or distributed consensus.

The system allows users to add transactions, mine new blocks, validate the blockchain, view account balances, and persist the blockchain to a JSON file.

---

# 2. System Design

The project is organized into several packages, each with a specific responsibility.

* **block** – Defines the block structure and SHA-256 hashing logic.
* **blockchain** – Manages the chain, mining, validation, balance calculation, printing, and persistence.
* **ledger** – Validates transactions and tracks account balances.
* **main.go** – Provides the command-line interface and coordinates the application flow.

This modular structure improves readability, maintainability, and testing while avoiding unnecessary dependencies between packages.

---

# 3. Block Structure

Each block contains the following information:

* Block Index
* Timestamp
* List of Transactions
* Previous Block Hash
* Nonce
* Current Block Hash

Each block is connected to the previous block through the `PreviousHash` field, forming a chain of blocks. The genesis block is created with a fixed timestamp of `0` and a fixed previous-hash value so that the first block is deterministic across runs.

---

# 4. SHA-256 Hashing

The project uses the SHA-256 cryptographic hash function to generate a unique identifier for each block.

The hash is calculated from a labeled string that includes:

* Block Index
* Timestamp
* Transactions
* Previous Hash
* Nonce

A key property of SHA-256 is that even a small change to the input produces a completely different output hash. The implementation also separates the block fields with labels and delimiters, so different blocks do not collapse into the same hash input string.

For example:

Original transaction:

```
Alice -> Bob : 20
```

Modified transaction:

```
Alice -> Bob : 200
```

Although only one digit changes, the resulting hash becomes completely different.

---

# 5. Proof of Work

To add a new block, the blockchain performs Proof of Work (PoW).

Mining repeatedly changes the block's nonce until the calculated SHA-256 hash begins with the required number of leading zeros. In the current implementation, the blockchain is initialized with a difficulty of `4`, and mining prints both the number of attempts and the elapsed time.

Example of a valid hash:

```
00005a1bbfe0f1f8139082808cb1357da5bf6acf0825b988920a87c3276e238a
```

Increasing the mining difficulty requires more nonce attempts before a valid hash can be found, increasing computational effort.

---

# 6. Transaction Validation

Before transactions are added to a block, the ledger validates each transaction.

The following rules are applied:

* Transaction amount must be a positive integer.
* The sender must exist.
* The sender must have sufficient balance.

Only valid transactions are included in newly mined blocks.

---

# 7. Blockchain Validation

The blockchain can be validated at any time using the validation command, and chains loaded from `chain.json` are validated before use.

Validation performs the following checks:

1. Recalculate the genesis block hash and compare it with the stored hash.
2. Recalculate each block's hash and compare it with the stored hash.
3. Verify that each block's `PreviousHash` matches the previous block's hash.
4. Verify that each block index increases by one.
5. Verify that timestamps do not move backward.
6. Verify that each block satisfies the configured Proof of Work difficulty.
7. Replay balances across the chain to detect invalid transaction flows.

If any of these checks fail, the blockchain is considered invalid, and the validation error identifies the block that failed.

---

# 8. Tamper Experiment

To demonstrate blockchain integrity, a tampering experiment was performed.

### Original Blockchain

Transaction:

```
Alice -> Bob : 20
```

Validation result:

```
Blockchain valid: true
```

### Tampered Blockchain

The transaction amount was manually changed:

```
Alice -> Bob : 999
```

Validation result:

```
Blockchain valid: false
```

### Observation

Changing transaction data modifies the calculated block hash. If the tampered block is not mined again correctly, validation fails because the stored hash, the proof-of-work prefix, or the replayed balances no longer match the chain. This demonstrates how cryptographic hashing protects the integrity of blockchain data.

---

# 9. Difficulty vs. Mining Effort

The blockchain runs with a fixed difficulty of 4 by default, but the mining function accepts difficulty as a parameter. To study the relationship between difficulty and effort, the same block was mined at difficulty levels 1 through 6 using the project's own `MineBlock` function, and the attempt count and elapsed time were recorded for each run.

| Difficulty | Attempts  | Time        |
| ---------- | --------- | ----------- |
| 1          | 5         | 0.28 ms     |
| 2          | 5         | 0.01 ms     |
| 3          | 5         | 0.01 ms     |
| 4          | 37,931    | 92.5 ms     |
| 5          | 722,077   | 1.78 s      |
| 6          | 4,949,561 | 11.85 s     |

At low difficulty (1-3 leading zero hex digits) a valid hash is found almost immediately, since roughly 1 in 16 hashes already satisfies a single leading zero. From difficulty 4 onward the cost grows sharply: each additional required hex digit multiplies the expected number of attempts by roughly 16, since each hex digit has 16 possible values and the hash function behaves like a uniform random source. The growth is therefore exponential in the difficulty, not linear — going from difficulty 4 to 6 increases attempts by roughly 130x, consistent with 16^2 = 256 in expectation. This matches the theoretical model of proof-of-work: doubling the difficulty target does not double the work, it multiplies it.

---

# 10. Proof of Work vs. Proof of Stake

| Proof of Work                          | Proof of Stake                                |
| -------------------------------------- | --------------------------------------------- |
| Uses computational work to mine blocks | Uses validators based on cryptocurrency stake |
| Requires significant computation       | Requires significantly less computation       |
| Higher energy consumption              | Lower energy consumption                      |
| Used by Bitcoin                        | Used by modern Ethereum                       |

Proof of Work provides strong security through computational effort but consumes more energy than Proof of Stake.

---

# 11. Toy Blockchain vs. Production Blockchain

| Toy Blockchain             | Production Blockchain            |
| -------------------------- | -------------------------------- |
| Single-node application    | Distributed peer-to-peer network |
| JSON file persistence      | Distributed replicated storage   |
| No digital signatures      | Public/private key cryptography  |
| Simple Proof of Work       | Advanced consensus algorithms    |
| Educational implementation | Highly secure production systems |
| No networking              | Full peer-to-peer communication  |

This implementation focuses on demonstrating blockchain concepts rather than providing a production-ready blockchain platform.

---

# 12. Testing

The project includes automated unit tests covering:

* Deterministic SHA-256 hashing
* Blockchain validation
* Tamper detection
* Ledger validation
* Balance calculation
* Mining difficulty
* Double-spending prevention
* Persistence round-trip behavior

Running the tests:

```
go test ./...
```

All implemented tests complete successfully.

---

# 13. Conclusion

This project demonstrates the core principles of blockchain technology through a simplified command-line application developed in Go. The implementation includes deterministic block creation, cryptographic hashing, Proof of Work mining, transaction validation, blockchain validation, persistence, and automated testing.

Although simplified, the project illustrates the essential mechanisms used by real blockchain systems to ensure data integrity, detect tampering, and maintain a secure sequence of transactions. It provides a foundation for understanding more advanced blockchain technologies such as distributed consensus, digital signatures, peer-to-peer networking, and smart contracts.
