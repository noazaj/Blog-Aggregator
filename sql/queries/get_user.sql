-- name: GetUserApiKey :one
SELECT * FROM users
WHERE api_key = $1;