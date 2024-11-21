package entity

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

var (
	ErrEmptyName  = errors.New("an user can't have an empty name")
	ErrFormatName = errors.New("an user can't have a name with special characters")

	nameRegex = regexp.MustCompile(`^[a-zA-Z]+$`)
)

type User struct {
	ID   uuid.UUID
	Name string
}

func newClient(name string) (User, error) {
	if name == "" {
		return User{}, ErrEmptyName
	}
	if !nameRegex.MatchString(name) {
		return User{}, ErrFormatName
	}
	return User{
		ID:   uuid.New(),
		Name: name,
	}, nil
}
