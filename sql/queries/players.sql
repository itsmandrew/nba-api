-- name: GetLeBronJames :one
SELECT * FROM players
WHERE name LIKE '%LeBron%' LIMIT 1;

-- name: GetPlayers :many
SELECT * FROM players
LIMIT 10;

-- name: GetPlayerByID :one
SELECT * FROM players
WHERE id = $1;


-- name: GetPlayerByName :many
SELECT * FROM players
WHERE name ILIKE $1;

-- name: GetPlayersFiltered :many
SELECT
    id,
    name,
    position,
    college,
    year_start,
    height,
    weight,
    birth_date
FROM players
WHERE ($1 = '' OR position ILIKE '%' || $1 || '%')
  AND ($2 = '' OR college ILIKE '%' || $2 || '%')
  AND ($3 = 0 OR year_start = $3)
LIMIT 10;
