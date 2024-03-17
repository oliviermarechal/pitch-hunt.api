package models

import (
	"encoding/json"
	"time"
)

func CreateUser(
	id string,
	email string,
	password string,
) *User {
	return &User{Id: id, Email: email, Password: password}
}

type User struct {
	Id        string    `gorm:"type:uuid;primary_key" json:"id"`
	Username  string    `json:"username"`
	Email     string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}
