package handlers

import (
	"errors"
	"testing"

	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/stretchr/testify/require"
)

func TestSignUp(t *testing.T) {
	t.Run("empty request", func(t *testing.T) {
		te := newTestEnv(t)

		req := &core.Client{}
		te.storageMock.EXPECT().CreateUser(*req).
			Return(0, errors.New("empty new user"))

		expectedErr := errors.New("empty new user")

		actual, err := te.serviceMock.Auth.CreateUser(*req)
		require.EqualError(t, err, expectedErr.Error())
		require.Equal(t, 0, actual)
	})

}
