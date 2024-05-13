package testing

import (
	"golang-mssql/connection"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	db, _ := connection.GetConnection()

	assert.NotNil(t, db)
}

func TestConnectionFailed(t *testing.T) {
	_, err := connection.GetConnection()

	assert.NotNil(t, err)

}
