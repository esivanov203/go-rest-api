package sqlstore_test

import (
	"os"
	"testing"
)

var dbUrl string

func TestMain(m *testing.M) {
	dbUrl = os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "evg:strong@tcp(localhost:13306)/test_api"
	}

	os.Exit(m.Run())
}
