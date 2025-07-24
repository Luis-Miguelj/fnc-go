package infra

type User struct {
	Id        string `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	CreatedAt string
}

type Finance struct {
	Id         string  `gorm:"primaryKey"`
	UserId     string  `gorm:"not null"`
	Type       string  `gorm:"not null"`
	Value      float64 `gorm:"not null"`
	CategoryId string  `gorm:"not null"`
	CreatedAt  string  `gorm:"not null"`
}

type Category struct {
	Id        string `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	UserId    string `gorm:"not null"`
	CreatedAt string `gorm:"not null"`
}
