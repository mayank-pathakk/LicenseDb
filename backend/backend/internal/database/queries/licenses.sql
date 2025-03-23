-- name: GetLicenseByID :one
SELECT id, name FROM licenses WHERE id = $1;
