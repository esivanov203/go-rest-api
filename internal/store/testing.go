package store

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, dbUrl string) (*Store, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.Url = dbUrl
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.ExecContext(
				context.Background(),
				fmt.Sprintf("TRUNCATE %s", strings.Join(tables, ","))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}
