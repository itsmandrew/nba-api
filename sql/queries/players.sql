-- name: GetLeBronJames :one
SELECT * FROM players
WHERE name LIKE '%LeBron%' LIMIT 1;

-- name: GetPlayers :many
SELECT * FROM players
LIMIT 10;

-- name: GetPlayerByID :one
SELECT * FROM PLAYERS
WHERE id = $1;


-- name: GetPlayerByName :many
SELECT * FROM PLAYERS
WHERE name = $1;

