package mysql_utils

import (
	"errors"
	"github.com/go-sql-driver/mysql"
  "github.com/katsun0921/go_utils/logger"
  "github.com/katsun0921/go_utils/rest_errors"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
      logger.Error("error when connect to mysql", err)
			return rest_errors.NewNotFoundError("no record matching given id", errors.New("database error"))
		}
    logger.Error("error when trying to mysql", err)
		return rest_errors.NewInternalServerError("error parsing database mysql response", errors.New("database error"))
	}

	switch sqlErr.Number {
	case 1062:
    logger.Error("error sql 1062 to mysql", err)
		return rest_errors.NewBadRequestError("invalid data error", errors.New("database error"))
	}
  logger.Error("error processing request", err)
	return rest_errors.NewInternalServerError("error processing request", errors.New("database error"))
}
