package db

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"samples-golang/initializer"
)

type MongoDB struct {
	Client *mongo.Client
	DbName string
}

func (m *MongoDB) Connect() {
	config, err := initializer.LoadConfig(".")
	clientOptions := options.Client().ApplyURI(config.UriAddress)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Error(err.Error())
		return
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Error(err.Error())
		return
	}
	fmt.Println("MongoDB connection successful")
	m.Client = client
}

func (m *MongoDB) Close() {
	err := m.Client.Disconnect(context.Background())
	if err != nil {
		log.Error(err.Error())
		return
	}
	fmt.Println("MongoDB connection closed")
}
