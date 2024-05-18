-- name: GetTasks :many
SELECT *
FROM task
WHERE task.project_id = $1;

-- name: GetAllStudentsByIndex :many
SELECT *
FROM student
WHERE student.index LIKE sqlc.arg(index)::text || '%';

-- name: GetStudentByIndex :one
SELECT *
FROM student
WHERE student.index LIKE $1
LIMIT 1;

-- name: GetProject :one
SELECT *
FROM project
WHERE project.id = $1
LIMIT 1;

-- name: GetProjects :many
SELECT *
FROM project
ORDER BY sqlc.arg(ord)::text
OFFSET sqlc.arg(page)::int * sqlc.arg(size)::int LIMIT sqlc.arg(size)::int;

-- name: GetProjectsByName :many
SELECT *
FROM project
WHERE upper(project.name) LIKE upper('%' || sqlc.arg(name)::text || '%')
ORDER BY sqlc.arg(ord)::text
OFFSET sqlc.arg(page)::int * sqlc.arg(size)::int LIMIT sqlc.arg(size)::int;

-- name: DeleteProject :exec
DELETE
FROM project
WHERE project.id = $1;

-- name: UpdateProject :exec
UPDATE project
SET name         = $2,
    description  = $3,
    start_date   = $4,
    turnout_date = $5
WHERE id = $1
RETURNING *;

-- name: CreateProject :one
INSERT INTO project (name, description, start_date, turnout_date)
VALUES ($1, $2, $3, $4)
RETURNING *;

