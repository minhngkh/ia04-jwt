package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"test-echo/internal/db"
)

type NewProfileInfo struct {
	Email    string
	Password string
}

type LoginInfo struct {
	Email    string
	Password string
}

var (
	ErrUserExists        = errors.New("User already exists")
	ErrUserNotFound      = errors.New("User not found")
	ErrIncorrectPassword = errors.New("Incorrect password")
)

func CreateUser(email string, password string) error {
	// Hash password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := db.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	// Create user in database
	result := db.Get().Create(&user)
	if result.RowsAffected == 0 {
		return ErrUserExists
	}

	return result.Error
}

func VerifyLoginInfo(info *LoginInfo) error {
	var user db.User
	result := db.Get().Where("email = ?", info.Email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ErrUserExists
		}
		return result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(info.Password))
	if err != nil {
		return ErrIncorrectPassword
	}

	return nil
}
