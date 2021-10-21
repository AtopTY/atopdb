package atopdb

import "go.mongodb.org/mongo-driver/bson"

func GetDatabaseStatus() (bson.M, error) {
	return bson.M{}, nil
}
