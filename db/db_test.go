package db_test

import (
	"testing"

	"github.com/boring-school-work/irb-tracker/db"
)

func TestDatabaseConnection(t *testing.T) {
	conn, err := db.Init()
	if err != nil {
		t.Fatalf("Could not connect to database: %s", err)
	}

	defer conn.Close()

	if err := conn.Ping(); err != nil {
		t.Fatalf("Ping Error: %s", err)
	}

	t.Log("Database connected!")
}
