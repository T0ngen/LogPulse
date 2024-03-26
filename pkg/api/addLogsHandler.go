package api

import (
	"fmt"
	"main/pkg/api/mappers"
	"main/pkg/api/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)



func genLog(err error, funcName, file string ) logrus.Fields {
  
	return logrus.Fields{
	  "Error":            err,
	  "File":             file,
	  "FromFunction":     funcName,
	}
  
  }


func (h *handler) AcceptLogs(c *gin.Context){
	key := c.Param("key")

	var requestBody mappers.NewLogForm


	//TODO: working with binding json to the ResetPasswordForm model
	{
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			logrus.WithFields(genLog(err, "ShouldBindJSON",
			"addLogsHandler")).Errorf("Unable to bind requested json to the model")
			c.JSON(http.StatusInternalServerError, gin.H{
					"error": responses.ErrorResponse{
						Error: "Unable to bind requested json to the model ",
						Description: responses.ErrorDescription{
							ErrorCode:           http.StatusBadRequest,
							TagError:            "inner_error",
							DetailedDescription: "Unable to bind requested json to the model, try again later",
						},
					},
				})
				return
		}
	}


	//TODO: check key existence and validating json input
	isKeyExist, err :=h.DB.CheckKeyExist(c, key)
	{
		if err != nil{
			logrus.WithFields(genLog(err, "AcceptLogs",
			"addLogsHandler")).Errorf("Error while checking key existence")
			c.JSON(http.StatusInternalServerError, gin.H{
					"error": responses.ErrorResponse{
						Error: "Error while checking key existence",
						Description: responses.ErrorDescription{
							ErrorCode:           http.StatusBadRequest,
							TagError:            "query_error",
							DetailedDescription: "Error while checking key existence, try again later",
						},
					},
				})
				return
		}

		//TODO: validating json inputs according to the ResetPasswordForm liters
		{
			err := h.Validator.Struct(requestBody)
			fieldErrors := make(map[string]string)
			if err != nil {
				if validationErrors, ok := err.(validator.ValidationErrors); ok {

					for _, fieldError := range validationErrors {

						fieldErrors[fieldError.Field()] = fieldError.ActualTag()
					}
				}
				logrus.WithFields(genLog(err, "AcceptLogs",
				"addLogsHandler")).Errorf("Error validating requested input in json")
				c.JSON(http.StatusBadRequest, gin.H{
					"error":            "Error validating json input. Check the requested json.",
					"validation_error": fieldErrors,
					})
					return
				}
		}

		if !isKeyExist{
			
			fmt.Println("pidor")
		}

	}

	newLog := mappers.GenerateNewLog(key, requestBody)
	
	_, err =h.MongoDB.AddLog(c, newLog)
	if err!= nil{
		fmt.Println("bad")
	}
	fmt.Println("good")
	
	
}