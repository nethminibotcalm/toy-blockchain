# Toy Blockchain and Ledger Simulator

## Software Engineering Internship Assessment Report

## 1. Introduction

This project is a command-line based Toy Blockchain and Ledger Simulator developed using **Go 1.22+**.

The objective of this project is to implement and demonstrate the fundamental concepts of blockchain technology, including block creation, cryptographic hashing, Proof of Work mining, transaction management, blockchain validation, persistence, and advanced blockchain features.

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
| -------------------- | --------------------------------- |
| Go 1.22+             | Backend implementation            |
| SHA-256              | Block hashing                     |
| ECDSA                | Digital signatures                |
| JSON                 | Blockchain and wallet persistence |
| Goroutines           | Concurrent mining                 |
| Go Testing Framework | Automated testing                 |

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
* Previous hash relationship
* Block ordering
* Timestamp ordering
* Proof of Work difficulty
* Transaction validity
* Genesis block correctness

---

# 4.3 Proof of Work Mining

The project implements a Proof of Work consensus mechanism.

Mining process:

1. A block is created with pending transactions.
2. The nonce value is changed repeatedly.
3. SHA-256 hash is calculated.
4. Mining continues until the hash satisfies the difficulty requirement.

Implemented features:

* Configurable difficulty
* Mining attempt counting
* Mining time measurement
* Concurrent mining support

---

# 4.4 Concurrent Mining

Concurrent mining was implemented using Go concurrency features.

Implementation:

* Multiple goroutines work simultaneously.
* Each worker searches a different nonce range.
* Workers use atomic counters to track attempts.
* Mutex protects shared mining results.
* Context cancellation stops remaining workers after success.

Example:

Worker 0:

```
0,4,8,12...
```

Worker 1:

```
1,5,9,13...
```

This improves mining performance by utilizing multiple CPU workers.

---

# 4.5 Transactions and Ledger

The ledger system manages transactions and balances.

Implemented features:

* Sender and receiver transactions
* Integer-based transaction amounts
* Pending transaction pool
* Balance calculation from blockchain history
* Transaction validation
* Double-spending prevention

Balances are not stored permanently. Instead, they are recalculated from blockchain transaction history.

---

# 4.6 Digital Signatures and Wallets

ECDSA-based digital signatures were implemented to authenticate transactions.

Implemented features:

* Wallet key generation
* Private and public key handling
* Transaction signing
* Signature verification
* Invalid signature rejection

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

This prevents unauthorized transaction modification.

---

# 4.7 Merkle Root Implementation

Merkle trees were implemented to summarize block transactions efficiently.

Instead of directly hashing the complete transaction list, transactions are converted into hashes and combined into a Merkle tree.

Benefits:

* Faster transaction verification
* Efficient transaction integrity checking
* Reduced dependency on raw transaction lists

Merkle root is included in block hashing, meaning any transaction modification changes the block hash.

---

# 4.8 Difficulty Retargeting

Automatic difficulty adjustment was implemented.

The system monitors block generation time and adjusts mining difficulty.

Purpose:

* Maintain consistent block creation time
* Increase difficulty when blocks are mined too quickly
* Reduce difficulty when mining becomes too slow

This simulates real blockchain difficulty adjustment mechanisms.

---

# 4.9 Fork Resolution

Fork resolution was implemented using the longest valid chain rule.

Process:

1. Receive a competing blockchain.
2. Check if the candidate chain is longer.
3. Validate the candidate chain.
4. Replace the current chain if valid.

Invalid or shorter chains are rejected.

---

# 5. Data Persistence

Blockchain data is stored using JSON format.

Implemented features:

* Save blockchain state
* Load blockchain state
* Restore balances after restart
* Validate loaded blockchain data

---

# 6. Testing

The project includes automated tests covering:

* SHA-256 hash generation
* Merkle root calculation
* Blockchain validation
* Tamper detection
* Mining difficulty
* Concurrent mining
* Difficulty adjustment
* Fork resolution
* Transaction validation
* Double-spending prevention
* Persistence
* Digital signature verification

Testing command:

```
go test ./...
```

All tests pass successfully.

---

# 7. Project Challenges and Solutions

## Challenge 1: Maintaining Blockchain Integrity

Problem:

Changing block data could make the blockchain invalid.

Solution:

Implemented SHA-256 hashing and chain validation checks.

---

## Challenge 2: Preventing Invalid Transactions

Problem:

Users could attempt transactions without enough balance.

Solution:

Implemented balance verification and transaction replay validation.

---

## Challenge 3: Improving Mining Performance

Problem:

Single-threaded mining was slower.

Solution:

Implemented concurrent mining using goroutines and synchronization techniques.

---

## Challenge 4: Handling Blockchain Forks

Problem:

Multiple valid chains may exist.

Solution:

Implemented longest-valid-chain fork resolution.

---

# 8. Conclusion

The Toy Blockchain and Ledger Simulator successfully implements the major concepts of blockchain systems.

The project demonstrates:

* Blockchain structure design
* Cryptographic hashing
* Proof of Work consensus
* Secure transaction processing
* Digital signatures
* Merkle tree optimization
* Concurrent programming
* Difficulty adjustment
* Fork resolution
* Automated testing

This project provided practical experience in designing and implementing a blockchain-based backend system using Go.
