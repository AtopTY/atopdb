package atopdb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// TaTableSettings the table setting schema (colleciton: tableSeting)
type TableSettings struct {
	Name             string   `bson:"name"`
	Unique           []string `bson:"unique"`
	SearchableFields []string `bson:"searchableFields"`
}

// getTableSettings
func (d *DB) getTableSettings(c string) (TableSettings, error) {
	collection := d.db.Collection("tableSetting")
	var ts TableSettings
	err := collection.FindOne(context.Background(), bson.M{"collection": c}).Decode(&ts)
	if err != nil {
		return TableSettings{}, err
	}
	return ts, nil
}

func (d *DB) getSearchFields(c string) []string {
	searchOption := d.db.Collection("tableSetting")
	var searchFields TableSettings
	err := searchOption.FindOne(context.Background(), bson.M{"collection": c}).Decode(&searchFields)
	if err != nil {
		return []string{}
	}
	return searchFields.SearchableFields
}
