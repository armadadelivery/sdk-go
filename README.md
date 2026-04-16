# armada (Go SDK)

```bash
go get github.com/armadadelivery/sdk-go
```

```go
package main

import (
    "context"
    "fmt"
    "os"

    armada "github.com/armadadelivery/sdk-go"
)

type Wallet struct {
    Balance  float64 `json:"balance"`
    Currency string  `json:"currency"`
}

func main() {
    c := armada.NewClient(armada.Options{
        APIKey:    os.Getenv("ARMADA_API_KEY"),
        APISecret: os.Getenv("ARMADA_API_SECRET"),
    })

    resp, err := c.Get(context.Background(), "/v2/wallet", nil)
    if err != nil { panic(err) }

    var w Wallet
    if err := armada.DecodeJSON(resp, &w); err != nil { panic(err) }

    fmt.Printf("balance: %.2f %s (remaining: %v)\n", w.Balance, w.Currency, resp.RateLimit.Remaining)
}
```

## What it handles

- HMAC-SHA256 request signing
- `RateLimit` struct on every `*Response`
- Convenience methods: `Get`, `PostJSON`, `PutJSON`, `Delete`, or the full `Do(ctx, method, path, query, body)`
- `DecodeJSON` helper that errors on 4xx/5xx with the decoded body

## Generated typed resources

Typed resource clients generated from the OpenAPI spec are rolling out
alongside the `scripts/regen.sh` pipeline. Until then, build request
bodies per the spec at
[docs.armadadelivery.com/v2](https://docs.armadadelivery.com/v2).
