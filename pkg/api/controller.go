package api

import (
	"context"
	"main/pkg/common/database/mongodb"
	"main/pkg/common/database/sqlc"
	"main/pkg/common/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/mongo"
)

type KeyExist interface{
	CheckKeyExist(ctx context.Context, key string) (bool, error)
}

type LogsMongo interface{
	AddLog(ctx context.Context, logData models.LogsModel) (*mongo.InsertOneResult, error)
	GetLogs(ctx context.Context, key string) ([]models.LogsModel, error)
	GetLogsById(ctx context.Context, key string, reqId string) ([]models.LogsModel, error)
}

type handler struct{
	DB KeyExist
	MongoDB LogsMongo
	Validator *validator.Validate
	

}


func RegisterRouter(r *gin.Engine, validate *validator.Validate, db *sqlc.Queries, client *mongodb.MyMongoClient ){
	
	h := &handler{DB: db, MongoDB: client, Validator: validate}
	
	routes := r.Group("/api/v1/")
	
	routes.POST("/sendlog/:key", h.AcceptLogs)
	routes.GET("/logs/:key", gin.BasicAuth(gin.Accounts{
       "tgbot" : "tgbot", 
    }), h.GetLogs)
	routes.GET("/logsbyid/:key/:id", gin.BasicAuth(gin.Accounts{
       "tgbot" : "tgbot", 
    }), h.GetLogsByReqId)
}