package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Email struct {
	Id        uuid.UUID 
	EmailFrom string 
	EmailTo   string 
	Subject   string 
	Html      string 
	FirstName string 
	LastName  string 
	Message   string 
	CreatedAt time.Time 
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

func NewEmail(emailFrom, emailTo, subject, html, firstName, lastName, message string) *Email {
	return &Email{
		Id:        	uuid.New(),
		EmailFrom:  emailFrom,
		EmailTo: 	emailTo,
		Subject:	subject,
		Html:		html,
		FirstName:	firstName,
		LastName:	lastName,
		Message:	message,
		CreatedAt:  time.Now(),
	}
}

