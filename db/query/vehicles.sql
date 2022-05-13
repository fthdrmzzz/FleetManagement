-- name: CreateVehicle :one
INSERT INTO vehicles (
    plate
) VALUES (
    $1
) RETURNING *;

-- name: GetVehicle :one
SELECT * FROM vehicles
WHERE plate = $1  LIMIT 1;

-- name: ListVehicles :many
SELECT * FROM vehicles
ORDER BY plate;

-- name: DeleteVehicle :exec
DELETE FROM vehicles WHERE plate = $1;


