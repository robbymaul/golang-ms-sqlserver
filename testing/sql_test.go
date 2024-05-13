package testing

import (
	"context"
	"database/sql"
	"golang-mssql/connection"
	"golang-mssql/model"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	ctx := context.Background()

	conn, err := connection.GetConnection()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		t.Fatal(err.Error())
	}

	user := model.User{
		Username: "robby",
		Password: "robby",
	}

	query := "INSERT INTO dbo.users (username, password) VALUES (@username, @password)"

	_, err = tx.ExecContext(ctx, query, sql.Named("username", user.Username), sql.Named("password", user.Password))
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
	}

	tx.Commit()

}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	user := model.User{
		Username: "robby",
		Password: "ganteng",
	}
	query := "UPDATE users SET password=@password WHERE username=@username"

	conn, err := connection.GetConnection()
	if err != nil {
		t.Fatal(err.Error())
	}

	tx, err := conn.Begin()
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = tx.ExecContext(ctx, query, sql.Named("password", user.Password), sql.Named("username", user.Username))
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
	}

	tx.Commit()
}

func TestSelect(t *testing.T) {
	ctx := context.Background()
	var user model.User
	query := "SELECT username, password FROM users WHERE username=@username"

	conn, err := connection.GetConnection()
	if err != nil {
		t.Fatal(err.Error())
	}

	tx, err := conn.Begin()
	if err != nil {
		t.Fatal(err.Error())
	}

	rows, err := tx.QueryContext(ctx, query, sql.Named("username", "robby"))
	if err != nil {
		tx.Rollback()
	}

	if rows.Next() {
		rows.Scan(&user.Username, &user.Password)
	}

	assert.Equal(t, "robby", user.Username)
	assert.Equal(t, "ganteng", user.Password)
}

func TestDelete(t *testing.T) {
	ctx := context.Background()
	query := "DELETE FROM users"

	conn, err := connection.GetConnection()
	if err != nil {
		t.Fatal(err.Error())
	}

	tx, err := conn.Begin()
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
	}

	tx.Commit()

}
