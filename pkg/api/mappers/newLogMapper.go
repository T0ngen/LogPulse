package mappers

import "time"

type NewLogForm struct {
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