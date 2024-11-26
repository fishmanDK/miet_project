package storage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetUserOrders_toDataset(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		actual, params, err := toDatasetGetUserOrders(nil).ToSQL()

		expected := `SELECT * FROM "orders"`

		require.NoError(t, err)
		require.Empty(t, params)
		require.Equal(t, expected, actual)
	})
}
