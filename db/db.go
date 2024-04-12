// Package db provides a connection to the database
package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

// Init initializes a new connection to the database
func Init() (*sql.DB, error) {
	addr := fmt.Sprintf("%s:%s", os.Getenv("DBHOST"), os.Getenv("DBPORT"))
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 addr,
		DBName:               os.Getenv("DBNAME"),
		AllowNativePasswords: true,
	}

	return sql.Open("mysql", cfg.FormatDSN())
}
