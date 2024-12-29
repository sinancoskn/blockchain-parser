# Project: Ethereum Transaction Notifier

This project is an Ethereum transaction notification system that monitors blockchain transactions for subscribed addresses and provides a notification mechanism. It includes a parser, storage, and configuration management.

## Features

- Monitors the Ethereum blockchain for transactions.
- Allows subscribing and unsubscribing to specific Ethereum addresses.
- Stores transaction data in an in-memory storage system.
- Exposes HTTP endpoints for interaction.
- Configurable via `.env` file.

## Project Structure

```
project-root/
├── config/                # Configuration management
│   └── config.go
├── internal/              # Core application logic
│   ├── app/               # Application initialization
│   ├── ethereum/          # Ethereum JSON-RPC client
│   ├── parser/            # Blockchain parser implementation
│   └── storage/           # In-memory storage system
├── test/                  # Test files
│   ├── integration/       # Integration tests
│   └── unit/              # Unit tests
├── .env                   # Environment variables
├── go.mod                 # Go module file
├── main.go                # Application entry point
└── README.md              # Project documentation
```

## Prerequisites

- [Go](https://golang.org/dl/) 1.18+
- Geth (Go Ethereum) or another Ethereum node for testing

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/yourproject.git
   cd yourproject
   ```

2. Create a `.env` file in the project root:
   ```env
   ETHEREUM_RPC_URL=http://127.0.0.1:8545
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the Ethereum node locally (example with Geth):
   ```bash
   geth --dev --http --http.api personal,eth,net,web3 --http.addr 127.0.0.1 --datadir ./eth_data
   ```

## Running the Application

Start the application:
```bash
go run main.go
```

## HTTP Endpoints

- **Get Current Block**
  - **GET** `/block/current`
  - Returns the current block being monitored.

- **Subscribe to an Address**
  - **POST** `/subscribe`
  - Body: `{ "address": "<ethereum_address>" }`

- **Unsubscribe from an Address**
    - **POST** `/unsubscribe`
    - Body: `{ "address": "<ethereum_address>" }`

- **Get Transactions for an Address**
  - **GET** `/transactions?address=<ethereum_address>`

## Tests

### Run All Tests
```bash
go test ./...
```

### Run Unit Tests Only
```bash
go test ./test/unit/...
```

### Run Integration Tests Only
```bash
go test ./test/integration/...
```