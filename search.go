package atopdb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// FuzzySearch search key in collection
func (d *DB) FuzzySearch(c, searchString string) ([]bson.M, error) {
	textFields := d.getSearchFields(c)

	if len(textFields) <= 0 {
		return nil, fmt.Errorf("Can not search %s without fields", c)
	}

	var queryCondition = []bson.M{}
	for _, v := range textFields {
		queryCondition = append(queryCondition, bson.M{v: bson.M{"$regex": searchString}})
	}
	var query = bson.M{"$or": queryCondition}

	collection := d.db.Collection(c)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, query)
	defer cur.Close(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var data []bson.M
	for cur.Next(ctx) {
		var item bson.M
		err := cur.Decode(&item)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}
