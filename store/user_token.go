package store

import (
	"time"
)

type UserToken struct {
	Token     string `sql:",notnull"`
	UserID    string `sql:"type:uuid,notnull"`
	User      *User
	CreatedAt time.Time `sql:",notnull"`
}
