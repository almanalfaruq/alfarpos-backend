package util_test

import (
	"testing"

	. "../util"
)

func TestRead(t *testing.T) {
	configTest := Config{Env: "test", Database: Database{
		Host:       "localhost",
		Port:       5432,
		Username:   "postgres",
		Password:   "",
		DBName:     "alfarpos",
		DBTestName: "alfarpos_test",
	}}

	var config Config
	err := config.Read(&config)

	if err != nil {
		t.Error("\nExpected: not nil\nGot: nil")
	}
	if config != configTest {
		t.Errorf("\nExpected: %v\nGot: %v", configTest, config)
	}
}
