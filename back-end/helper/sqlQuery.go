package helper

import (
	"database/sql"
)

func ExecuteQuery(query string, arg ...interface{}) (rows *sql.Rows, err error) {
	rows, err = SqlHandle.Query(query, arg...)
	return
}
