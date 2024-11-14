package db

import (
	"time"

	"github.com/google/uuid"
)

type EmailDatabase struct {
	IdEmail   uuid.UUID `gorm:"column:IdEmail;primaryKey"`
	EmailFrom string 	`gorm:"column:EmailFrom"`
	EmailTo   string 	`gorm:"column:EmailTo"`
	Subject   string 	`gorm:"column:Subject"`
	Html      string 	`gorm:"column:Html"`
	FirstName string 	`gorm:"column:FirstName"`
	LastName  string 	`gorm:"column:LastName"`
	Message   string 	`gorm:"column:Message"`
	Company   string 	`gorm:"column:Company"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	ItisRead  bool     	`gorm:"column:ItisRead"`
}

func (EmailDatabase) TableName() string {
	return "EmailsRegister"
}