package storage

import (
	"testing"

	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/stretchr/testify/require"
)

func TestAuth_toDataset(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		newUser := core.Client{}
		actual, params, err := toDatasetAuth(&newUser).ToSQL()

		expected := `SELECT "id", "email", "role" FROM "users"`

		require.NoError(t, err)
		require.Empty(t, params)
		require.Equal(t, expected, actual)
	})

	t.Run("with email field", func(t *testing.T) {
		newUser := core.Client{
			Email: "test-email",
		}

		actual, _, err := toDatasetAuth(&newUser).ToSQL()
		expected := `SELECT "id", "email", "role" FROM "users" WHERE ("email" = 'test-email')`

		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
}
