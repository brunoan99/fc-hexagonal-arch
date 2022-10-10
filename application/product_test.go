package application_test

import (
	"testing"

	"github.com/brunoan99/hexagonal-arch/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.ID = "123"
	product.Name = "Product 1"
	product.Status = application.ENABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.ID = "123"
	product.Name = "Product 1"
	product.Status = application.ENABLED
	product.Price = 10

	err := product.Disable()
	require.Equal(t, "the price must be zero to disable the product", err.Error())

	product.Price = 0
	err = product.Disable()
	require.Nil(t, err)
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Status = application.ENABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "a"
	_, err = product.IsValid()
	require.Equal(t, "status must be ENABLED or DISABLED", err.Error())
	product.Status = application.ENABLED

	product.Price = -1
	_, err = product.IsValid()
	require.Equal(t, "price must be greater or equal zero", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	product := application.Product{}
	id := uuid.NewV4().String()
	product.ID = id

	recivedID := product.GetID()
	require.Equal(t, id, recivedID)
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	name := "Product Name Example"
	product.Name = name

	recivedName := product.GetName()
	require.Equal(t, name, recivedName)
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	status := application.ENABLED
	product.Status = status

	recivedStatus := product.GetStatus()
	require.Equal(t, status, recivedStatus)
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	price := 10.0
	product.Price = price

	recivedPrice := product.GetPrice()
	require.Equal(t, price, recivedPrice)
}
