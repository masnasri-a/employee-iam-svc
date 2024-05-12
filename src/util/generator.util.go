package util

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateID() string {
	objectID := primitive.NewObjectID()
	objectIDString := objectID.Hex()
	return objectIDString
}

func GetTimeNow() string {
	return primitive.NewDateTimeFromTime(time.Now()).Time().String()
}
