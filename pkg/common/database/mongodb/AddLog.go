package mongodb

import (
	"context"
	"main/pkg/common/models"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)
func (m *MyMongoClient) AddLog(ctx context.Context, logData models.LogsModel) (*mongo.InsertOneResult, error) {
    database := m.Database("main")

    logsCollection := database.Collection("standart_logs")

    insertResult, err := logsCollection.InsertOne(ctx, logData)
    if err != nil {
        logrus.WithFields(logrus.Fields{
            "Error":        err.Error(),
            "FromFunction": "AddLog",
            "File":         "logManager.go", 
        }).Warningf("Can't insert ")

        return nil, err
    }
    return insertResult, nil
}
