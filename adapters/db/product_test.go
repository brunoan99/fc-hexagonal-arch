package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/brunoan99/hexagonal-arch/adapters/db"
	"github.com/brunoan99/hexagonal-arch/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
		);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abc","Product Test",0,"disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	sut := db.NewProductDb(Db)

	product, err := sut.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "abc", product.GetID())
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	sut := db.NewProductDb(Db)
	newProduct := application.NewProduct()
	newProduct.Name = "Product Test 2"
	newProduct.Price = 25

	product, err := sut.Save(newProduct)
	require.Nil(t, err)
	require.Equal(t, "Product Test 2", product.GetName())
	require.Equal(t, 25.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())

	newProduct.Status = "enabled"
	product, err = sut.Save(newProduct)
	require.Nil(t, err)
	require.Equal(t, "enabled", product.GetStatus())
}
