package mongox

import (
	"context"

	"github.com/go-bamboo/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client = mongo.Client

func MustConnect(c *Conf) *Client {
	client, err := Connect(c)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func Connect(c *Conf) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout.AsDuration())
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.Source))
	if err != nil {
		return nil, err
	}
	return client, nil
}
