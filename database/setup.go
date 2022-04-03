package database

import (
	"github.com/mehdisadeghidev/go-mongodb-template/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Database

func Setup() error {
	client, err := NewMongoClient(config.Database.Uri)

	if err != nil {
		return err
	}

	DB = client.Database(config.Database.Name)

	err = createIndexes()

	return err
}
