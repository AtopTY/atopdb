package atopdb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *DB) Error(err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	collection := d.db.Collection("errorLogs")
	collection.InsertOne(ctx, bson.M{
		"error":  err.Error(),
		"lastTS": time.Now(),
	})
}
