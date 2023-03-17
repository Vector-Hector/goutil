package util

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
	"runtime/debug"
	"strings"
)

type HttpError struct {
	StatusCode int
	Message    string
}

func (err HttpError) Error() string {
	return fmt.Sprintf("http error %d: %s", err.StatusCode, err.Message)
}

func BadRequest(message ...interface{}) error {
	return HttpError{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprintln(message...),
	}
}

func PanicBadRequest(message ...interface{}) {
	panic(BadRequest(message...))
}

func InternalServerError(message ...interface{}) error {
	return HttpError{
		StatusCode: http.StatusInternalServerError,
		Message:    fmt.Sprintln(message...),
	}
}

func PanicInternalServerError(message ...interface{}) {
	panic(InternalServerError(message...))
}

func NewHttpError(code int, message ...interface{}) error {
	return HttpError{
		StatusCode: code,
		Message:    fmt.Sprintln(message...),
	}
}

func PanicHttp(code int, message ...interface{}) {
	panic(NewHttpError(code, message...))
}

func HandleErrors(c *gin.Context) {
	if r := recover(); r != nil {
		err := r.(error)

		if httpErr, ok := err.(HttpError); ok {
			c.JSON(httpErr.StatusCode, gin.H{
				"data": httpErr.Message,
			})
			return
		}

		// ignored errors:
		if strings.Contains(err.Error(), "An established connection was aborted by the software in your host machine") {
			return
		}
		fmt.Println(err)
		debug.PrintStack()

		if _, ok := err.(*mysql.MySQLError); ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"data": "Error while executing on database",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"data": "Internal error",
		})
		return
	}
}

func HandleUpdaterErrors() {
	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			err = errors.New(r.(string))
		}
		fmt.Println(err)
		debug.PrintStack()
		return
	}
}

func Rollback(tx *sqlx.Tx) { // defer rollback
	if r := recover(); r != nil {
		_ = tx.Rollback()
		panic(r) // fall back to default error handling
	} else {
		_ = tx.Commit()
	}
}
