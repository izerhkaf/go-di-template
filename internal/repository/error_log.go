package repository

import (
	"template/internal/model/dbmodel"

	"github.com/jmoiron/sqlx"
)

type ErrorLogRepository interface {
	InsertErrorLog(payload dbmodel.ErrorLog)
}

type errorLogRepository struct {
	DB *sqlx.DB
}

func NewErrorLogRepository(db *sqlx.DB) ErrorLogRepository {
	return &errorLogRepository{DB: db}
}

func (r *errorLogRepository) InsertErrorLog(payload dbmodel.ErrorLog) {
	query := `INSERT INTO error_logs  
	(
		error_message,
	) VALUES (
		:error_message,
	)`

	r.DB.NamedExec(query, payload)
}
