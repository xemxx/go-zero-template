package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ NoteModel = (*customNoteModel)(nil)

type (
	// NoteModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNoteModel.
	NoteModel interface {
		noteModel
	}

	customNoteModel struct {
		*defaultNoteModel
	}
)

// NewNoteModel returns a model for the database table.
func NewNoteModel(conn sqlx.SqlConn) NoteModel {
	return &customNoteModel{
		defaultNoteModel: newNoteModel(conn),
	}
}
