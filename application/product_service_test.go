package application_test

import (
	"testing"

	"github.com/brunoan99/hexagonal-arch/application"
	mock_application "github.com/brunoan99/hexagonal-arch/application/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockIProduct(ctrl)
	persistence := mock_application.NewMockIProductPersistence(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	sut := application.ProductService{Persistence: persistence}
	result, err := sut.Get("any_id")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockIProduct(ctrl)
	persistence := mock_application.NewMockIProductPersistence(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	sut := application.ProductService{Persistence: persistence}
	result, err := sut.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockIProduct(ctrl)
	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)
	persistence := mock_application.NewMockIProductPersistence(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	sut := application.ProductService{Persistence: persistence}
	result, err := sut.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = sut.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
