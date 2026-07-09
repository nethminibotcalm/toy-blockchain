# Toy Blockchain and Ledger Simulator – Research Report

## 1. Introduction

This project implements a command-line (CLI) Toy Blockchain and Ledger Simulator using Go 1.22+. The objective is to demonstrate the fundamental concepts behind blockchain technology, including cryptographic hashing, Proof of Work (PoW), transaction validation, blockchain integrity, and persistent storage.

Unlike production blockchain systems, this implementation focuses on educational purposes by providing a simplified blockchain that runs on a single machine without networking or distributed consensus.

The system allows users to add transactions, mine new blocks, validate the blockchain, view account balances, and persist the blockchain to a JSON file.

---

# 2. System Design

The project is organized into several packages, each with a specific responsibility.

* **block** – Defines the structure of a block and implements SHA-256 hashing.
* **blockchain** – Manages the blockchain, mining, validation, balance calculation, printing, and persistence.
* **ledger** – Validates transactions and manages account balances.
* **main.go** – Provides the command-line interface and coordinates the entire application.

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

Each block is connected to the previous block through the `PreviousHash` field, forming a chain of blocks. Any modification to a block changes its hash, causing all subsequent blocks to become invalid.

---

# 4. SHA-256 Hashing

The project uses the SHA-256 cryptographic hash function to generate a unique identifier for each block.

The hash is calculated using:

* Block Index
* Timestamp
* Transactions
* Previous Hash
* Nonce

A key property of SHA-256 is that even a small change to the input produces a completely different output hash. This characteristic makes it possible to detect any modification to blockchain data.

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

Mining repeatedly changes the block's nonce until the calculated SHA-256 hash begins with the required number of leading zeros.

Current mining difficulty:

```
Difficulty = 4
```

Example of a valid hash:

```
00005a1bbfe0f1f8139082808cb1357da5bf6acf0825b988920a87c3276e238a
```

Increasing the mining difficulty requires more nonce attempts before a valid hash can be found, increasing computational effort.

---

# 6. Transaction Validation

Before transactions are added to a block, the ledger validates each transaction.

The following rules are applied:

* Transaction amount must be greater than zero.
* The sender must exist.
* The sender must have sufficient balance.

Only valid transactions are included in newly mined blocks.

---

# 7. Blockchain Validation

The blockchain can be validated at any time using the validation command.

Validation performs the following checks:

1. Recalculate each block's hash and compare it with the stored hash.
2. Verify that each block's `PreviousHash` matches the previous block's hash.
3. Verify that each block satisfies the configured Proof of Work difficulty.

If any of these checks fail, the blockchain is considered invalid.

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

Changing transaction data modifies the calculated block hash. Since the stored hash no longer matches the recalculated hash, blockchain validation fails. This demonstrates how cryptographic hashing protects the integrity of blockchain data.

---

# 9. Difficulty vs. Mining Effort

| Difficulty      | Expected Mining Effort |
| --------------- | ---------------------- |
| 2 leading zeros | Low                    |
| 3 leading zeros | Medium                 |
| 4 leading zeros | High                   |
| 5 leading zeros | Very High              |

As the required number of leading zeros increases, the probability of finding a valid hash decreases. Consequently, more nonce values must be tested before successfully mining a block.

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

Running the tests:

```
go test ./...
```

All implemented tests complete successfully.

---

# 13. Conclusion

This project successfully demonstrates the core principles of blockchain technology through a simplified command-line application developed in Go. The implementation includes block creation, cryptographic hashing, Proof of Work mining, transaction validation, blockchain validation, persistence, and automated testing.

Although simplified, the project illustrates the essential mechanisms used by real blockchain systems to ensure data integrity, detect tampering, and maintain a secure sequence of transactions. It provides a strong foundation for understanding more advanced blockchain technologies such as distributed consensus, digital signatures, peer-to-peer networking, and smart contracts.
