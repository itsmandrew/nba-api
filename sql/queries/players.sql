-- name: GetLeBronJames :one
SELECT * FROM players
WHERE name LIKE '%LeBron%' LIMIT 1;

-- name: GetPlayers :many
SELECT * FROM players
LIMIT 10;
