-- name: CreateDeliveryPoint :one
INSERT INTO delivery_points (
    id,
    name
) VALUES (
    $1,
    $2
) RETURNING *;

-- name: GetDeliveryPoint :one
SELECT * FROM delivery_points
WHERE id = $1  LIMIT 1;


-- name: DeleteDeliveryPoint :exec
DELETE FROM delivery_points WHERE id = $1;
