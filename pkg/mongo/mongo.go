package mongo

import (
	"context"
	"github.com/bebraRPC/microservice/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	connTimeout     = 30 * time.Second
	maxConnIdleTime = 3 * time.Minute
	minPoolSize     = 20
	maxPoolSize     = 300
)

func NewMongoDBConn(ctx context.Context, cfg *config.Config) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDB.URI).
		SetAuth(options.Credential{
			Username: cfg.MongoDB.User,
			Password: cfg.MongoDB.Password,
		}).
		SetConnectTimeout(connTimeout).
		SetMaxConnIdleTime(maxConnIdleTime).
		SetMinPoolSize(minPoolSize).
		SetMaxPoolSize(maxPoolSize))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
