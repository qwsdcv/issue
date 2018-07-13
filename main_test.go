package main_test

import (
	"issues/db"
	"testing"
)

func Test_main(t *testing.T) {
	dsn := db.ConnectString
	t.Logf("%s", dsn)
}
