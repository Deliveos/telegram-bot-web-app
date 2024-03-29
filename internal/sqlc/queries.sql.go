// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: queries.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createGroup = `-- name: CreateGroup :one
INSERT INTO schedules.groups (name, chat_id) VALUES ($1, $2) RETURNING id
`

type CreateGroupParams struct {
	Name   sql.NullString `json:"name"`
	ChatID sql.NullInt32  `json:"chat_id"`
}

func (q *Queries) CreateGroup(ctx context.Context, arg CreateGroupParams) (int32, error) {
	row := q.queryRow(ctx, q.createGroupStmt, createGroup, arg.Name, arg.ChatID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createNotification = `-- name: CreateNotification :one
INSERT INTO schedules.notifications (chat_id, text, datetime) VALUES ($1, $2, $3) RETURNING id
`

type CreateNotificationParams struct {
	ChatID   sql.NullInt32 `json:"chat_id"`
	Text     sql.NullInt32 `json:"text"`
	Datetime sql.NullTime  `json:"datetime"`
}

func (q *Queries) CreateNotification(ctx context.Context, arg CreateNotificationParams) (int32, error) {
	row := q.queryRow(ctx, q.createNotificationStmt, createNotification, arg.ChatID, arg.Text, arg.Datetime)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createReplacement = `-- name: CreateReplacement :one
INSERT INTO schedules.replacements (schedule_id, subject_id, datetime) VALUES ($1, $2, $3) RETURNING id
`

type CreateReplacementParams struct {
	ScheduleID sql.NullInt32 `json:"schedule_id"`
	SubjectID  sql.NullInt32 `json:"subject_id"`
	Datetime   sql.NullTime  `json:"datetime"`
}

func (q *Queries) CreateReplacement(ctx context.Context, arg CreateReplacementParams) (int32, error) {
	row := q.queryRow(ctx, q.createReplacementStmt, createReplacement, arg.ScheduleID, arg.SubjectID, arg.Datetime)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createSchedule = `-- name: CreateSchedule :one
INSERT INTO schedules.schedules (group_id, subject_id, weekday, time) VALUES ($1, $2, $3, $4) RETURNING id
`

type CreateScheduleParams struct {
	GroupID   sql.NullInt32 `json:"group_id"`
	SubjectID sql.NullInt32 `json:"subject_id"`
	Weekday   sql.NullInt32 `json:"weekday"`
	Time      sql.NullTime  `json:"time"`
}

func (q *Queries) CreateSchedule(ctx context.Context, arg CreateScheduleParams) (int32, error) {
	row := q.queryRow(ctx, q.createScheduleStmt, createSchedule,
		arg.GroupID,
		arg.SubjectID,
		arg.Weekday,
		arg.Time,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createSubject = `-- name: CreateSubject :one
INSERT INTO schedules.subjects (name, teacher, classrom) VALUES ($1, $2, $3) RETURNING id
`

type CreateSubjectParams struct {
	Name     sql.NullString `json:"name"`
	Teacher  sql.NullString `json:"teacher"`
	Classrom sql.NullString `json:"classrom"`
}

func (q *Queries) CreateSubject(ctx context.Context, arg CreateSubjectParams) (int32, error) {
	row := q.queryRow(ctx, q.createSubjectStmt, createSubject, arg.Name, arg.Teacher, arg.Classrom)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteGroup = `-- name: DeleteGroup :exec
DELETE FROM schedules.groups WHERE id = $1
`

func (q *Queries) DeleteGroup(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteGroupStmt, deleteGroup, id)
	return err
}

const deleteNotification = `-- name: DeleteNotification :exec
DELETE FROM schedules.notifications WHERE id = $1
`

func (q *Queries) DeleteNotification(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteNotificationStmt, deleteNotification, id)
	return err
}

const deleteReplacement = `-- name: DeleteReplacement :exec
DELETE FROM schedules.replacements WHERE id = $1
`

func (q *Queries) DeleteReplacement(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteReplacementStmt, deleteReplacement, id)
	return err
}

const deleteSchedule = `-- name: DeleteSchedule :exec
DELETE FROM schedules.schedules WHERE id = $1
`

func (q *Queries) DeleteSchedule(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteScheduleStmt, deleteSchedule, id)
	return err
}

const deleteSubject = `-- name: DeleteSubject :exec
DELETE FROM schedules.subjects WHERE id = $1
`

func (q *Queries) DeleteSubject(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteSubjectStmt, deleteSubject, id)
	return err
}

const getAllGroups = `-- name: GetAllGroups :many
SELECT id, name, chat_id FROM schedules.groups
`

func (q *Queries) GetAllGroups(ctx context.Context) ([]SchedulesGroup, error) {
	rows, err := q.query(ctx, q.getAllGroupsStmt, getAllGroups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SchedulesGroup
	for rows.Next() {
		var i SchedulesGroup
		if err := rows.Scan(&i.ID, &i.Name, &i.ChatID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllReplacements = `-- name: GetAllReplacements :many
SELECT id, schedule_id, subject_id, datetime FROM schedules.replacements
`

func (q *Queries) GetAllReplacements(ctx context.Context) ([]SchedulesReplacement, error) {
	rows, err := q.query(ctx, q.getAllReplacementsStmt, getAllReplacements)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SchedulesReplacement
	for rows.Next() {
		var i SchedulesReplacement
		if err := rows.Scan(
			&i.ID,
			&i.ScheduleID,
			&i.SubjectID,
			&i.Datetime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllSchedule = `-- name: GetAllSchedule :many
SELECT id, group_id, subject_id, weekday, time FROM schedules.schedules
`

func (q *Queries) GetAllSchedule(ctx context.Context) ([]SchedulesSchedule, error) {
	rows, err := q.query(ctx, q.getAllScheduleStmt, getAllSchedule)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SchedulesSchedule
	for rows.Next() {
		var i SchedulesSchedule
		if err := rows.Scan(
			&i.ID,
			&i.GroupID,
			&i.SubjectID,
			&i.Weekday,
			&i.Time,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllSubjects = `-- name: GetAllSubjects :many
SELECT id, name, teacher, classrom FROM schedules.subjects
`

func (q *Queries) GetAllSubjects(ctx context.Context) ([]SchedulesSubject, error) {
	rows, err := q.query(ctx, q.getAllSubjectsStmt, getAllSubjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SchedulesSubject
	for rows.Next() {
		var i SchedulesSubject
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Teacher,
			&i.Classrom,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGroupByID = `-- name: GetGroupByID :one
SELECT id, name, chat_id FROM schedules.groups WHERE id = $1
`

func (q *Queries) GetGroupByID(ctx context.Context, id int32) (SchedulesGroup, error) {
	row := q.queryRow(ctx, q.getGroupByIDStmt, getGroupByID, id)
	var i SchedulesGroup
	err := row.Scan(&i.ID, &i.Name, &i.ChatID)
	return i, err
}

const getGroupByName = `-- name: GetGroupByName :one
SELECT id, name, chat_id FROM schedules.groups WHERE "name"=$1
`

func (q *Queries) GetGroupByName(ctx context.Context, name sql.NullString) (SchedulesGroup, error) {
	row := q.queryRow(ctx, q.getGroupByNameStmt, getGroupByName, name)
	var i SchedulesGroup
	err := row.Scan(&i.ID, &i.Name, &i.ChatID)
	return i, err
}

const getNotificationByID = `-- name: GetNotificationByID :one
SELECT id, chat_id, text, datetime FROM schedules.notifications WHERE id = $1
`

func (q *Queries) GetNotificationByID(ctx context.Context, id int32) (SchedulesNotification, error) {
	row := q.queryRow(ctx, q.getNotificationByIDStmt, getNotificationByID, id)
	var i SchedulesNotification
	err := row.Scan(
		&i.ID,
		&i.ChatID,
		&i.Text,
		&i.Datetime,
	)
	return i, err
}

const getReplacementByID = `-- name: GetReplacementByID :one
SELECT id, schedule_id, subject_id, datetime FROM schedules.replacements WHERE id = $1
`

func (q *Queries) GetReplacementByID(ctx context.Context, id int32) (SchedulesReplacement, error) {
	row := q.queryRow(ctx, q.getReplacementByIDStmt, getReplacementByID, id)
	var i SchedulesReplacement
	err := row.Scan(
		&i.ID,
		&i.ScheduleID,
		&i.SubjectID,
		&i.Datetime,
	)
	return i, err
}

const getReplacementsByscheduleID = `-- name: GetReplacementsByscheduleID :one
SELECT id, schedule_id, subject_id, datetime FROM schedules.replacements WHERE schedule_id=$1
`

func (q *Queries) GetReplacementsByscheduleID(ctx context.Context, scheduleID sql.NullInt32) (SchedulesReplacement, error) {
	row := q.queryRow(ctx, q.getReplacementsByscheduleIDStmt, getReplacementsByscheduleID, scheduleID)
	var i SchedulesReplacement
	err := row.Scan(
		&i.ID,
		&i.ScheduleID,
		&i.SubjectID,
		&i.Datetime,
	)
	return i, err
}

const getScheduleByID = `-- name: GetScheduleByID :one
SELECT id, group_id, subject_id, weekday, time FROM schedules.schedules WHERE id = $1
`

func (q *Queries) GetScheduleByID(ctx context.Context, id int32) (SchedulesSchedule, error) {
	row := q.queryRow(ctx, q.getScheduleByIDStmt, getScheduleByID, id)
	var i SchedulesSchedule
	err := row.Scan(
		&i.ID,
		&i.GroupID,
		&i.SubjectID,
		&i.Weekday,
		&i.Time,
	)
	return i, err
}

const getSchedulesByGroupID = `-- name: GetSchedulesByGroupID :one
SELECT id, group_id, subject_id, weekday, time FROM schedules.schedules WHERE group_id=$1
`

func (q *Queries) GetSchedulesByGroupID(ctx context.Context, groupID sql.NullInt32) (SchedulesSchedule, error) {
	row := q.queryRow(ctx, q.getSchedulesByGroupIDStmt, getSchedulesByGroupID, groupID)
	var i SchedulesSchedule
	err := row.Scan(
		&i.ID,
		&i.GroupID,
		&i.SubjectID,
		&i.Weekday,
		&i.Time,
	)
	return i, err
}

const getSubjectByID = `-- name: GetSubjectByID :one
SELECT id, name, teacher, classrom FROM schedules.subjects WHERE id = $1
`

func (q *Queries) GetSubjectByID(ctx context.Context, id int32) (SchedulesSubject, error) {
	row := q.queryRow(ctx, q.getSubjectByIDStmt, getSubjectByID, id)
	var i SchedulesSubject
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Teacher,
		&i.Classrom,
	)
	return i, err
}

const getSubjectByName = `-- name: GetSubjectByName :one
SELECT id, name, teacher, classrom FROM schedules.subjects WHERE "name"=$1
`

func (q *Queries) GetSubjectByName(ctx context.Context, name sql.NullString) (SchedulesSubject, error) {
	row := q.queryRow(ctx, q.getSubjectByNameStmt, getSubjectByName, name)
	var i SchedulesSubject
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Teacher,
		&i.Classrom,
	)
	return i, err
}

const updateGroup = `-- name: UpdateGroup :exec
UPDATE schedules.groups SET name = $2, chat_id = $3 WHERE id = $1
`

type UpdateGroupParams struct {
	ID     int32          `json:"id"`
	Name   sql.NullString `json:"name"`
	ChatID sql.NullInt32  `json:"chat_id"`
}

func (q *Queries) UpdateGroup(ctx context.Context, arg UpdateGroupParams) error {
	_, err := q.exec(ctx, q.updateGroupStmt, updateGroup, arg.ID, arg.Name, arg.ChatID)
	return err
}

const updateNotification = `-- name: UpdateNotification :exec
UPDATE schedules.notifications SET chat_id = $2, text = $3, datetime = $4 WHERE id = $1
`

type UpdateNotificationParams struct {
	ID       int32         `json:"id"`
	ChatID   sql.NullInt32 `json:"chat_id"`
	Text     sql.NullInt32 `json:"text"`
	Datetime sql.NullTime  `json:"datetime"`
}

func (q *Queries) UpdateNotification(ctx context.Context, arg UpdateNotificationParams) error {
	_, err := q.exec(ctx, q.updateNotificationStmt, updateNotification,
		arg.ID,
		arg.ChatID,
		arg.Text,
		arg.Datetime,
	)
	return err
}

const updateReplacement = `-- name: UpdateReplacement :exec
UPDATE schedules.replacements SET schedule_id = $2, subject_id = $3, datetime = $4 WHERE id = $1
`

type UpdateReplacementParams struct {
	ID         int32         `json:"id"`
	ScheduleID sql.NullInt32 `json:"schedule_id"`
	SubjectID  sql.NullInt32 `json:"subject_id"`
	Datetime   sql.NullTime  `json:"datetime"`
}

func (q *Queries) UpdateReplacement(ctx context.Context, arg UpdateReplacementParams) error {
	_, err := q.exec(ctx, q.updateReplacementStmt, updateReplacement,
		arg.ID,
		arg.ScheduleID,
		arg.SubjectID,
		arg.Datetime,
	)
	return err
}

const updateSchedule = `-- name: UpdateSchedule :exec
UPDATE schedules.schedules SET group_id = $2, subject_id = $3, weekday = $4, time = $5 WHERE id = $1
`

type UpdateScheduleParams struct {
	ID        int32         `json:"id"`
	GroupID   sql.NullInt32 `json:"group_id"`
	SubjectID sql.NullInt32 `json:"subject_id"`
	Weekday   sql.NullInt32 `json:"weekday"`
	Time      sql.NullTime  `json:"time"`
}

func (q *Queries) UpdateSchedule(ctx context.Context, arg UpdateScheduleParams) error {
	_, err := q.exec(ctx, q.updateScheduleStmt, updateSchedule,
		arg.ID,
		arg.GroupID,
		arg.SubjectID,
		arg.Weekday,
		arg.Time,
	)
	return err
}

const updateSubject = `-- name: UpdateSubject :exec
UPDATE schedules.subjects SET name = $2, teacher = $3, classrom = $4 WHERE id = $1
`

type UpdateSubjectParams struct {
	ID       int32          `json:"id"`
	Name     sql.NullString `json:"name"`
	Teacher  sql.NullString `json:"teacher"`
	Classrom sql.NullString `json:"classrom"`
}

func (q *Queries) UpdateSubject(ctx context.Context, arg UpdateSubjectParams) error {
	_, err := q.exec(ctx, q.updateSubjectStmt, updateSubject,
		arg.ID,
		arg.Name,
		arg.Teacher,
		arg.Classrom,
	)
	return err
}
