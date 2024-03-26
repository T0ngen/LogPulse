package api

import (
	"main/pkg/common/database/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)


type handler struct{
	DB *sqlc.Queries
	
	Validator *validator.Validate
	

}


func RegisterRouter(r *gin.Engine, validate *validator.Validate, db *sqlc.Queries ){
	

	h := &handler{DB: db, Validator: validate}

	routes := r.Group("/api/v1/")
	
	routes.POST("/sendlog/:key", h.AcceptLogs)

}