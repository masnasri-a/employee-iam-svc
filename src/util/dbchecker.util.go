package util

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckFieldDB(client *mongo.Client, key string, value string, collection string) (bool, error) {
	// Check if the field exists in the database
	coll := client.Database("emp-iam").Collection(collection)
	cursor, err := coll.Find(context.Background(), bson.M{key: value})
	if err != nil {
		return false, err
	}
	defer cursor.Close(context.Background())
	if cursor.Next(context.Background()) {
		return true, nil
	}
	return false, nil
}
