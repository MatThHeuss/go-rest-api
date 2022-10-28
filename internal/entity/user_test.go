package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Matheus Santos", "test@gmail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Matheus Santos", user.Name)
	assert.Equal(t, "test@gmail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Matheus Santos", "test@gmail.com", "123456")
	assert.Nil(t, err)

	assert.NotEqual(t, "123456", user.Password)

	err = user.ValidatePassword("123456")
	assert.Nil(t, err)

	err = user.ValidatePassword("1234567")
	assert.NotNil(t, err)

}

func TestUser_ValidateEmail(t *testing.T) {
	user, _ := NewUser("Matheus Santos", "test@gmail.com", "123456")

	assert.False(t, user.ValidateEmail("2234.com"))

	assert.True(t, user.ValidateEmail("test@gmail.com"))
}
