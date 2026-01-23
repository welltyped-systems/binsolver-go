package binsolver_test

import (
	"context"
	"net/http"

	"github.com/welltyped-systems/binsolver-go"
)

func ExampleClientWithResponses_PostPack() {
	client, err := binsolver.NewClientWithResponses(
		"https://api.binsolver.com",
		binsolver.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("x-api-key", "test-api-key")
			return nil
		}),
	)
	if err != nil {
		panic(err)
	}

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

	_ = req
	_ = client

}

func ptr[T any](v T) *T {
	return &v
}
