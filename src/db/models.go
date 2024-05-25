// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"
)

type Chatroom struct {
	ID        int64
	Name      string
	Creator   string
	CreatedAt *time.Time
}

type Message struct {
	ID        int64
	Author    string
	Message   string
	Color     string
	RoomID    int64
	CreatedAt *time.Time
}

type Todo struct {
	ID        int64
	Body      string
	Username  string
	CreatedAt *time.Time
}

type Token struct {
	ID      int64
	JwtType string
	Jwt     string
	Valid   bool
	Family  string
}

type User struct {
	ID                string
	Email             string
	Username          string
	EncryptedPassword string
	PasswordCreatedAt time.Time
}

type UsersToken struct {
	UserID  string
	TokenID int64
}
