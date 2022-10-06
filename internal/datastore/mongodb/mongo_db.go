package mongodb

import (
	"context"
	"fmt"
	repository2 "github.com/senpathi/kafkajet/internal/datastore/repository"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(conf repository2.DBConfig) (repository2.DB, error) {
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s/?maxPoolSize=%d", conf.User, conf.Password, conf.Address, conf.MaxConnection)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MongoDB")

	return &mongoDB{
		client:   client,
		database: conf.Database,
	}, nil
}

type mongoDB struct {
	client   *mongo.Client
	database string
}

func (m *mongoDB) Table(name string) repository2.Repo {
	return &table{
		collection: m.client.Database(m.database).Collection(name),
	}
}

func (m *mongoDB) Close() {
	err := m.client.Disconnect(context.TODO())
	if err != nil {
		log.Println(err)
	}
}

type table struct {
	collection *mongo.Collection
}

func (m *table) Read(filter map[string]interface{}) (value interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (m *table) Write(value interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m *table) Delete(filter map[string]interface{}) {
	//TODO implement me
	panic("implement me")
}
