# BinSolver Go SDK

Official Go SDK for the [BinSolver API](https://binsolver.com).

## Installation

```bash
go get github.com/welltyped-systems/binsolver-go
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/welltyped-systems/binsolver-go"
)

func main() {
	// Initialize the client
	client, err := binsolver.NewClientWithResponses(
		"https://api.binsolver.com",
		binsolver.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("x-api-key", "your-api-key")
			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Check health
	health, err := client.GetHealthWithResponse(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if health.JSON200 != nil {
		fmt.Println("API is healthy")
	}

	// Prepare request
	req := binsolver.PackRequest{
		Objective: ptr(binsolver.MinBins),
		Items: []binsolver.ItemInput{
			{
				Id:       ptr("item-1"),
				W:        5,
				H:        5,
				D:        5,
				Quantity: ptr(12),
			},
		},
		Bins: []binsolver.BinInput{
			{
				Id:       ptr("box-small"),
				W:        10,
				H:        10,
				D:        10,
				Quantity: ptr(10),
			},
		},
	}

	// Pack items
	resp, err := client.PostPackWithResponse(
		context.Background(),
		&binsolver.PostPackParams{},
		binsolver.PostPackJSONRequestBody(req)
	)
	if err != nil {
		log.Fatal(err)
	}

	if resp.JSON200 != nil {
		fmt.Printf("Placed %d items in %d bins.\n", resp.JSON200.Stats.Placed, resp.JSON200.Stats.BinsUsed)
		for _, bin := range resp.JSON200.Bins {
			fmt.Printf("Bin %s (%s): %d items\n", bin.BinId, bin.TemplateId, len(bin.Placements))
		}
	} else if resp.JSON400 != nil {
		fmt.Printf("Error: %s\n", resp.JSON400.Error.Message)
	} else {
		fmt.Printf("Unexpected status: %d\n", resp.StatusCode())
	}
}

func ptr[T any](v T) *T {
	return &v
}
```

## Features

- **Fully Typed:** Generated from OpenAPI specification.
- **Modern:** Context-aware and strongly typed.
- **Easy to Use:** simple interface for the BinSolver API.

## License

MIT
