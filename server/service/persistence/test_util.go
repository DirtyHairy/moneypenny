package persistence

import (
	"database/sql"
	"testing"
)

func failIfError(t *testing.T, err error, format string, args ...interface{}) {
	if err != nil {
		t.Fatalf(format, args...)
	}
}

func failIfOK(t *testing.T, err error, format string, args ...interface{}) {
	if err == nil {
		t.Fatalf(format, args...)
	}
}

func createConnection(t *testing.T) *sql.DB {
	connection, err := sql.Open("sqlite3", ":memory:")
	failIfError(t, err, "unable to create fixture DB: %s", err)

	return connection
}

func createProviderFromConnection(t *testing.T, connection *sql.DB) (p Provider) {
	var err error
	if p, err = FromDbConnection(connection); err != nil {
		t.Errorf("failed to create DB: %s", err)
	}

	return
}

func createProvider(t *testing.T) (p Provider) {
	var err error
	if p, err = FromSqlite(":memory:"); err != nil {
		t.Errorf("failed to create DB: %s", err)
	}

	return
}
