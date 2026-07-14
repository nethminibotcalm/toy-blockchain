# Toy Blockchain and Ledger Simulator

A command-line blockchain simulator developed using **Go 1.22+** as part of a Software Engineering Internship assessment.

This project demonstrates the core concepts of blockchain technology including block creation, SHA-256 hashing, Proof of Work mining, transaction validation, ledger management, blockchain validation, persistence, and testing.

---

## Features

### Blockchain

- Genesis block creation
- Deterministic genesis block
- Block linking using previous block hashes
- SHA-256 hash generation
- Structured hash calculation
- Blockchain validation
- Tamper detection
- Detailed validation error reporting
- Validation after loading blockchain data

### Proof of Work

- Configurable mining difficulty
- Nonce-based mining
- Hash difficulty verification
- Mining attempt count reporting
- Mining execution time reporting

### Transactions and Ledger

- Transaction model with sender, receiver, and amount
- Integer-based transaction amounts
- Pending transaction pool
- Transaction validation
- Balance calculation from blockchain history
- Double-spending prevention
- Ledger state derived from blockchain data

### Persistence

- Save blockchain data into JSON format
- Load blockchain data from JSON file
- Validate blockchain after loading
- Maintain blockchain state after restarting application

### Testing

Implemented tests for:

- Hash generation
- Blockchain validation
- Tamper detection
- Mining difficulty
- Ledger validation
- Persistence
- Double-spending prevention

---

## Project Structure

```
toy-blockchain/
│
├── main.go
├── chain.json
├── go.mod
├── README.md
├── REPORT.md
│
├── block/
│   ├── block.go
│   ├── hash.go
│   └── hash_test.go
│
├── blockchain/
│   ├── balance.go
│   ├── balance_validation.go
│   ├── blockchain.go
│   ├── double_spend_test.go
│   ├── mining.go
│   ├── mining_test.go
│   ├── print.go
│   ├── storage.go
│   ├── storage_test.go
│   ├── tamper_test.go
│   ├── validate.go
│   └── validate_test.go
│
└── ledger/
    ├── ledger.go
    ├── ledger_test.go
    └── transaction.go
```

---

## Requirements

- Go 1.22 or later

Check Go installation:

```bash
go version
```

---

## Installation

Clone the repository:

```bash
git clone <repository-url>
```

Navigate into the project:

```bash
cd toy-blockchain
```

Install dependencies:

```bash
go mod tidy
```

---

## Running the Application

### Add a Transaction

Command:

```bash
go run . add Alice Bob 20
```

Example output:

```
Transaction added
```

The transaction is added to the pending transaction pool.

### Mine Pending Transactions

Command:

```bash
go run . mine
```

During mining:

- Pending transactions are validated
- A new block is created
- Proof of Work is performed
- The mined block is added to the blockchain
- Blockchain data is saved to `chain.json`

Example output:

```
Mining attempts: 1619
Mining time: 5.462ms
Mining completed
```

### Print Blockchain

Command:

```bash
go run . print
```

Displays:

- Block index
- Transactions
- Nonce
- Previous hash
- Block hash

Example output:

```
Index: 1
Transactions: [{Alice Bob 20}]
Nonce: 1618
Previous Hash: ...
Hash: 0000....
```

### Validate Blockchain

Command:

```bash
go run . validate
```

Example output:

```
Blockchain is valid
```

If validation fails:

```
Blockchain is invalid: block 2: invalid hash
```

Validation checks:

- Hash correctness
- Previous block connection
- Block order
- Timestamp order
- Proof of Work difficulty
- Transaction balance validity

### View Balances

Command:

```bash
go run . balance
```

Example output:

```
Balances: map[Alice:80 Bob:120 Charlie:100]
```

Balances are calculated from blockchain transaction history.

---
## Concurrent Mining

The blockchain supports concurrent Proof of Work mining using Go goroutines.

Implementation details:
- Multiple workers search different nonce ranges.
- sync.WaitGroup manages worker completion.
- context cancellation stops remaining workers after a valid nonce is found.
- atomic counters track mining attempts.
- mutex protects shared mining results.

---

## Example Workflow

1. Add transaction

   ```bash
   go run . add Alice Bob 20
   ```

2. Mine transaction

   ```bash
   go run . mine
   ```

3. View blockchain

   ```bash
   go run . print
   ```

4. Validate blockchain

   ```bash
   go run . validate
   ```

5. View balances

   ```bash
   go run . balance
   ```

---

## How It Works

1. A user creates a transaction.
2. The transaction is stored in the pending transaction pool.
3. Pending transactions are validated before mining.
4. A new block is created with valid transactions.
5. Proof of Work searches for a valid nonce.
6. The block is added to the blockchain.
7. Blockchain data is stored in `chain.json`.
8. When restarting, blockchain data is loaded and validated.
9. Account balances are calculated from blockchain history.

---

## Proof of Work

This project uses a simple Proof of Work algorithm.

Mining changes the block nonce until the SHA-256 hash satisfies the required difficulty.

Example (difficulty: 4), a valid hash:

```
00005a1bbfe0f1f8139082808cb1357da5bf6acf0825b988920a87c3276e238a
```

Higher difficulty requires more mining attempts.

---

## Persistence

Blockchain data is stored in:

```
chain.json
```

The saved blockchain is loaded during application startup and validated before use.

---

## Running Tests

Run all tests:

```bash
go test ./...
```

Example output:

```
ok      toy-blockchain/block
ok      toy-blockchain/blockchain
ok      toy-blockchain/ledger
```

---

## Current Limitations

This project is created for educational purposes and does not include:

- Peer-to-peer networking
- Digital signatures
- Wallet management
- Distributed consensus
- Multiple blockchain nodes
- Smart contracts
- Merkle trees

---

## Future Improvements

Possible improvements:

- Digital wallet implementation
- Cryptographic signatures
- REST API support
- Web interface
- Multiple blockchain nodes
- Adaptive mining difficulty
- Merkle tree implementation

---

## Author

Developed as part of a Software Engineering Internship take-home assessment using Go.