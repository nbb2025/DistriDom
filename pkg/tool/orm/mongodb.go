package orm

import (
	"context"
	"github.com/nbb2025/distri-domain/app/static/config"
	"github.com/nbb2025/distri-domain/pkg/util/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var ContextTodo = context.TODO()
var client *mongo.Client
var database *mongo.Database

func MongoDB() *mongo.Database {
	return database
}

func MongoInit() {
	if client != nil && database != nil {
		panic("MongoDB already initialized")
	}

	_config := config.Conf.MongoConfig
	clientOptions := options.Client().ApplyURI(_config.URI)

	// 连接到MongoDB
	newClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.Error("MongoDB connect error: " + err.Error())
		panic(err)
	}

	// 检查连接
	err = newClient.Ping(context.TODO(), nil)
	if err != nil {
		logger.Error("MongoDB ping error: " + err.Error())
		panic(err)
	}

	client = newClient
	database = client.Database(_config.DBName)
}

func MongoUnInit() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	log.Println("MongoDB disconnected")
}
