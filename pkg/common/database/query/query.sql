-- name: CheckKeyExist :one
SELECT EXISTS (
    SELECT 1 FROM "projects" WHERE "key" = $1
) AS "exists";