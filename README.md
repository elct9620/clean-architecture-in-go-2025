Clean Architecture in Go 2025
===

This is example code for the series [Clean Architecture in Go](https://blog.aotoki.me/series/clean-architecture-in-go/).

## gRPCurl Example

### Lookup Order

```shell
grpcurl -plaintext -d '{"id": "d6ec-1c06-4e0b-aa57-9b2335fc56c3"}' localhost:8080 order.OrderService/LookupOrder
```

### Place Order

```shell
grpcurl -plaintext -d '{"name": "Aotoki", "items": [{"name": "Apple", "quantity": 1, "unit_price": 10 }]}' localhost:8080 order.OrderService/PlaceOrder
```
