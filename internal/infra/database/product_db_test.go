package database

import (
	"testing"

	"github.com/MatThHeuss/go-rest-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	t.Run("Create Product", func(t *testing.T) {
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

}
