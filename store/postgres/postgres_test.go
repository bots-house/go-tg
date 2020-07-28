package postgres

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-txdb"
)

func init() {
	const env = "BRZ_TEST_DATABASE"

	dsn := os.Getenv(env)

	if dsn == "" {
		log.Panicf("'%s' is not provided, but required for integration tests", env)
	}

	txdb.Register("txdb", "postgres", dsn)
}

func newPostgres(t *testing.T) *Postgres {
	t.Helper()

	db, err := sql.Open("txdb", t.Name())
	if err != nil {
		t.Fatalf("can't open db: %s", err)
	}

	t.Cleanup(func() {
		db.Close()
	})

	return NewPostgres(db)
}
