package model

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}
