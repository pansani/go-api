package database

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pansani/go-api/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Macbook", 2000.00)
	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	var productFromDB entity.Product
	err = db.First(&productFromDB, "id = ?", product.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFromDB.ID)
	assert.Equal(t, product.Name, productFromDB.Name)
	assert.Equal(t, product.Price, productFromDB.Price)
}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Macbook", 2000.00)
	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	productFromDB, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFromDB.ID)
	assert.Equal(t, product.Name, productFromDB.Name)
	assert.Equal(t, product.Price, productFromDB.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Macbook", 2000.00)
	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	product.Name = "Linux"
	product.Price = 1000.00

	err = productDB.Update(product)
	assert.Nil(t, err)

	var productFromDB entity.Product
	err = db.First(&productFromDB, "id = ?", product.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFromDB.ID)
	assert.Equal(t, product.Name, productFromDB.Name)
	assert.Equal(t, product.Price, productFromDB.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Macbook", 2000.00)
	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	err = productDB.Delete(product.ID.String())
	assert.Nil(t, err)

	var productFromDB entity.Product
	err = db.First(&productFromDB, "id = ?", product.ID).Error

	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err, "expected record to be deleted")
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&entity.Product{})
	db.Exec("DELETE FROM products")

	for i := 1; i <= 23; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		err = db.Create(product).Error
		assert.NoError(t, err)
	}

	productDB := NewProduct(db)

	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)

	for _, product := range products {
		assert.NotEmpty(t, product.ID)
		assert.NotEmpty(t, product.Name)
		assert.Greater(t, product.Price, 0.0)
	}
}
