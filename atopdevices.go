package atopdb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const AtopDeviceCollection = "atopDevices"

func (d *DB) GetDeviceInformation() ([]bson.M, error) {

	data, err := d.Read(AtopDeviceCollection, bson.M{})
	return data, err
}

func (d *DB) UpdateDeviceInformation(data bson.M) (bson.M, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	collection := d.db.Collection(AtopDeviceCollection)

	var doc = bson.M{}
	res := collection.FindOneAndUpdate(ctx,
		bson.M{"mac": data["mac"]},
		bson.M{"$set": data},
		options.FindOneAndUpdate().SetUpsert(true).SetBypassDocumentValidation(true))

	decodeErr := res.Decode(doc)
	fmt.Printf("FindoneAndUpdate result.err = %q resdecode=%q\n", res.Err(), doc)
	return data, decodeErr
}
