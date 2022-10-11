package cli_test

import (
	"fmt"
	"testing"

	"github.com/brunoan99/hexagonal-arch/adapters/cli"
	mock_application "github.com/brunoan99/hexagonal-arch/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Product Test"
	price := 10.0
	status := "enabled"
	id := "abc"

	productMock := mock_application.NewMockIProduct(ctrl)
	productMock.EXPECT().GetID().Return(id).AnyTimes()
	productMock.EXPECT().GetName().Return(name).AnyTimes()
	productMock.EXPECT().GetPrice().Return(price).AnyTimes()
	productMock.EXPECT().GetStatus().Return(status).AnyTimes()

	service := mock_application.NewMockIProductService(ctrl)
	service.EXPECT().Create(name, price).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(id).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID: %s with the name %s has been created with the price %f and status %s.",
		id, name, price, status)
	result, err := cli.Run(service, "create", "", name, price, "")
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been enabled.",
		name)
	result, err = cli.Run(service, "enable", id, "", 0, "")
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been disabled.",
		name)
	result, err = cli.Run(service, "disable", id, "", 0, "")
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		id, name, price, status)
	result, err = cli.Run(service, "anything cause in default it do a Get", id, "", 0, "")
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
