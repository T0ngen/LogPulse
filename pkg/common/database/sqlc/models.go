// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"time"
)

type Project struct {
	ID          int64     `json:"id"`
	TgID        int64     `json:"tg_id"`
	ProjectName string    `json:"project_name"`
	Key         string    `json:"key"`
	CreatedAt   time.Time `json:"created_at"`
}

type User struct {
	ID        int32     `json:"id"`
	TgID      int64     `json:"tg_id"`
	Username  string    `json:"username"`
	Balance   int64     `json:"balance"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
