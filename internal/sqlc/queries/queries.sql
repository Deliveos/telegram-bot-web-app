-- name: GetGroupByID :one
SELECT id, name, chat_id FROM schedules.groups WHERE id = $1;

-- name: CreateGroup :one
INSERT INTO schedules.groups (name, chat_id) VALUES ($1, $2) RETURNING id;

-- name: UpdateGroup :exec
UPDATE schedules.groups SET name = $2, chat_id = $3 WHERE id = $1;

-- name: DeleteGroup :exec
DELETE FROM schedules.groups WHERE id = $1;

-- name: GetAllGroups :many
SELECT * FROM schedules.groups;

-- name: GetGroupByName :one
SELECT * FROM schedules.groups WHERE "name"=$1;



-- name: GetSubjectByID :one
SELECT id, name, teacher, classrom FROM schedules.subjects WHERE id = $1;

-- name: CreateSubject :one
INSERT INTO schedules.subjects (name, teacher, classrom) VALUES ($1, $2, $3) RETURNING id;

-- name: UpdateSubject :exec
UPDATE schedules.subjects SET name = $2, teacher = $3, classrom = $4 WHERE id = $1;

-- name: DeleteSubject :exec
DELETE FROM schedules.subjects WHERE id = $1;

-- name: GetAllSubjects :many
SELECT * FROM schedules.subjects;

-- name: GetSubjectByName :one
SELECT * FROM schedules.subjects WHERE "name"=$1;



-- name: GetScheduleByID :one
SELECT id, group_id, subject_id, weekday, time FROM schedules.schedules WHERE id = $1;

-- name: CreateSchedule :one
INSERT INTO schedules.schedules (group_id, subject_id, weekday, time) VALUES ($1, $2, $3, $4) RETURNING id;

-- name: UpdateSchedule :exec
UPDATE schedules.schedules SET group_id = $2, subject_id = $3, weekday = $4, time = $5 WHERE id = $1;

-- name: DeleteSchedule :exec
DELETE FROM schedules.schedules WHERE id = $1;

-- name: GetAllSchedule :many
SELECT * FROM schedules.schedules;

-- name: GetSchedulesByGroupID :one
SELECT * FROM schedules.schedules WHERE group_id=$1;



-- name: GetReplacementByID :one
SELECT id, schedule_id, subject_id, datetime FROM schedules.replacements WHERE id = $1;

-- name: CreateReplacement :one
INSERT INTO schedules.replacements (schedule_id, subject_id, datetime) VALUES ($1, $2, $3) RETURNING id;

-- name: UpdateReplacement :exec
UPDATE schedules.replacements SET schedule_id = $2, subject_id = $3, datetime = $4 WHERE id = $1;

-- name: DeleteReplacement :exec
DELETE FROM schedules.replacements WHERE id = $1;

-- name: GetAllReplacements :many
SELECT * FROM schedules.replacements;

-- name: GetReplacementsByscheduleID :one
SELECT * FROM schedules.replacements WHERE schedule_id=$1;



-- name: GetNotificationByID :one
SELECT id, chat_id, text, datetime FROM schedules.notifications WHERE id = $1;

-- name: CreateNotification :one
INSERT INTO schedules.notifications (chat_id, text, datetime) VALUES ($1, $2, $3) RETURNING id;

-- name: UpdateNotification :exec
UPDATE schedules.notifications SET chat_id = $2, text = $3, datetime = $4 WHERE id = $1;

-- name: DeleteNotification :exec
DELETE FROM schedules.notifications WHERE id = $1;
