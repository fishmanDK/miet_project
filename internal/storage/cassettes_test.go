package storage

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListCassttes_toDataset(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		actual, params, err := toDatasetAuth(nil).ToSQL()

		expected := `SELECT * FROM "cassettes"`

		require.NoError(t, err)
		require.Empty(t, params)
		require.Equal(t, expected, actual)
	})

	t.Run("with all filters", func(t *testing.T) {
		id := 1

		actual, _, err := toDatasetListCassettes(&id).ToSQL()
		expected := `SELECT "cassettes"."id", "cassettes"."name", "cassettes"."genre", "cassettes"."year_of_release" FROM "cassettes" INNER JOIN "cassetteavailability" ON ("cassetteavailability"."cassette_id" = "cassettes"."id") WHERE ("cassetteavailability"."store_id" = 1)`

		// Удаляем лишние пробелы в обеих строках
		normalizedExpected := strings.ReplaceAll(strings.TrimSpace(expected), "\n", "")
		normalizedActual := strings.ReplaceAll(strings.TrimSpace(actual), "\n", "")

		require.NoError(t, err)
		require.Equal(t, normalizedExpected, normalizedActual)
	})
}
