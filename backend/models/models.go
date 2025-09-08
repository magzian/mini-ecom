package models

import (
	"time"
)

type User struct {
	ID          uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	Username    string        `json:"username" gorm:"unique;not null"`
	Password    string        `json:"password" gorm:"not null"`
	Permissions []Permissions `json:"permissions" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
}

type Permissions struct {
	ID        uint `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint `json:"user_id" gorm:"not null;index"`
	Entry     int  `json:"entry"`
	AddFlag   bool `json:"add_flag" gorm:"default:false"`
	AdminFlag bool `json:"admin_flag" gorm:"default:false"`
	User      User `json:"-" gorm:"foreignKey:UserID"`
}

type Product struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Price       float64   `json:"price" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type Order struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
