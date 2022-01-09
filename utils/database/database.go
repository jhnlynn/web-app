package database

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"web-app/config"
	"web-app/utils/processErr"
)

func OpenDB(c *gin.Context) *sql.DB {

	db, err := sql.Open("mysql", config.DBConnection)
	processErr.ProcessInternalErr(c, err, "mysql opening error")

	return db
}

