package dbmodel

type ErrorLog struct {
	ErrorMessage string `db:"error_message"`
}
