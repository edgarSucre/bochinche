-- name: CreateChat :exec
INSERT INTO chats (
    room,
    author,
    message
) VALUES (
    $1, $2, $3
);

-- name: ListChats :many
SELECT * FROM chats
WHERE room = $1
ORDER BY created_at
LIMIT 50;