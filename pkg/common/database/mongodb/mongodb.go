package mongodb

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type MyMongoClient struct {
    *mongo.Client
}


func Init(url string) (*MyMongoClient, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)

    client, err := mongo.Connect(ctx, opts)
    if err != nil {
        logrus.Fatal("Can't create DB client: ", err)
        return nil, err 
    }

    
    return &MyMongoClient{client}, nil
}
