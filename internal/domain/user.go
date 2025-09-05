package domain

import (
	"time"
)

type Role string

const (
	USER  Role = "USER"
	ADMIN Role = "ADMIN"
)

type User struct {
	ID        uint      `gorm:"unique;primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"size:25;not null;unique" json:"username"`
	Password  string    `gorm:"size:128;not null" json:"-"`
	Firstname string    `gorm:"size:50;" json:"firstname"`
	Lastname  string    `gorm:"size:50;" json:"lastname"`
	Role      Role      `gorm:"size:10;default:'USER'" json:"role"`
	Token     string    `gorm:"-" json:"token"`
	Device    string    `gorm:"type:text" json:"device"`
	Ip        string    `json:"ip"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserList []User

type LoginResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Role      Role   `json:"role"`
	Token     string `json:"token"`
}

type UserRegisterResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Role      Role      `json:"role"`
	Ip        string    `json:"ip"`
	Device    string    `json:"device"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

type UsersWithPaginate struct {
	Users    []User `json:"users"`
	LastPage int    `json:"last_page"`
	Total    int64  `json:"total"`
}

// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)

// 	if err != nil {
// 		return errors.New("error hashing password")
// 	}

// 	u.Password = string(hash)
// 	return
// }
