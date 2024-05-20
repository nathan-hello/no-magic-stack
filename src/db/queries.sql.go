// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: queries.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const deleteChatroom = `-- name: DeleteChatroom :exec
DELETE FROM chatrooms WHERE id = $1
`

func (q *Queries) DeleteChatroom(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteChatroom, id)
	return err
}

const deleteMessage = `-- name: DeleteMessage :exec
DELETE FROM messages WHERE id = $1
`

func (q *Queries) DeleteMessage(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMessage, id)
	return err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1
`

func (q *Queries) DeleteTodo(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const deleteTokensByUserId = `-- name: DeleteTokensByUserId :exec
DELETE FROM tokens
WHERE tokens.id IN (
        SELECT token_id FROM users_tokens WHERE users_tokens.user_id = $1
    )
`

func (q *Queries) DeleteTokensByUserId(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTokensByUserId, userID)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const insertChatroom = `-- name: InsertChatroom :one
INSERT INTO chatrooms (name, creator) VALUES ($1, $2) RETURNING id
`

type InsertChatroomParams struct {
	Name    string
	Creator string
}

func (q *Queries) InsertChatroom(ctx context.Context, arg InsertChatroomParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, insertChatroom, arg.Name, arg.Creator)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const insertMessage = `-- name: InsertMessage :exec
INSERT INTO messages (author, message, room_id, created_at) VALUES ($1, $2, $3, $4)
`

type InsertMessageParams struct {
	Author    string
	Message   string
	RoomID    int64
	CreatedAt string
}

func (q *Queries) InsertMessage(ctx context.Context, arg InsertMessageParams) error {
	_, err := q.db.ExecContext(ctx, insertMessage,
		arg.Author,
		arg.Message,
		arg.RoomID,
		arg.CreatedAt,
	)
	return err
}

const insertTodo = `-- name: InsertTodo :one
INSERT INTO todos (body, username) VALUES ($1, $2) RETURNING id, created_at, body, username
`

type InsertTodoParams struct {
	Body     string
	Username string
}

func (q *Queries) InsertTodo(ctx context.Context, arg InsertTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, insertTodo, arg.Body, arg.Username)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Body,
		&i.Username,
	)
	return i, err
}

const insertToken = `-- name: InsertToken :one
INSERT INTO tokens (jwt_type, jwt, valid, family) VALUES ($1, $2, $3, $4) RETURNING id
`

type InsertTokenParams struct {
	JwtType string
	Jwt     string
	Valid   bool
	Family  uuid.UUID
}

func (q *Queries) InsertToken(ctx context.Context, arg InsertTokenParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, insertToken,
		arg.JwtType,
		arg.Jwt,
		arg.Valid,
		arg.Family,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO users ( email, username, encrypted_password, password_created_at)
values ($1, $2, $3, $4)
RETURNING id, email, username
`

type InsertUserParams struct {
	Email             sql.NullString
	Username          string
	EncryptedPassword string
	PasswordCreatedAt time.Time
}

type InsertUserRow struct {
	ID       uuid.UUID
	Email    sql.NullString
	Username string
}

// table: users
func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (InsertUserRow, error) {
	row := q.db.QueryRowContext(ctx, insertUser,
		arg.Email,
		arg.Username,
		arg.EncryptedPassword,
		arg.PasswordCreatedAt,
	)
	var i InsertUserRow
	err := row.Scan(&i.ID, &i.Email, &i.Username)
	return i, err
}

const insertUsersTokens = `-- name: InsertUsersTokens :exec
INSERT INTO users_tokens (user_id, token_id) VALUES ($1, $2)
`

type InsertUsersTokensParams struct {
	UserID  uuid.UUID
	TokenID int64
}

func (q *Queries) InsertUsersTokens(ctx context.Context, arg InsertUsersTokensParams) error {
	_, err := q.db.ExecContext(ctx, insertUsersTokens, arg.UserID, arg.TokenID)
	return err
}

const selectChatrooms = `-- name: SelectChatrooms :many
SELECT id, created_at, name, creator FROM chatrooms ORDER BY created_at DESC  LIMIT $1
`

// table: chatrooms
func (q *Queries) SelectChatrooms(ctx context.Context, limit int64) ([]Chatroom, error) {
	rows, err := q.db.QueryContext(ctx, selectChatrooms, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chatroom
	for rows.Next() {
		var i Chatroom
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Creator,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectEmailOrUsernameAlreadyExists = `-- name: SelectEmailOrUsernameAlreadyExists :one
SELECT email FROM users WHERE users.email = $1 OR users.username = $2
`

type SelectEmailOrUsernameAlreadyExistsParams struct {
	Email    sql.NullString
	Username string
}

func (q *Queries) SelectEmailOrUsernameAlreadyExists(ctx context.Context, arg SelectEmailOrUsernameAlreadyExistsParams) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, selectEmailOrUsernameAlreadyExists, arg.Email, arg.Username)
	var email sql.NullString
	err := row.Scan(&email)
	return email, err
}

const selectMessagesByChatroom = `-- name: SelectMessagesByChatroom :many
SELECT id, created_at, author, message, room_id FROM messages WHERE room_id = $1 ORDER BY created_at DESC LIMIT $2
`

type SelectMessagesByChatroomParams struct {
	RoomID int64
	Limit  int64
}

// table: messages
func (q *Queries) SelectMessagesByChatroom(ctx context.Context, arg SelectMessagesByChatroomParams) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, selectMessagesByChatroom, arg.RoomID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Author,
			&i.Message,
			&i.RoomID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectMessagesByUser = `-- name: SelectMessagesByUser :many
SELECT id, created_at, author, message, room_id FROM messages WHERE author = $1 ORDER BY created_at DESC LIMIT $2
`

type SelectMessagesByUserParams struct {
	Author string
	Limit  int64
}

func (q *Queries) SelectMessagesByUser(ctx context.Context, arg SelectMessagesByUserParams) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, selectMessagesByUser, arg.Author, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Author,
			&i.Message,
			&i.RoomID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectTodosByUsername = `-- name: SelectTodosByUsername :many
SELECT id, created_at, body, username FROM todos WHERE username = $1
`

// table: todos
func (q *Queries) SelectTodosByUsername(ctx context.Context, username string) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, selectTodosByUsername, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Body,
			&i.Username,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectTokenFromId = `-- name: SelectTokenFromId :one
SELECT id, jwt_type, jwt, valid, family FROM tokens WHERE id = $1
`

// table: tokens
func (q *Queries) SelectTokenFromId(ctx context.Context, id int64) (Token, error) {
	row := q.db.QueryRowContext(ctx, selectTokenFromId, id)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.JwtType,
		&i.Jwt,
		&i.Valid,
		&i.Family,
	)
	return i, err
}

const selectTokenFromJwtString = `-- name: SelectTokenFromJwtString :one
SELECT id, jwt_type, jwt, valid, family FROM tokens WHERE jwt = $1
`

func (q *Queries) SelectTokenFromJwtString(ctx context.Context, jwt string) (Token, error) {
	row := q.db.QueryRowContext(ctx, selectTokenFromJwtString, jwt)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.JwtType,
		&i.Jwt,
		&i.Valid,
		&i.Family,
	)
	return i, err
}

const selectUserByEmail = `-- name: SelectUserByEmail :one
SELECT created_at, username, email, encrypted_password, password_created_at, id FROM users WHERE email = $1
`

func (q *Queries) SelectUserByEmail(ctx context.Context, email sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, selectUserByEmail, email)
	var i User
	err := row.Scan(
		&i.CreatedAt,
		&i.Username,
		&i.Email,
		&i.EncryptedPassword,
		&i.PasswordCreatedAt,
		&i.ID,
	)
	return i, err
}

const selectUserByUsername = `-- name: SelectUserByUsername :one
SELECT created_at, username, email, encrypted_password, password_created_at, id FROM users WHERE username = $1
`

func (q *Queries) SelectUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, selectUserByUsername, username)
	var i User
	err := row.Scan(
		&i.CreatedAt,
		&i.Username,
		&i.Email,
		&i.EncryptedPassword,
		&i.PasswordCreatedAt,
		&i.ID,
	)
	return i, err
}

const selectUserIdFromToken = `-- name: SelectUserIdFromToken :one
SELECT user_id FROM users_tokens WHERE token_id = $1 LIMIT 1
`

func (q *Queries) SelectUserIdFromToken(ctx context.Context, tokenID int64) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, selectUserIdFromToken, tokenID)
	var user_id uuid.UUID
	err := row.Scan(&user_id)
	return user_id, err
}

const selectUsersTokens = `-- name: SelectUsersTokens :many
SELECT id, user_id, token_id FROM users_tokens WHERE user_id = $1
`

// table: users_tokens
func (q *Queries) SelectUsersTokens(ctx context.Context, userID uuid.UUID) ([]UsersToken, error) {
	rows, err := q.db.QueryContext(ctx, selectUsersTokens, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UsersToken
	for rows.Next() {
		var i UsersToken
		if err := rows.Scan(&i.ID, &i.UserID, &i.TokenID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateChatroomName = `-- name: UpdateChatroomName :one
UPDATE chatrooms SET name = $1 WHERE id = $2 RETURNING id, created_at, name, creator
`

type UpdateChatroomNameParams struct {
	Name string
	ID   int64
}

func (q *Queries) UpdateChatroomName(ctx context.Context, arg UpdateChatroomNameParams) (Chatroom, error) {
	row := q.db.QueryRowContext(ctx, updateChatroomName, arg.Name, arg.ID)
	var i Chatroom
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Creator,
	)
	return i, err
}

const updateMessage = `-- name: UpdateMessage :one
UPDATE messages SET message = $1 WHERE id = $2 RETURNING id, created_at, author, message, room_id
`

type UpdateMessageParams struct {
	Message string
	ID      int64
}

func (q *Queries) UpdateMessage(ctx context.Context, arg UpdateMessageParams) (Message, error) {
	row := q.db.QueryRowContext(ctx, updateMessage, arg.Message, arg.ID)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Author,
		&i.Message,
		&i.RoomID,
	)
	return i, err
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE todos SET body = $1 WHERE id = $2 RETURNING id, created_at, body, username
`

type UpdateTodoParams struct {
	Body string
	ID   int64
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateTodo, arg.Body, arg.ID)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Body,
		&i.Username,
	)
	return i, err
}

const updateTokenValid = `-- name: UpdateTokenValid :one
UPDATE tokens SET valid = $1 WHERE jwt = $2 RETURNING id
`

type UpdateTokenValidParams struct {
	Valid bool
	Jwt   string
}

func (q *Queries) UpdateTokenValid(ctx context.Context, arg UpdateTokenValidParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, updateTokenValid, arg.Valid, arg.Jwt)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const updateTokensFamilyInvalid = `-- name: UpdateTokensFamilyInvalid :exec
UPDATE tokens SET valid = FALSE WHERE family = $1
`

func (q *Queries) UpdateTokensFamilyInvalid(ctx context.Context, family uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, updateTokensFamilyInvalid, family)
	return err
}

const updateUserTokensToInvalid = `-- name: UpdateUserTokensToInvalid :exec
UPDATE tokens SET valid = FALSE FROM users_tokens
INNER JOIN tokens AS t ON users_tokens.token_id = t.id
    WHERE users_tokens.user_id = $1
    AND tokens.id = t.id
`

func (q *Queries) UpdateUserTokensToInvalid(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, updateUserTokensToInvalid, userID)
	return err
}
