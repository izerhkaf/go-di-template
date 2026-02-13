package database

// import (
// 	"context"
// 	"crypto/tls"
// 	"homeez/internal/config"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func NewMongo(cfg *config.Config) *mongo.Database {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	client, err := mongo.Connect(ctx, options.Client().
// 		ApplyURI(cfg.MongoURI).
// 		SetTLSConfig(&tls.Config{InsecureSkipVerify: true}),
// 	)
// 	if err != nil {
// 		return nil
// 	}

// 	return client.Database(cfg.MongoDBName)
// }
