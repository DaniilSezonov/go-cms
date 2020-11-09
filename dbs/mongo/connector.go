package dbs

import (
	"context"
	"github.com/nerlin/go-cms/dbs"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once

type MongoConnector struct {
	instance dbs.DBConnector
	client   *mongo.Client
}

func (mc *MongoConnector) GetInstance() dbs.DBConnector {
	once.Do(func() {
		var instance = new(MongoConnector)
		mc.instance = instance
	})
	return mc.instance
}
func (mc *MongoConnector) Connect(uri string) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return
	}
	mc.client = client
	err = mc.client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return
	}
	err = mc.client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB has been established")
}
func (mc *MongoConnector) Disconnect() {
	if mc.client == nil {
		log.Println("Already disconnected")
		return
	}
	mc.client.Disconnect(context.TODO())
	mc.client = nil
}
