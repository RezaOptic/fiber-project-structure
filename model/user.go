package model

import "time"

type SubmitUserRequest struct {
	Username string `json:"username"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
}
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseUsers struct {
	Users       []User `json:"users"`
	HasNextPage bool   `json:"has_next_page"`
}
