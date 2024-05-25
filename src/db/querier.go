// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	//DeleteChatroom
	//
	//  DELETE FROM chatrooms WHERE id = ?
	DeleteChatroom(ctx context.Context, id int64) error
	//DeleteMessage
	//
	//  DELETE FROM messages WHERE id = ?
	DeleteMessage(ctx context.Context, id int64) error
	//DeleteTodo
	//
	//  DELETE FROM todos WHERE id = ?
	DeleteTodo(ctx context.Context, id int64) error
	//DeleteTokensByUserId
	//
	//  DELETE FROM tokens WHERE id IN (
	//          SELECT token_id FROM users_tokens WHERE user_id = ?
	//      )
	DeleteTokensByUserId(ctx context.Context, userID string) error
	//DeleteUser
	//
	//  DELETE FROM users WHERE id = ?
	DeleteUser(ctx context.Context, id string) error
	//InsertChatroom
	//
	//  INSERT INTO chatrooms (name, creator, created_at) VALUES (?, ?, ?) RETURNING id
	InsertChatroom(ctx context.Context, arg InsertChatroomParams) (int64, error)
	//InsertMessage
	//
	//  INSERT INTO messages (author, message, color, room_id, created_at) VALUES (?, ?, ?, ?, ?)
	InsertMessage(ctx context.Context, arg InsertMessageParams) error
	//InsertTodo
	//
	//  INSERT INTO todos (body, username, created_at) VALUES (?, ?, ?) RETURNING id, body, username, created_at
	InsertTodo(ctx context.Context, arg InsertTodoParams) (Todo, error)
	//InsertToken
	//
	//  INSERT INTO tokens (jwt_type, jwt, valid, family) VALUES (?, ?, ?, ?) RETURNING id, jwt_type, jwt, valid, family
	InsertToken(ctx context.Context, arg InsertTokenParams) (Token, error)
	// table: users
	//
	//  INSERT INTO users (id, email, username, encrypted_password, password_created_at)
	//  VALUES (?, ?, ?, ?, ?) RETURNING id, email, username
	InsertUser(ctx context.Context, arg InsertUserParams) (InsertUserRow, error)
	//InsertUsersTokens
	//
	//  INSERT INTO users_tokens (user_id, token_id) VALUES (?, ?)
	InsertUsersTokens(ctx context.Context, arg InsertUsersTokensParams) error
	// table: chatrooms
	//
	//  SELECT id, name, creator, created_at FROM chatrooms ORDER BY created_at DESC LIMIT ?
	SelectChatrooms(ctx context.Context, limit int64) ([]Chatroom, error)
	//SelectEmailOrUsernameAlreadyExists
	//
	//  SELECT email FROM users WHERE email = ? OR username = ?
	SelectEmailOrUsernameAlreadyExists(ctx context.Context, arg SelectEmailOrUsernameAlreadyExistsParams) (string, error)
	// table: messages
	//
	//  SELECT id, author, message, color, room_id, created_at FROM messages WHERE room_id = ? ORDER BY created_at DESC LIMIT ?
	SelectMessagesByChatroom(ctx context.Context, arg SelectMessagesByChatroomParams) ([]Message, error)
	//SelectMessagesByUser
	//
	//  SELECT id, author, message, color, room_id, created_at FROM messages WHERE author = ? ORDER BY created_at DESC LIMIT ?
	SelectMessagesByUser(ctx context.Context, arg SelectMessagesByUserParams) ([]Message, error)
	// table: todos
	//
	//  SELECT id, body, username, created_at FROM todos WHERE username = ?
	SelectTodosByUsername(ctx context.Context, username string) ([]Todo, error)
	// table: tokens
	//
	//  SELECT id, jwt_type, jwt, valid, family FROM tokens WHERE id = ?
	SelectTokenFromId(ctx context.Context, id int64) (Token, error)
	//SelectTokenFromJwtString
	//
	//  SELECT id, jwt_type, jwt, valid, family FROM tokens WHERE jwt = ?
	SelectTokenFromJwtString(ctx context.Context, jwt string) (Token, error)
	//SelectUserByEmail
	//
	//  SELECT id, email, username, encrypted_password, password_created_at FROM users WHERE email = ?
	SelectUserByEmail(ctx context.Context, email string) (User, error)
	//SelectUserByUsername
	//
	//  SELECT id, email, username, encrypted_password, password_created_at FROM users WHERE username = ?
	SelectUserByUsername(ctx context.Context, username string) (User, error)
	//SelectUserIdFromToken
	//
	//  SELECT user_id FROM users_tokens WHERE token_id = ? LIMIT 1
	SelectUserIdFromToken(ctx context.Context, tokenID int64) (string, error)
	// table: users_tokens
	//
	//  SELECT user_id, token_id FROM users_tokens WHERE user_id = ?
	SelectUsersTokens(ctx context.Context, userID string) ([]UsersToken, error)
	//UpdateChatroomName
	//
	//  UPDATE chatrooms SET name = ? WHERE id = ? RETURNING id, name, creator, created_at
	UpdateChatroomName(ctx context.Context, arg UpdateChatroomNameParams) (Chatroom, error)
	//UpdateMessage
	//
	//  UPDATE messages SET message = ? WHERE id = ? RETURNING id, author, message, color, room_id, created_at
	UpdateMessage(ctx context.Context, arg UpdateMessageParams) (Message, error)
	//UpdateTodo
	//
	//  UPDATE todos SET body = ? WHERE id = ? RETURNING id, body, username, created_at
	UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error)
	//UpdateTokenValid
	//
	//  UPDATE tokens SET valid = ? WHERE jwt = ? RETURNING id
	UpdateTokenValid(ctx context.Context, arg UpdateTokenValidParams) (int64, error)
	//UpdateTokensFamilyInvalid
	//
	//  UPDATE tokens SET valid = FALSE WHERE family = ?
	UpdateTokensFamilyInvalid(ctx context.Context, family string) error
	//UpdateUserTokensToInvalid
	//
	//  UPDATE tokens SET valid = FALSE WHERE id IN (
	//          SELECT token_id FROM users_tokens WHERE user_id = ?
	//      )
	UpdateUserTokensToInvalid(ctx context.Context, userID string) error
}

var _ Querier = (*Queries)(nil)
