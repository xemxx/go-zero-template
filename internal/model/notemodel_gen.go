// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	noteFieldNames          = builder.RawFieldNames(&Note{})
	noteRows                = strings.Join(noteFieldNames, ",")
	noteRowsExpectAutoSet   = strings.Join(stringx.Remove(noteFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	noteRowsWithPlaceHolder = strings.Join(stringx.Remove(noteFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	noteModel interface {
		Insert(ctx context.Context, data *Note) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Note, error)
		Update(ctx context.Context, data *Note) error
		Delete(ctx context.Context, id int64) error
	}

	defaultNoteModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Note struct {
		Id         int64          `db:"id"`
		Guid       sql.NullString `db:"guid"`
		Uid        sql.NullInt64  `db:"uid"`
		Bid        sql.NullString `db:"bid"`
		Title      sql.NullString `db:"title"`
		Content    sql.NullString `db:"content"`
		SC         sql.NullInt64  `db:"SC"`
		AddDate    sql.NullTime   `db:"addDate"`
		ModifyDate sql.NullTime   `db:"modifyDate"`
		IsDel      sql.NullInt64  `db:"isDel"`
	}
)

func newNoteModel(conn sqlx.SqlConn) *defaultNoteModel {
	return &defaultNoteModel{
		conn:  conn,
		table: "`note`",
	}
}

func (m *defaultNoteModel) withSession(session sqlx.Session) *defaultNoteModel {
	return &defaultNoteModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`note`",
	}
}

func (m *defaultNoteModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultNoteModel) FindOne(ctx context.Context, id int64) (*Note, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", noteRows, m.table)
	var resp Note
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultNoteModel) Insert(ctx context.Context, data *Note) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, noteRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Guid, data.Uid, data.Bid, data.Title, data.Content, data.SC, data.AddDate, data.ModifyDate, data.IsDel)
	return ret, err
}

func (m *defaultNoteModel) Update(ctx context.Context, data *Note) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, noteRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Guid, data.Uid, data.Bid, data.Title, data.Content, data.SC, data.AddDate, data.ModifyDate, data.IsDel, data.Id)
	return err
}

func (m *defaultNoteModel) tableName() string {
	return m.table
}