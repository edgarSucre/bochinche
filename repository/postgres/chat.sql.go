// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: chat.sql

package postgres

import (
	"context"
)

const createChat = `-- name: CreateChat :exec
INSERT INTO chats (
    room,
    author,
    message
) VALUES (
    $1, $2, $3
)
`

type CreateChatParams struct {
	Room    string
	Author  string
	Message string
}

func (q *Queries) CreateChat(ctx context.Context, arg CreateChatParams) error {
	_, err := q.db.ExecContext(ctx, createChat, arg.Room, arg.Author, arg.Message)
	return err
}

const listChats = `-- name: ListChats :many
SELECT id, room, author, message, created_at FROM chats
WHERE room = $1
ORDER BY created_at
LIMIT 50
`

func (q *Queries) ListChats(ctx context.Context, room string) ([]Chat, error) {
	rows, err := q.db.QueryContext(ctx, listChats, room)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Chat{}
	for rows.Next() {
		var i Chat
		if err := rows.Scan(
			&i.ID,
			&i.Room,
			&i.Author,
			&i.Message,
			&i.CreatedAt,
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