package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Bruno", "bruno@mail.com", "123456")
	assert.Nil(t, err)
	assert.NotNilf(t, user, "User should not be nil")
	assert.NotEmptyf(t, user.ID, "User ID should not be empty")
	assert.NotEmptyf(t, user.Password, "User Password should not be empty")
	assert.Equal(t, "Bruno", user.Name)
	assert.Equal(t, "bruno@mail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Bruno", "bruno@mail.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, "123456", user.Password) // Password should be hashed
}
