package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Username string `json:"username" gorm:"type:citext;unique;not null"`
	Password string `json:"-" gorm:"not null"`
	Role     string `json:"role" gorm:"default:admin"` // admin, super_admin, content_editor, moderator
	IsActive bool   `json:"is_active" gorm:"default:true"`
}

// HashPassword hashes the user's password
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword compares the provided password with the hashed password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// BeforeCreate hook to hash password before creating user
func (u *User) BeforeCreate(tx *gorm.DB) error {
	return u.HashPassword()
}
