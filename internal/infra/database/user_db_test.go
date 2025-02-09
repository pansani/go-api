package database

import (
	"testing"

	internalEntity "github.com/pansani/go-api/internal/entity"
	pkgEntity "github.com/pansani/go-api/pkg/entity"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&internalEntity.User{})
	user, _ := internalEntity.NewUser("pansani@gmail.com", "pansani", "123456")
	userDB := NewUser(db)

	// Explicitly setting the ID before saving to ensure GORM does not overwrite it.
	user.ID = pkgEntity.NewID()

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFromDB internalEntity.User
	err = db.First(&userFromDB, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFromDB.ID)
	assert.Equal(t, user.Email, userFromDB.Email)
	assert.Equal(t, user.Name, userFromDB.Name)
	assert.NotNil(t, user.Password)
}

func TestFindUserByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&internalEntity.User{})
	user, _ := internalEntity.NewUser("pansani@gmail.com", "pansani", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	userFromDB, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFromDB.ID)
	assert.Equal(t, user.Email, userFromDB.Email)
	assert.Equal(t, user.Name, userFromDB.Name)
	assert.NotNil(t, user.Password)
}
