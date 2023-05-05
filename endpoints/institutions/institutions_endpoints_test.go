package institutions_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/weportfolio/go-nordigen"
	"github.com/weportfolio/go-nordigen/consts"
	"testing"
)

func TestClient_List(t *testing.T) {
	t.Parallel()

	t.Run("list institutions", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			consts.GetSecrets(t),
		)
		assert.NotNil(t, client)

		token, err := client.Token().New(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		institutions, err := client.Institutions().List(context.Background(), token.Access, consts.NetherlandsInstitution, true)
		assert.NoError(t, err)
		assert.NotNil(t, institutions)
	})
}
