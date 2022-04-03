package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func createIndexes() error {
	var err error
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)

	_, err = DB.Collection("users").Indexes().CreateMany(
		context.TODO(),
		[]mongo.IndexModel{
			{
				Keys: bson.D{{Key: "email", Value: 1}},
				Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{
					"email": bson.M{
						"$type":   "string",
						"$exists": true,
					},
				}),
			},
		},
		opts,
	)

	return err
}
