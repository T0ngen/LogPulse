package api

import (
	"fmt"
	"main/pkg/api/mappers"

	"github.com/gin-gonic/gin"
)





func (h *handler) AcceptLogs(c *gin.Context){
	key := c.Param("key")

	var requestBody mappers.NewLogForm

	isKeyExist, err :=h.DB.CheckKeyExist(c, key)
	if err != nil{
		fmt.Println(err)
	}
	if isKeyExist{
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			fmt.Println(err)
		}
		newLog := mappers.GenerateNewLog(requestBody)

		_, err :=h.MongoDB.AddLog(c, newLog)
		if err!= nil{
			fmt.Println("bad")
		}
		fmt.Println("good")

	}
	
}