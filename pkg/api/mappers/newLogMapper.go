package mappers

import (
	"main/pkg/common/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewLogForm struct {
	Key string `json:"key"`
    Timestamp   time.Time `json:"timestamp"`   // Временная метка события
    Level       string    `json:"level"`       // Уровень лога, например, INFO, ERROR, DEBUG
    Message     string    `json:"message"`     // Основное сообщение лога
    UserID      string    `json:"userId,omitempty"` // ID пользователя, если применимо
    RequestID   string    `json:"requestId,omitempty"` // ID запроса, для связывания логов одного запроса
    IP          string    `json:"ip,omitempty"`       // IP-адрес запроса, если применимо
    UserAgent   string    `json:"userAgent,omitempty"`// User-Agent запроса, если применимо
    ErrorCode   string    `json:"errorCode,omitempty"`// Код ошибки, если применимо
    StackTrace  string    `json:"stackTrace,omitempty"` // Стек вызовов для ошибок
    CustomData  map[string]interface{} `json:"customData,omitempty"` // Дополнительные пользовательские данные
}
func GenerateNewLog(newLogForm NewLogForm) models.LogsModel {
    
    accountID := primitive.NewObjectID().Hex()

    
    currentTime := time.Now()

   
    newLog := models.LogsModel{
        
        Timestamp:  currentTime,
        Level:      "INFO", 
        Message:    "New user account created",
        UserID:     accountID,
        RequestID:  "",        
        IP:         "",        
        UserAgent:  "",       
        ErrorCode:  "",       
        StackTrace: "",       
        CustomData: map[string]interface{}{
           
        },
    }

    return newLog
}
