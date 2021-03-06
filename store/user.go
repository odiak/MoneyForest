package store

import (
	"fmt"
	"github.com/go-pg/pg/orm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                string `sql:"type:uuid"`
	Email             string `sql:",notnull"`
	Name              string `sql:",notnull"`
	EncryptedPassword string `sql:",notnull"`
}

func (u User) String() string {
	return fmt.Sprintf("<User ID=%s Email=%s Name=%s EncryptedPassword=%s>",
		u.ID, u.Email, u.Name, u.EncryptedPassword)
}

func (u *User) SetPassword(pw string) error {
	ep, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.EncryptedPassword = string(ep)
	return nil
}

func (u User) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pw)) == nil
}

func (u *User) validate() error {
	if u.Name == "" {
		return ValidationError("name is required")
	}
	if u.Email == "" {
		return ValidationError("email is required")
	}
	if u.EncryptedPassword == "" {
		return ValidationError("password is required")
	}
	return nil
}

func (u *User) BeforeInsert(db orm.DB) error {
	return u.validate()
}

func (u *User) BeforeUpdate(db orm.DB) error {
	return u.validate()
}
