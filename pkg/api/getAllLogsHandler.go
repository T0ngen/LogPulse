package api

import (
	"main/pkg/api/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)



func (h *handler) GetLogs(c *gin.Context) {
	key := c.Param("key")


	//TODO: check key existence and validating json input
	isKeyExist, err :=h.DB.CheckKeyExist(c, key)
	{
		if err != nil{
			logrus.WithFields(genLog(err, "GetLogs",
			"getAllLogsHandler")).Errorf("Error while checking key existence")
			c.JSON(http.StatusInternalServerError, gin.H{
					"error": responses.ErrorResponse{
						Error: "Error while checking key existence",
						Description: responses.ErrorDescription{
							ErrorCode:           http.StatusInternalServerError,
							TagError:            "inner_error",
							DetailedDescription: "Error while checking key existence, try again later",
						},
					},
				})
				return
		}



		if !isKeyExist{
			logrus.WithFields(genLog(err, "GetLogs",
			"getAllLogsHandler")).Infof("Error! The key is not registered in the system")
			c.JSON(http.StatusNotFound, gin.H{
					"error": responses.ErrorResponse{
						Error: "Error! The key is not registered in the system",
						Description: responses.ErrorDescription{
							ErrorCode:           http.StatusNotFound,
							TagError:            "error",
							DetailedDescription: "The key is not registered in the system, try again with other key",
						},
					},
				})
				return
			
		}


	}

	logs, err := h.MongoDB.GetLogs(c,key)
	if err != nil {
		logrus.WithFields(genLog(err, "GetLogs",
			"getAllLogsHandler")).Errorf("Error while getting logs from DB")
			c.JSON(http.StatusInternalServerError, gin.H{
					"error": responses.ErrorResponse{
						Error: "Error while getting logs from DB",
						Description: responses.ErrorDescription{
							ErrorCode:           http.StatusInternalServerError,
							TagError:            "inner_error",
							DetailedDescription: "Error while getting logs from DB, try again later",
						},
					},
				})
				return
	}
	
	c.JSON(200, gin.H{"logs": logs})
	



	
}