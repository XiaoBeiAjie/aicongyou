package models

import (
	"time"
	"github.com/XiaoBeiAjie/aicongyou/db"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
    Name      string         `gorm:"size:100;not null" json:"name"`
    Email     string         `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Password  string 		 `gorm:"size:255;not null" json:"password"`
	Role      UserRole       `gorm:"type:varchar(20);not null;default:'student'" json:"role"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *User) Create() error {
    db := db.GetDB()
    return db.Create(u).Error
}

func (u *User) GetByID(id uint) error {
    db := db.GetDB()
    return db.First(u, id).Error
}

func GetAllUsers() ([]User, error) {
    var users []User
	db := db.GetDB()
    err := db.Find(&users).Error
    return users, err
}
