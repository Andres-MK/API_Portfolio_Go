package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Email struct {
	IdEmail   uuid.UUID `gorm:"column:IdEmail"`
	EmailFrom string    `gorm:"column:EmailFrom"`
	EmailTo   string    `gorm:"column:EmailTo"`
	Subject   string 	`gorm:"column:Subject"`
	Html      string 	`gorm:"column:Html"`
	FirstName string 	`gorm:"column:FirstName"`
	LastName  string 	`gorm:"column:LastName"`
	Message   string 	`gorm:"column:Message"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	Company   string	`gorm:"column:Company"`
	ItisRead  bool 		`gorm:"column:ItisRead"`
}

func (p *Email) validate() error {
	if p.EmailFrom == "" {
		return errors.New("El email no debe ir vacio")
	}
	if p.Subject == "" {
		return errors.New("El correo debe tener un asunto")
	}
	if p.FirstName == "" {
		return errors.New("Debe llenar su nombre")
	}
	if p.LastName == "" {
		return errors.New("Debe llenar su apellido")
	}
	if p.Message == "" {
		return errors.New("El correo debe tener su mensaje")
	}

	return nil
}

func NewEmail(emailFrom, emailTo, subject, html, firstName, lastName, message, company string) *Email {
	return &Email{
		IdEmail:    uuid.New(),
		EmailFrom:  emailFrom,
		EmailTo: 	emailTo,
		Subject:	subject,
		Html:		html,
		FirstName:	firstName,
		LastName:	lastName,
		Message:	message,
		Company:    company,
	}
}

