package sqlconnect

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_pass, host, db_port, db_name)
	database, err := sql.Open("mysql", connectionStr)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to database")
	return database, nil
}
