package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/MatThHeuss/go-rest-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestProduct(t *testing.T) {
	t.Run("Create Product", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}
		db.AutoMigrate(&entity.Product{})
		product, _ := entity.NewProduct("notebook", 1500.00)
		productDb := NewProduct(db)

		err = productDb.Create(product)
		assert.Nil(t, err)

		var productFound entity.Product
		err = db.First(&productFound, "id = ?", product.ID).Error
		assert.Nil(t, err)

		assert.Equal(t, productFound.Name, product.Name)
		assert.Equal(t, productFound.Price, product.Price)
		assert.Equal(t, productFound.ID, product.ID)
		assert.Equal(t, productFound.CreatedAt.Local(), product.CreatedAt.Local())
	})

	t.Run("Find by id", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}
		db.AutoMigrate(&entity.Product{})
		product, _ := entity.NewProduct("notebook", 1500.00)
		productDb := NewProduct(db)
		err = productDb.Create(product)
		assert.Nil(t, err)

		productById, err := productDb.FindByID(product.ID.String())

		assert.Nil(t, err)
		assert.Equal(t, productById.ID, product.ID)
		assert.Equal(t, productById.Name, product.Name)
		assert.Equal(t, productById.Price, product.Price)
		assert.Equal(t, productById.CreatedAt.Local(), product.CreatedAt.Local())

	})

	t.Run("Find all products", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}
		db.AutoMigrate(&entity.Product{})
		for i := 1; i < 24; i++ {
			product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
			assert.NoError(t, err)
			db.Create(product)
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
	})

	t.Run("Update product", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}
		db.AutoMigrate(&entity.Product{})
		product, err := entity.NewProduct("Product 1", 10.00)
		assert.NoError(t, err)
		db.Create(product)
		productDB := NewProduct(db)
		product.Name = "Product 2"
		err = productDB.Update(product)
		assert.NoError(t, err)
		product, err = productDB.FindByID(product.ID.String())
		assert.NoError(t, err)
		assert.Equal(t, "Product 2", product.Name)

	})

	t.Run("Delete product", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}
		db.AutoMigrate(&entity.Product{})
		product, err := entity.NewProduct("Product 1", 10.00)
		assert.NoError(t, err)
		db.Create(product)
		productDB := NewProduct(db)

		err = productDB.Delete(product.ID.String())
		assert.NoError(t, err)

		_, err = productDB.FindByID(product.ID.String())
		assert.Error(t, err)
	})

}
