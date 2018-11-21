package main

import (
	"context"
	"log"

	"cloud.google.com/go/spanner"
)

func CreateClient(ctx context.Context, db string, spannerMinOpened uint64) *spanner.Client {
	ctx, span := startSpan(ctx, "createClient")
	defer span.End()

	o := spanner.ClientConfig{
		SessionPoolConfig: spanner.SessionPoolConfig{
			MinOpened: spannerMinOpened,
		},
	}
	dataClient, err := spanner.NewClientWithConfig(ctx, db, o)
	if err != nil {
		log.Fatal(err)
	}

	return dataClient
}

func WarmUp(ctx context.Context, client *spanner.Client) error {
	ctx, span := startSpan(ctx, "warmUp")
	defer span.End()

	err := client.Single().Query(ctx, spanner.NewStatement("SELECT 1")).Do(func(r *spanner.Row) error {
		return nil
	})
	return err
}
