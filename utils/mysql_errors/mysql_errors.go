package mysql_errors

import (
	"database/sql"
	"net/http"

	mysql "github.com/go-mysql/errors"

	"github.com/VividCortex/mysqlerr"
)

type MysqlError interface {
	Message() string
	Status() int64
	Code() string
}

type mysqlError struct {
	MEMessage string
	MEStatus  int64
	MECode    string
}

func (me *mysqlError) Message() string {
	return me.MEMessage
}

func (me *mysqlError) Status() int64 {
	return me.MEStatus
}

func (me *mysqlError) Code() string {
	return me.MECode
}

func DatabaseErrorHandler(err error) MysqlError {
	if ok, mysqlErr := mysql.Error(err); ok { //if it is the [MYSQL ERROR]
		switch mysqlErr {
		case mysql.ErrCannotConnect: //check if it is a connection error
			return &mysqlError{
				MEMessage: mysqlErr.Error(),
				MEStatus:  http.StatusRequestTimeout,
				MECode:    "cannot_connect",
			}
		case sql.ErrNoRows: //no data present
			return &mysqlError{
				MEMessage: "data not found",
				MEStatus:  http.StatusNotFound,
				MECode:    "no_data",
			}
		case mysql.ErrConnLost: //database connection lost
			return &mysqlError{
				MEMessage: mysqlErr.Error(),
				MEStatus:  http.StatusInternalServerError,
				MECode:    "connection_lost",
			}
		case mysql.ErrDupeKey: //same value or duplicate value
			return &mysqlError{
				MEMessage: "same value or value already existed",
				MEStatus:  http.StatusConflict,
				MECode:    "duplicate_data",
			}
		}

		//get the error code
		errorNumber := mysql.MySQLErrorCode(mysqlErr)
		if errorNumber == mysqlerr.ER_DUP_ENTRY {
			return &mysqlError{
				MEMessage: "same value or value already existed",
				MEStatus:  http.StatusConflict,
				MECode:    "duplicate_data",
			}
		}
	}

	//if it is not the [MYSQL ERROR]
	return &mysqlError{
		MEMessage: err.Error(),
		MEStatus:  http.StatusInternalServerError,
		MECode:    "not_mysql_error",
	}
}