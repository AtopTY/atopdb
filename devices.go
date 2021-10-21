package atopdb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DEVICES_INFO = "devicesInfo"

type DataNode struct {
	Type   string `json:"type" bson:"type"`
	Name   string `json:"name" bson:"name"`
	Access string `json:"access" bson:"access"`
}

type DeviceInfo struct {
	Inet  string     `json:"inet" bson:"inet"`
	Mac   string     `json:"mac" bson:"mac"`
	Nodes []DataNode `json:"nodes" bson:"nodes"`
}

func (d *DB) WriteDeviceInfo(info *DeviceInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	collection := d.db.Collection(DEVICES_INFO)
	ops := options.FindOneAndUpdate()
	ops.SetUpsert(true)
	data, err := bson.Marshal(info)
	if err != nil {
		return err
	}
	ret := collection.FindOneAndUpdate(ctx, bson.M{"mac": info.Mac}, data, ops)
	return ret.Err()
}
