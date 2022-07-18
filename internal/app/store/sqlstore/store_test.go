package sqlstore_test

import (
	"os"
	"testing"
)

var databaseURL string

func TestMain(m *testing.M) {
	databaseURL = "host=localhost dbname=todoapi_test password=123 sslmode=disable"
	os.Exit(m.Run())
}
