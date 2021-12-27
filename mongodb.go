package atopdb

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB struct of monge database
type DB struct {
	client *mongo.Client
	db     *mongo.Database
}

var instance *DB
var once sync.Once

var defaultURL string = "mongodb://localhost:27017"
var clientoption *options.ClientOptions = nil

// GetDB get unique database pointer
func GetDB() *DB {
	once.Do(func() {
		instance = initDB(defaultURL)
	})
	return instance
}
func SettingOptions(option *options.ClientOptions) {
	clientoption = option
}

// Initial database
func initDB(url string) *DB {
	const defaultDatabase = "atop"
	db := new(DB)
	var err error
	if clientoption == nil {
		db.client, err = mongo.NewClient(options.Client().ApplyURI(url))
	} else {
		db.client, err = mongo.NewClient(clientoption)
	}
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = db.client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = db.client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	db.db = db.client.Database(defaultDatabase)
	log.Printf("Database %s connected... \n", url)
	return db
}

func (d *DB) GetCollections() []string {
	result, _ := d.db.ListCollectionNames(context.TODO(), bson.M{})

	return result
}

// RunCommand mongo runCommand()
func (d *DB) RunCommand(arg bson.M) (bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var ret bson.M
	err := d.db.RunCommand(ctx, arg).Decode(&ret)
	if err != nil {
		return bson.M{}, err
	}
	return ret, nil
}

// GetStats db.runCommand({dbStates:1})
func (d *DB) GetStats() (bson.M, error) {
	return d.RunCommand(bson.M{"dbStats": 1})
}

func (d *DB) Test() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	collection := d.db.Collection("test")

	if _, err := collection.InsertOne(ctx, bson.M{"name": "hello"}); err != nil {
		log.Println("test insert fail")
	}

}
