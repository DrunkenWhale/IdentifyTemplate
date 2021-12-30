package model

type User struct {
	UUID     int64 `gorm:"primaryKey"`
	Name     string
	Password string
	Mailbox  string `gorm:"index;unique"`
}

