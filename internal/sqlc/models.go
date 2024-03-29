// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package sqlc

import (
	"database/sql"
)

type SchedulesGroup struct {
	ID     int32          `json:"id"`
	Name   sql.NullString `json:"name"`
	ChatID sql.NullInt32  `json:"chat_id"`
}

type SchedulesNotification struct {
	ID       int32         `json:"id"`
	ChatID   sql.NullInt32 `json:"chat_id"`
	Text     sql.NullInt32 `json:"text"`
	Datetime sql.NullTime  `json:"datetime"`
}

type SchedulesReplacement struct {
	ID         int32         `json:"id"`
	ScheduleID sql.NullInt32 `json:"schedule_id"`
	SubjectID  sql.NullInt32 `json:"subject_id"`
	Datetime   sql.NullTime  `json:"datetime"`
}

type SchedulesSchedule struct {
	ID        int32         `json:"id"`
	GroupID   sql.NullInt32 `json:"group_id"`
	SubjectID sql.NullInt32 `json:"subject_id"`
	Weekday   sql.NullInt32 `json:"weekday"`
	Time      sql.NullTime  `json:"time"`
}

type SchedulesSubject struct {
	ID       int32          `json:"id"`
	Name     sql.NullString `json:"name"`
	Teacher  sql.NullString `json:"teacher"`
	Classrom sql.NullString `json:"classrom"`
}

type UsersPermission struct {
	ID   int32          `json:"id"`
	Name sql.NullString `json:"name"`
}

type UsersPermissionsRole struct {
	RoleID       sql.NullInt32 `json:"role_id"`
	PermissionID sql.NullInt32 `json:"permission_id"`
}

type UsersRole struct {
	ID   int32          `json:"id"`
	Name sql.NullString `json:"name"`
}

type UsersUser struct {
	ID        int32          `json:"id"`
	Firstname sql.NullString `json:"firstname"`
	Lastname  sql.NullString `json:"lastname"`
	Chatid    sql.NullInt32  `json:"chatid"`
}

type UsersUsersRole struct {
	RoleID sql.NullInt32 `json:"role_id"`
	UserID sql.NullInt32 `json:"user_id"`
}
