## Getting Started

```
https://github.com/LGROW101/Block-Blockchain.git

cd Block-Blockchain

go run main.go
```

## Blockchain API Testing

# Create first wallet

```
$ curl -X POST http://localhost:8080/wallet

{"address":"abc123..."}
```

# Create second wallet

```
$ curl -X POST http://localhost:8080/wallet

{"address":"def456..."}
```

# Create a transaction

```
$ curl -X POST http://localhost:8080/transaction \
     -H "Content-Type: application/json" \
     -d '{"from":"abc123...","to":"def456...","amount":10}'

{"message":"Transaction created and new block mined"}
```

# Validate the blockchain

```
$ curl http://localhost:8080/validate

{"isValid":true}

```

# View the entire blockchain

```
$ curl http://localhost:8080/blockchain


[{...},{...}]  # This will show the full blockchain data

```
