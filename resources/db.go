package resources

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	"os"
)

var (
	db *sql.DB
	err error
)

func ConnectDB() {
	url := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?tls=skip-verify",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST")+":"+os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	db, err = sql.Open(os.Getenv("DRIVER"), url)

	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() error {
	return db.Close()
}