package bug

import (
	"context"
	stdsql "database/sql"
	"fmt"
	"os"
	"testing"

	"entgo.io/bug/internal/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	_ "entgo.io/bug/internal/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func createClient(t *testing.T, db *stdsql.DB) *ent.Client {
	return ent.NewClient(ent.Driver(sql.OpenDB(dialect.SQLite, db)))
}

func ensureDatabaseRemoved(t *testing.T, path string) {
	_, err := os.Stat(path)
	switch {
	case os.IsNotExist(err):
		// Do nothing, the file did not exist
	case err != nil:
		t.Fatal(err)
	default:
		t.Logf("Deleting the existing %q file", path)

		if err := os.Remove(path); err != nil {
			t.Logf("Failed to remove %q: %s", path, err.Error())
		}
	}
}

func Test_CGOSQLite(t *testing.T) {
	ctx := context.Background()
	path := "cgo.sqlite.db"

	// Ensure we remove the database file (if it exists)
	ensureDatabaseRemoved(t, path)

	// Create a sql.DB instance
	db, err := stdsql.Open("sqlite3", fmt.Sprintf("file:%s?_fk=1", path))
	if err != nil {
		t.Fatal(err)
	}

	client := createClient(t, db)
	defer client.Close()

	// First, run Schema.Create() once (first time, for example)
	//
	// This will create the sqlite.db file in the directory root.
	if err := client.Schema.Create(ctx); err != nil {
		t.Fatal(err)
	}

	// Simulate a second run, for my usecase this only happens on the second (or subsequent) startup. It passes here.
	t.Log("Simulating a startup after the database has already been migrated by calling Schema.Create() again.")

	if err := client.Schema.Create(ctx); err != nil {
		t.Fatal(err.Error())
	}
}

func Test_NonCGOSQLite(t *testing.T) {
	ctx := context.Background()
	path := "noncgo.sqlite.db"

	// Ensure we remove the database file (if it exists)
	ensureDatabaseRemoved(t, path)

	// Create a sql.DB instance
	// Using a custom SQLite driver that enables foreign key constraints: https://github.com/ent/ent/discussions/1667#discussioncomment-1132296
	db, err := stdsql.Open("sqliter", fmt.Sprintf("file:%s?_fk=1", path))
	if err != nil {
		t.Fatal(err)
	}

	client := createClient(t, db)
	defer client.Close()

	// First, run Schema.Create() once (first time, for example)
	//
	// This will create the sqlite.db file in the directory root.
	if err := client.Schema.Create(ctx); err != nil {
		t.Fatal(err)
	}

	// Simulate a second run, for my usecase this only happens on the second (or subsequent) startup. It fails here.
	t.Log("Simulating a startup after the database has already been migrated by calling Schema.Create() again.")

	if err := client.Schema.Create(ctx); err != nil {
		t.Fatalf("Failed to run Schema.Create: %s", err.Error())
	}
}
