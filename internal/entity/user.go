package entity

import (
	"errors"
	"net/mail"

	"github.com/MatThHeuss/go-rest-api/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailIsRequired    = errors.New("email is required")
	ErrPasswordIsRequired = errors.New("password is required")
	ErrInvalidEmail       = errors.New("invalid email")
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}

	err = user.Validate()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}

func (u *User) Validate() error {
	if u.ID.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := entity.ParseID(u.ID.String()); err != nil {
		return ErrInvalidID
	}

	if u.Name == "" {
		return ErrNameIsRequired
	}

	if u.Email == "" {
		return ErrEmailIsRequired
	}

	if err := u.ValidateEmail(u.Email); !err {
		return ErrInvalidEmail

	}

	if u.Password == "" {
		return ErrPasswordIsRequired
	}

	return nil
}
