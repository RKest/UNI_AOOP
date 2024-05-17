-- name: GetProjectByName :many
SELECT * FROM project WHERE upper(project.name) = upper("%$1%") LIMIT 1;

-- name: GetTasks :many
SELECT * FROM task WHERE task.project_id = $1;

-- name: GetAllStudentsByIndex :many
SELECT * FROM student WHERE student.index LIKE "$1%";

-- name: GetStudentByIndex :one
SELECT * FROM student WHERE student.index LIKE $1 LIMIT 1;

-- name: GetProject :one
SELECT * FROM project WHERE project.id = $1 LIMIT 1;

-- name: DeleteProject :exec
DELETE FROM project WHERE project.id = $1;

-- name: CreateProject :one
INSERT INTO project (name, description, start_date, turnout_date)
VALUES ($1, $2, $3, $4) RETURNING *;

