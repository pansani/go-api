package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("product", 100)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "product", product.Name)
	assert.Equal(t, 100, product.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 100)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("product", 0)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("product", -100)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsInvalid, err)
}

func TestProductValidate(t *testing.T) {
	product, err := NewProduct("product", 100)
	assert.Nil(t, err)
	assert.Nil(t, product.Validate())
}
