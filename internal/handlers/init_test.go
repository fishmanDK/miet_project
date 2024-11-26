package handlers

import (
	"context"
	"testing"

	"github.com/fishmanDK/miet_project/internal/service"
	"github.com/fishmanDK/miet_project/internal/storage"
	"github.com/golang/mock/gomock"
)

type testEnvAuth struct {
	ctx  context.Context
	ctrl *gomock.Controller

	storageMock *storage.MockAuth

	serviceMock *service.Service
}

func (te *testEnvAuth) initMocks() {
	//te.storageMock = &storage.Storage{
	//	Auth: storage.NewMockAuth(te.ctrl),
	//	Cassettes: storage.NewMockCassettes(te.ctrl),
	//	Store: storage.NewMockStore(te.ctrl),
	//	Reservation : storage.NewMockReservation(te.ctrl),
	//	Orders: storage.NewMockOrders(te.ctrl),
	//}

	te.storageMock = storage.NewMockAuth(te.ctrl)
}

func newTestEnv(t *testing.T) *testEnvAuth {
	te := &testEnvAuth{
		ctx:  context.Background(),
		ctrl: gomock.NewController(t),
	}

	te.initMocks()
	te.serviceMock = &service.Service{
		Auth: service.NewMockAuth(te.ctrl),
		//Cassettes: service.NewMockCassettes(te.ctrl),
		//Store: service.NewMockStore(te.ctrl),
		//Reservation : service.NewMockReservation(te.ctrl),
		//Orders: service.NewMockOrders(te.ctrl),
	}

	return te
}
