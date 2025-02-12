package database

import (
	"testing"

	"github.com/pansani/go-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("pansani@gmail.com", "pansani", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFromDB entity.User
	err = db.First(&userFromDB, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.Email, userFromDB.Email)
	assert.Equal(t, user.Name, userFromDB.Name)
	assert.NotNil(t, userFromDB.Password)
}

func TestFindUserByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("pansani@gmail.com", "pansani", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	userFromDB, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, userFromDB.Email)
	assert.Equal(t, user.Name, userFromDB.Name)
	assert.NotNil(t, userFromDB.Password)
}
