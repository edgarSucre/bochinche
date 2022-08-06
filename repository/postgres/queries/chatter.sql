-- name: RegisterChatter :exec
INSERT INTO chatters (
    username,
    password,
    email
) VALUES (
    $1, $2, $3
);

-- name: VerifyChatter :one
SELECT * FROM chatters
WHERE username = $1;