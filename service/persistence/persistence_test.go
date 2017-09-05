package persistence

import "testing"

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
		t.Fatalf("failed to create DB: %s", err)
	}

	if _, err = FromDbConnection(connection); err != nil {
		t.Fatalf("failed to open DB: %s", err)
	}
}

func TestFailIfVersionOutdated(t *testing.T) {
	connection := createConnection(t)
	defer connection.Close()

	createProviderFromConnection(t, connection)

	var err error

	_, err = connection.Exec("UPDATE meta SET version = ?", SCHEMA_VERSION-1)
	failIfError(t, err, "SQL error: %s", err)

	_, err = FromDbConnection(connection)
	failIfOK(t, err, "opening an outdated version should fail")
}

func TestFailIfVersionFuture(t *testing.T) {
	connection := createConnection(t)
	defer connection.Close()

	createProviderFromConnection(t, connection)

	var err error

	_, err = connection.Exec("UPDATE meta SET version = ?", SCHEMA_VERSION+1)
	failIfError(t, err, "SQL error: %s", err)

	_, err = FromDbConnection(connection)
	failIfOK(t, err, "opening a future version should fail")
}

func TestFailIfMultipleMetadataEntries(t *testing.T) {
	connection := createConnection(t)
	defer connection.Close()

	createProviderFromConnection(t, connection)

	var err error

	_, err = connection.Exec("INSERT INTO meta (version) VALUES (?)", SCHEMA_VERSION)
	failIfError(t, err, "SQL error: %s", err)

	_, err = FromDbConnection(connection)
	failIfOK(t, err, "opening should fail if the meta table contains multiple entries")
}
