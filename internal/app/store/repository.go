package store

import "github.com/gopherschool/http-rest-api/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
