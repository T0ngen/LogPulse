package mappers

import (
	"main/pkg/common/models"
	"time"

)

type NewLogForm struct {
    Timestamp   time.Time `json:"timestamp"`   
    Level       string    `json:"level" validate:"required"`       
    Message     string    `json:"message" validate:"required"`     
    UserID      string    `json:"user_id,omitempty"` 
    RequestID   string    `json:"request_id,omitempty"` 
    IP          string    `json:"ip,omitempty"`      
    UserAgent   string    `json:"user_agent,omitempty"`
    ErrorCode   string    `json:"error_code,omitempty"`
    StackTrace  string    `json:"stack_trace,omitempty"`
}
func GenerateNewLog(key string,  newLogForm NewLogForm) models.LogsModel {
    
    
    currentTime := time.Now()

   
    newLog := models.LogsModel{
        Key: key,
        Timestamp:  currentTime,
        Level:      newLogForm.Level, 
        Message:    newLogForm.Message,
        UserID:     newLogForm.UserID,
        RequestID:  newLogForm.RequestID,        
        IP:         newLogForm.IP,        
        UserAgent:  newLogForm.UserAgent,       
        ErrorCode:  newLogForm.ErrorCode,       
        StackTrace: newLogForm.StackTrace,       
    }

    return newLog
}
