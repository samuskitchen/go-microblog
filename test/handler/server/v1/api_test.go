package v1

import (
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	data "microblog/database"
	"microblog/handler/server/v1"
	"testing"
)

var (
	connMock *data.Data
)

func TestNew(t *testing.T) {

	db, _, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	connMock = &data.Data{
		DB: db,
	}

	defer func() {
		db.Close()
	}()

	v1.New(connMock)
}
