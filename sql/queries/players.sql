-- name: GetLeBronJames :one
SELECT * FROM players
WHERE name LIKE '%LeBron%' LIMIT 1;
