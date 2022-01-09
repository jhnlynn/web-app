package database

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"joke/config"
	"joke/utils/processErr"
)

func OpenDB(c *gin.Context) *sql.DB {

	db, err := sql.Open("mysql", config.DBConnection)
	processErr.ProcessInternalErr(c, err, "mysql opening error")

	return db
}

