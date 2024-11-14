package db

import (
	"github.com/Andres-MK/internal/domain/entities"
)

func toDBEmail(email *entities.ValidatedEmail) *EmailDatabase {
	var e = &EmailDatabase{
		EmailFrom: email.EmailFrom,
		EmailTo:   email.EmailTo,
		Subject:   email.Subject,
		Html:      email.Html,
		Message:   email.Message,
		FirstName: email.FirstName,
		LastName:  email.LastName,
		Company:   email.Company,	
	}
	
	return e
}

func fromDBEmail(email *EmailDatabase) *entities.Email{
	var e = &entities.Email{
		IdEmail: 	email.IdEmail,
		EmailFrom: 	email.EmailFrom,
		EmailTo  : 	email.EmailTo, 
		Subject   : email.Subject,
		Html      : email.Html,
		FirstName : email.FirstName,
		LastName  : email.LastName,
		Message   : email.Message,
		CreatedAt : email.CreatedAt,
		Company   : email.Company,
	}
	return e
}
