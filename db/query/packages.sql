-- name: CreatePackage :one
INSERT INTO packages (
    barcode,
    package_status,
    package_weight,
    delivery_id
) VALUES (
    $1,$2,$3,$4
) RETURNING *;

-- name: GetPackage :one
SELECT * FROM packages
WHERE barcode = $1  LIMIT 1;

-- name: ListPackages :many
SELECT * FROM packages
ORDER BY barcode;

-- name: UpdatePackage :one
UPDATE packages SET package_status = $2
WHERE barcode = $1
RETURNING *;

-- name: DeletePackage :exec
DELETE FROM packages WHERE barcode = $1;


