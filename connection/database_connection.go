package connection

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

func GetConnection() (*sql.DB, error) {
	err := godotenv.Load("../.env.dev")
	if err != nil {
		log.Println(err.Error())
	}

	db_driver := os.Getenv("DB_DRIVER")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	dns := fmt.Sprintf(os.Getenv("DB_DNS"), db_driver, db_user, db_password, db_host, db_port, db_name)

	db, err := sql.Open(db_driver, dns)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	err = db.PingContext(context.Background())
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
