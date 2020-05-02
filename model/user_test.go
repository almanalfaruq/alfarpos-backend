package model_test

import (
	"testing"

	. "github.com/almanalfaruq/alfarpos-backend/model"
	"golang.org/x/crypto/bcrypt"
)

func TestUserModel(t *testing.T) {
	t.Run("Creating user model", func(t *testing.T) {
		username := "almanalfaruq"
		password := "adminalman"
		fullname := "Almantera"
		address := "Boyolali"
		phone := "081225812599"
		role := RoleAdmin

		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			t.Fatal("Bcrypt error")
		}

		user := User{Template: Template{ID: 1}, Username: username, Password: string(encryptedPassword), FullName: fullname, Address: address, Phone: phone, RoleID: RoleAdmin}

		if user.ID != 1 {
			FatalMessage(t, 1, user.ID)
		}
		if user.Username != username {
			FatalMessage(t, username, user.Username)
		}
		if bcrypt.CompareHashAndPassword(encryptedPassword, []byte(password)) != nil {
			FatalMessage(t, true, false)
		}
		if user.FullName != fullname {
			FatalMessage(t, fullname, user.FullName)
		}
		if user.Address != address {
			FatalMessage(t, address, user.Address)
		}
		if user.Phone != phone {
			FatalMessage(t, phone, user.Phone)
		}
		if user.RoleID != role {
			FatalMessage(t, role, user.RoleID)
		}
	})
}

func FatalMessage(t *testing.T, args ...interface{}) {
	t.Helper()
	t.Fatalf("Expected %v but got %v", args[0], args[1])
}
