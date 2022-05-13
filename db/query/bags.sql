-- name: CreateBag :one
INSERT INTO bags (
    barcode,
    bag_status,
    delivery_id
) VALUES (
    $1,$2,$3
) RETURNING *;

-- name: GetBag :one
SELECT * FROM bags
WHERE barcode = $1  LIMIT 1;

-- name: ListBags :many
SELECT * FROM bags
ORDER BY barcode;

-- name: UpdateBag :one
UPDATE bags SET bag_status = $2
WHERE barcode = $1
RETURNING *;

-- name: DeleteBag :exec
DELETE FROM bags WHERE barcode = $1;
