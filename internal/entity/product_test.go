package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	t.Run("Should return success when create product", func(t *testing.T) {
		product, err := NewProduct("notebook", 1000.00)
		assert.Nil(t, err)
		assert.NotEmpty(t, product.ID)
		assert.NotEmpty(t, product.Name)
		assert.NotEmpty(t, product.Price)
		assert.NotEmpty(t, product.CreatedAt)
		assert.Equal(t, product.Name, "notebook")
		assert.Equal(t, product.Price, 1000.00)
	})

	t.Run("Should return error when name is empty", func(t *testing.T) {
		_, err := NewProduct("", 1000.00)
		assert.NotNil(t, err)
		assert.Error(t, err, "name is required")
	})

	t.Run("Should return error when price is zero", func(t *testing.T) {
		_, err := NewProduct("notebook", 0)
		assert.NotNil(t, err)
		assert.Error(t, err, "price is required")
	})

	t.Run("Should return error when price is less than zero", func(t *testing.T) {
		_, err := NewProduct("notebook", -10.00)
		assert.NotNil(t, err)
		assert.Error(t, err, "invalid price")
	})

	t.Run("Test product Validate", func(t *testing.T) {
		product, err := NewProduct("Product 1", 35.90)
		assert.Nil(t, err)
		assert.NotNil(t, product)
		assert.Nil(t, product.Validate())
	})

}
