package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/almanalfaruq/alfarpos-backend/util"
)

func TestRead(t *testing.T) {
	expectedResult := Config{Env: "test", Database: Database{
		Host:       "localhost",
		Port:       5432,
		Username:   "postgres",
		Password:   "",
		DBName:     "alfarpos",
		DBTestName: "alfarpos_test",
	}}

	t.Run("Read - Pass", func(t *testing.T) {
		var actualResult Config
		err := actualResult.Read("github.com/almanalfaruq/alfarpos-backend/config.yaml", &actualResult)

		assert.Nil(t, err)
		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("Read - Fail File Error", func(t *testing.T) {
		var actualResult Config
		err := actualResult.Read("", &actualResult)
		assert.NotNil(t, err)
	})

	t.Run("Read - Fail Wrong Content", func(t *testing.T) {
		var actualResult Config
		err := actualResult.Read("github.com/almanalfaruq/alfarpos-backend/test/resources/config.yaml", &actualResult)

		assert.NotNil(t, err)
	})
}
