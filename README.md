To run the ledger provider:

```make run```

then cURL:

```
curl -X POST http://localhost:8080/accept \
  -H "Content-Type: application/json" \
  -d '{
    "id": "block-abc123",
    "parent_ids": ["block-xyz789"],
    "timestamp": "2025-06-19T12:00:00Z",
    "transactions": ["tx1", "tx2"],
    "signature": "mock-signature"
  }'
```
