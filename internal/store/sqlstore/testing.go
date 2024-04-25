package sqlstore

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

func TestDb(t *testing.T, dbUrl string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
	return db, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := db.ExecContext(
				context.Background(),
				fmt.Sprintf("TRUNCATE %s", strings.Join(tables, ","))); err != nil {
				t.Fatal(err)
			}
		}
		db.Close()
	}
}
