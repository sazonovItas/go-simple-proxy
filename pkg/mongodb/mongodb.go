package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, uri string) (*mongo.Client, error) {
	const op = "pkg.mongodb.Connect"

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	db, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: failed connect to db: %w", op, err)
	}

	return db, nil
}
