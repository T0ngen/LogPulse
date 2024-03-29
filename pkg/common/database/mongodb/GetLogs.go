package mongodb

import (
	"context"
	"main/pkg/common/models"

	"go.mongodb.org/mongo-driver/bson"
)


func (m *MyMongoClient) GetLogs(ctx context.Context, key string) ([]models.LogsModel, error){
	var logs []models.LogsModel

	database := m.Database("main")
	
	logsCollection := database.Collection("standart_logs")

	filter := bson.D{{Key: "key", Value: key}}
	cursor, err := logsCollection.Find(ctx, filter)
	if err != nil {
		return logs, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var log models.LogsModel
		err := cursor.Decode(&log)
		if err != nil {
			return logs, err
		}
		logs = append(logs, log)
	}
	return logs, nil


}
