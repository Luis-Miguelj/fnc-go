package infra

import "time"

type User struct {
	Id        string    `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
	Admin     bool
}

type Finance struct {
	Id         string    `gorm:"primaryKey"`
	UserId     string    `gorm:"not null"`
	Type       string    `gorm:"not null"`
	Value      float64   `gorm:"not null"`
	CategoryId string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null"`
	UpdatedAt  time.Time
}

type Category struct {
	Id        string    `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	UserId    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
}
