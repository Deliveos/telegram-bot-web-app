// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package sqlc

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createGroupStmt, err = db.PrepareContext(ctx, createGroup); err != nil {
		return nil, fmt.Errorf("error preparing query CreateGroup: %w", err)
	}
	if q.createNotificationStmt, err = db.PrepareContext(ctx, createNotification); err != nil {
		return nil, fmt.Errorf("error preparing query CreateNotification: %w", err)
	}
	if q.createReplacementStmt, err = db.PrepareContext(ctx, createReplacement); err != nil {
		return nil, fmt.Errorf("error preparing query CreateReplacement: %w", err)
	}
	if q.createScheduleStmt, err = db.PrepareContext(ctx, createSchedule); err != nil {
		return nil, fmt.Errorf("error preparing query CreateSchedule: %w", err)
	}
	if q.createSubjectStmt, err = db.PrepareContext(ctx, createSubject); err != nil {
		return nil, fmt.Errorf("error preparing query CreateSubject: %w", err)
	}
	if q.deleteGroupStmt, err = db.PrepareContext(ctx, deleteGroup); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteGroup: %w", err)
	}
	if q.deleteNotificationStmt, err = db.PrepareContext(ctx, deleteNotification); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteNotification: %w", err)
	}
	if q.deleteReplacementStmt, err = db.PrepareContext(ctx, deleteReplacement); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteReplacement: %w", err)
	}
	if q.deleteScheduleStmt, err = db.PrepareContext(ctx, deleteSchedule); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSchedule: %w", err)
	}
	if q.deleteSubjectStmt, err = db.PrepareContext(ctx, deleteSubject); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSubject: %w", err)
	}
	if q.getAllGroupsStmt, err = db.PrepareContext(ctx, getAllGroups); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllGroups: %w", err)
	}
	if q.getAllReplacementsStmt, err = db.PrepareContext(ctx, getAllReplacements); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllReplacements: %w", err)
	}
	if q.getAllScheduleStmt, err = db.PrepareContext(ctx, getAllSchedule); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllSchedule: %w", err)
	}
	if q.getAllSubjectsStmt, err = db.PrepareContext(ctx, getAllSubjects); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllSubjects: %w", err)
	}
	if q.getGroupByIDStmt, err = db.PrepareContext(ctx, getGroupByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetGroupByID: %w", err)
	}
	if q.getGroupByNameStmt, err = db.PrepareContext(ctx, getGroupByName); err != nil {
		return nil, fmt.Errorf("error preparing query GetGroupByName: %w", err)
	}
	if q.getNotificationByIDStmt, err = db.PrepareContext(ctx, getNotificationByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetNotificationByID: %w", err)
	}
	if q.getReplacementByIDStmt, err = db.PrepareContext(ctx, getReplacementByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetReplacementByID: %w", err)
	}
	if q.getReplacementsByscheduleIDStmt, err = db.PrepareContext(ctx, getReplacementsByscheduleID); err != nil {
		return nil, fmt.Errorf("error preparing query GetReplacementsByscheduleID: %w", err)
	}
	if q.getScheduleByIDStmt, err = db.PrepareContext(ctx, getScheduleByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetScheduleByID: %w", err)
	}
	if q.getSchedulesByGroupIDStmt, err = db.PrepareContext(ctx, getSchedulesByGroupID); err != nil {
		return nil, fmt.Errorf("error preparing query GetSchedulesByGroupID: %w", err)
	}
	if q.getSubjectByIDStmt, err = db.PrepareContext(ctx, getSubjectByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetSubjectByID: %w", err)
	}
	if q.getSubjectByNameStmt, err = db.PrepareContext(ctx, getSubjectByName); err != nil {
		return nil, fmt.Errorf("error preparing query GetSubjectByName: %w", err)
	}
	if q.updateGroupStmt, err = db.PrepareContext(ctx, updateGroup); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateGroup: %w", err)
	}
	if q.updateNotificationStmt, err = db.PrepareContext(ctx, updateNotification); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateNotification: %w", err)
	}
	if q.updateReplacementStmt, err = db.PrepareContext(ctx, updateReplacement); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateReplacement: %w", err)
	}
	if q.updateScheduleStmt, err = db.PrepareContext(ctx, updateSchedule); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateSchedule: %w", err)
	}
	if q.updateSubjectStmt, err = db.PrepareContext(ctx, updateSubject); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateSubject: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createGroupStmt != nil {
		if cerr := q.createGroupStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createGroupStmt: %w", cerr)
		}
	}
	if q.createNotificationStmt != nil {
		if cerr := q.createNotificationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createNotificationStmt: %w", cerr)
		}
	}
	if q.createReplacementStmt != nil {
		if cerr := q.createReplacementStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createReplacementStmt: %w", cerr)
		}
	}
	if q.createScheduleStmt != nil {
		if cerr := q.createScheduleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createScheduleStmt: %w", cerr)
		}
	}
	if q.createSubjectStmt != nil {
		if cerr := q.createSubjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createSubjectStmt: %w", cerr)
		}
	}
	if q.deleteGroupStmt != nil {
		if cerr := q.deleteGroupStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteGroupStmt: %w", cerr)
		}
	}
	if q.deleteNotificationStmt != nil {
		if cerr := q.deleteNotificationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteNotificationStmt: %w", cerr)
		}
	}
	if q.deleteReplacementStmt != nil {
		if cerr := q.deleteReplacementStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteReplacementStmt: %w", cerr)
		}
	}
	if q.deleteScheduleStmt != nil {
		if cerr := q.deleteScheduleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteScheduleStmt: %w", cerr)
		}
	}
	if q.deleteSubjectStmt != nil {
		if cerr := q.deleteSubjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteSubjectStmt: %w", cerr)
		}
	}
	if q.getAllGroupsStmt != nil {
		if cerr := q.getAllGroupsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllGroupsStmt: %w", cerr)
		}
	}
	if q.getAllReplacementsStmt != nil {
		if cerr := q.getAllReplacementsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllReplacementsStmt: %w", cerr)
		}
	}
	if q.getAllScheduleStmt != nil {
		if cerr := q.getAllScheduleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllScheduleStmt: %w", cerr)
		}
	}
	if q.getAllSubjectsStmt != nil {
		if cerr := q.getAllSubjectsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllSubjectsStmt: %w", cerr)
		}
	}
	if q.getGroupByIDStmt != nil {
		if cerr := q.getGroupByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGroupByIDStmt: %w", cerr)
		}
	}
	if q.getGroupByNameStmt != nil {
		if cerr := q.getGroupByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGroupByNameStmt: %w", cerr)
		}
	}
	if q.getNotificationByIDStmt != nil {
		if cerr := q.getNotificationByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getNotificationByIDStmt: %w", cerr)
		}
	}
	if q.getReplacementByIDStmt != nil {
		if cerr := q.getReplacementByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getReplacementByIDStmt: %w", cerr)
		}
	}
	if q.getReplacementsByscheduleIDStmt != nil {
		if cerr := q.getReplacementsByscheduleIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getReplacementsByscheduleIDStmt: %w", cerr)
		}
	}
	if q.getScheduleByIDStmt != nil {
		if cerr := q.getScheduleByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getScheduleByIDStmt: %w", cerr)
		}
	}
	if q.getSchedulesByGroupIDStmt != nil {
		if cerr := q.getSchedulesByGroupIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSchedulesByGroupIDStmt: %w", cerr)
		}
	}
	if q.getSubjectByIDStmt != nil {
		if cerr := q.getSubjectByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSubjectByIDStmt: %w", cerr)
		}
	}
	if q.getSubjectByNameStmt != nil {
		if cerr := q.getSubjectByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSubjectByNameStmt: %w", cerr)
		}
	}
	if q.updateGroupStmt != nil {
		if cerr := q.updateGroupStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateGroupStmt: %w", cerr)
		}
	}
	if q.updateNotificationStmt != nil {
		if cerr := q.updateNotificationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateNotificationStmt: %w", cerr)
		}
	}
	if q.updateReplacementStmt != nil {
		if cerr := q.updateReplacementStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateReplacementStmt: %w", cerr)
		}
	}
	if q.updateScheduleStmt != nil {
		if cerr := q.updateScheduleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateScheduleStmt: %w", cerr)
		}
	}
	if q.updateSubjectStmt != nil {
		if cerr := q.updateSubjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateSubjectStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                              DBTX
	tx                              *sql.Tx
	createGroupStmt                 *sql.Stmt
	createNotificationStmt          *sql.Stmt
	createReplacementStmt           *sql.Stmt
	createScheduleStmt              *sql.Stmt
	createSubjectStmt               *sql.Stmt
	deleteGroupStmt                 *sql.Stmt
	deleteNotificationStmt          *sql.Stmt
	deleteReplacementStmt           *sql.Stmt
	deleteScheduleStmt              *sql.Stmt
	deleteSubjectStmt               *sql.Stmt
	getAllGroupsStmt                *sql.Stmt
	getAllReplacementsStmt          *sql.Stmt
	getAllScheduleStmt              *sql.Stmt
	getAllSubjectsStmt              *sql.Stmt
	getGroupByIDStmt                *sql.Stmt
	getGroupByNameStmt              *sql.Stmt
	getNotificationByIDStmt         *sql.Stmt
	getReplacementByIDStmt          *sql.Stmt
	getReplacementsByscheduleIDStmt *sql.Stmt
	getScheduleByIDStmt             *sql.Stmt
	getSchedulesByGroupIDStmt       *sql.Stmt
	getSubjectByIDStmt              *sql.Stmt
	getSubjectByNameStmt            *sql.Stmt
	updateGroupStmt                 *sql.Stmt
	updateNotificationStmt          *sql.Stmt
	updateReplacementStmt           *sql.Stmt
	updateScheduleStmt              *sql.Stmt
	updateSubjectStmt               *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                              tx,
		tx:                              tx,
		createGroupStmt:                 q.createGroupStmt,
		createNotificationStmt:          q.createNotificationStmt,
		createReplacementStmt:           q.createReplacementStmt,
		createScheduleStmt:              q.createScheduleStmt,
		createSubjectStmt:               q.createSubjectStmt,
		deleteGroupStmt:                 q.deleteGroupStmt,
		deleteNotificationStmt:          q.deleteNotificationStmt,
		deleteReplacementStmt:           q.deleteReplacementStmt,
		deleteScheduleStmt:              q.deleteScheduleStmt,
		deleteSubjectStmt:               q.deleteSubjectStmt,
		getAllGroupsStmt:                q.getAllGroupsStmt,
		getAllReplacementsStmt:          q.getAllReplacementsStmt,
		getAllScheduleStmt:              q.getAllScheduleStmt,
		getAllSubjectsStmt:              q.getAllSubjectsStmt,
		getGroupByIDStmt:                q.getGroupByIDStmt,
		getGroupByNameStmt:              q.getGroupByNameStmt,
		getNotificationByIDStmt:         q.getNotificationByIDStmt,
		getReplacementByIDStmt:          q.getReplacementByIDStmt,
		getReplacementsByscheduleIDStmt: q.getReplacementsByscheduleIDStmt,
		getScheduleByIDStmt:             q.getScheduleByIDStmt,
		getSchedulesByGroupIDStmt:       q.getSchedulesByGroupIDStmt,
		getSubjectByIDStmt:              q.getSubjectByIDStmt,
		getSubjectByNameStmt:            q.getSubjectByNameStmt,
		updateGroupStmt:                 q.updateGroupStmt,
		updateNotificationStmt:          q.updateNotificationStmt,
		updateReplacementStmt:           q.updateReplacementStmt,
		updateScheduleStmt:              q.updateScheduleStmt,
		updateSubjectStmt:               q.updateSubjectStmt,
	}
}
