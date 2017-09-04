package persistence

import "testing"
import "database/sql"

func failIfError(t *testing.T, err error, format string, args ...interface{}) {
	if err != nil {
		t.Errorf(format, args...)
	}
}

func failIfOK(t *testing.T, err error, format string, args ...interface{}) {
	if err == nil {
		t.Errorf(format, args...)
	}
}

func createConnection(t *testing.T) *sql.DB {
	connection, err := sql.Open("sqlite3", ":memory:")
	failIfError(t, err, "unable to create fixture DB: %s", err)

	return connection
}

func createProvider(t *testing.T, connection *sql.DB) (p Provider) {
	var err error
	if p, err = FromDbConnection(connection); err != nil {
		t.Errorf("failed to create DB: %s", err)
	}

	return
}

func TestCreateDB(t *testing.T) {
	p, err := FromSqlite(":memory:")
	failIfError(t, err, "failed to create DB: %s", err)

	failIfError(t, p.Close(), "failed to close DB: %s", err)
}

func TestOpenDB(t *testing.T) {
	connection := createConnection(t)
	defer connection.Close()

	var err error
	if _, err = FromDbConnection(connection); err != nil {
		t.Errorf("failed to create DB: %s", err)
	}

	if _, err = FromDbConnection(connection); err != nil {
		t.Errorf("failed to open DB: %s", err)
	}
}

func TestFailIfVersionOutdated(t *testing.T) {
	connection := createConnection(t)
	defer connection.Close()

	createProvider(t, connection)

	var err error

	_, err = connection.Exec("UPDATE meta SET version = ?", SCHEMA_VERSION-1)
	failIfError(t, err, "SQL error: %s", err)

	_, err = FromDbConnection(connection)
	failIfOK(t, err, "opening an outdated version should fail")
}

func TestFailIfVersionFuture(t *testing.T) {
	connection := createConnection(t)
	defer connection.Close()

	createProvider(t, connection)

	var err error

	_, err = connection.Exec("UPDATE meta SET version = ?", SCHEMA_VERSION+1)
	failIfError(t, err, "SQL error: %s", err)

	_, err = FromDbConnection(connection)
	failIfOK(t, err, "opening a future version should fail")
}

func TestFailIfMultipleMetadataEntries(t *testing.T) {
	connection := createConnection(t)
	defer connection.Close()

	createProvider(t, connection)

	var err error

	_, err = connection.Exec("INSERT INTO meta (version) VALUES (?)", SCHEMA_VERSION)
	failIfError(t, err, "SQL error: %s", err)

	_, err = FromDbConnection(connection)
	failIfOK(t, err, "opening should fail if the meta table contains multiple entries")
}
