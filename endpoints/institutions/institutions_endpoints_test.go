package institutions_test

import (
	"context"
	"testing"

	"github.com/weportfolio/go-nordigen/endpoints/tests"

	"github.com/stretchr/testify/assert"
	"github.com/weportfolio/go-nordigen/consts"
)

func TestClient_List(t *testing.T) {
	t.Parallel()

	t.Run("list institutions", func(t *testing.T) {
		t.Parallel()

		client := tests.GetTestClient(t)
		assert.NotNil(t, client)

		token, err := client.Token().New(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		institutions, err := client.Institutions().List(context.Background(), token.Access, consts.NetherlandsInstitution, true)
		assert.NoError(t, err)
		assert.NotNil(t, institutions)
	})

	t.Run("list institutions with invalid token", func(t *testing.T) {
		t.Parallel()

		client := tests.GetTestClient(t)
		assert.NotNil(t, client)

		institutions, err := client.Institutions().List(context.Background(), "invalid", consts.NetherlandsInstitution, true)
		assert.Error(t, err)
		assert.Nil(t, institutions)

		checkErr := consts.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}
