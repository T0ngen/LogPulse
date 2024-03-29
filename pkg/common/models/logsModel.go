package models

import "time"


type LogsModel struct {
    Key string `bson:"key"`
	Timestamp        time.Time `bson:"timestamp"`
	Level       string    `bson:"level"`       
    Message     string    `bson:"message"`     
    UserID      string    `bson:"user_id"`
    RequestID   string    `bson:"request_id"` 
    IP          string    `bson:"ip"`       
    UserAgent   string    `bson:"user_agent"`
    ErrorCode   string    `bson:"error_code"`
    StackTrace  string    `bson:"stack_trace"`
}
type GetLogsModel struct {
	Timestamp        time.Time `bson:"timestamp"`
	Level       string    `bson:"level"`       
    Message     string    `bson:"message"`     
    UserID      string    `bson:"user_id"`
    RequestID   string    `bson:"request_id"` 
    IP          string    `bson:"ip"`       
    UserAgent   string    `bson:"user_agent"`
    ErrorCode   string    `bson:"error_code"`
    StackTrace  string    `bson:"stack_trace"`
}
