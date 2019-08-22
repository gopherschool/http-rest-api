package store_test

import (
	"testing"

	"github.com/gopherschool/http-rest-api/internal/app/model"
	"github.com/gopherschool/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: email,
	})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
