package atopdb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MQTTTopicsSettings = "mqttTopicsSettings"
var MQTTTopicsLogs = "mqttTopicsLogs"

// TopicExtraInformation mqtt topic extra informations
type TopicExtraInformation struct {
	Device string `bson:"device" json:"device"` //device name (Atop devices)
}

// TopicSetting
type TopicSetting struct {
	EnableLog bool `bson:"enableLog" json:"enableLog"`
}

// TopicInformation subscribe topics settings & informations
type TopicInformation struct {
	Topic       string    `bson:"topic" json:"topic"`
	LastPayload string    `bson:"lastPayload" json:"lastPayload"`
	Qos         byte      `bson:"qos" json:"qos"`
	LastTS      time.Time `bson:"lastTS" json:"lastTS"`
}

// KeepLastMqttMessage update mqtt message by topic
func (d *DB) KeepLastMqttMessage(topic string, data string, qos byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	collection := d.db.Collection(MQTTTopicsSettings)
	ops := options.FindOneAndUpdate()
	ops.SetUpsert(true)

	rawbody := TopicInformation{
		Topic: topic, LastPayload: data, Qos: qos, LastTS: time.Now().Local(),
	}

	w2mongo := bson.M{
		"$set": rawbody,
	}

	fmt.Printf("%q\n", w2mongo)

	ret := collection.FindOneAndUpdate(ctx, bson.M{"topic": topic}, w2mongo, ops)
	if ret.Err() != nil {
		return ret.Err()
	}
	return d.LogMqttMessage(topic, data, qos)

}

// LogMqttMessage append mqtt message into database "mqttTopicsLogs" collection
func (d *DB) LogMqttMessage(topic string, data string, qos byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	collection := d.db.Collection(MQTTTopicsLogs)
	// ops := options.FindOneAndUpdate()
	// ops.SetUpsert(true)
	body := bson.M{
		"topic":   topic,
		"payload": data,
		"lastTS":  time.Now(),
	}

	_, err := collection.InsertOne(ctx, body)
	return err
}

// GetExcludeLog Get exclude log topic list
// func (d *DB) GetExcludeLog() []string {
// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()
// 	collection := d.db.Collection(MQTTTopicsSettings)

// 	cursor, err := collection.Find(ctx, bson.M{})
// 	if err != nil {
// 		log.Println("READ find error", err)
// 		return nil, err
// 	}
// 	var data []TopicSetting
// 	err = cursor.All(context.TODO(), &data)
// 	if err != nil {
// 		log.Println("READ decode error", err)
// 		return nil, err
// 	}
// 	return data, nil
// 	return err
// }

// GetTopicSettings get topic
// func (d *DB) GetTopicSettings(topic) bson.A {
// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// defer cancel()
// collection := d.db.Collection(MQTTTopicsSettings)

// body := bson.M{
// 	"$set": bson.M{
// 		"topic":       topic,
// 		"lastPayload": data,
// 		"qos":         qos,
// 		"lastTS":      time.Now(),
// 	}}
// 	collection.FindOne()
// ret := collection.FindOneAndUpdate(ctx, bson.M{"topic": topic}, body, ops)
// return ret.Err()
// }
