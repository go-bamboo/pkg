package mongox

import (
	"bls/pkg/log"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MustConnect(c *Conf) *mongo.Client {
	client, err := Connect(c)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func Connect(c *Conf) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout.AsDuration())
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.Source))
	if err != nil {
		return nil, err
	}
	return client, nil
}
