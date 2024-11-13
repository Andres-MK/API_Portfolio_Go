package db

import (
	"github.com/Andres-MK/internal/domain/entities"
	"github.com/Andres-MK/internal/domain/repositories"
)

type EmailRepository struct {
	
}

func NewEmailRepository() repositories.EmailRepository {
	return &EmailRepository{}
}

func (repo *EmailRepository) Create(email *entities.ValidatedEmail) (*entities.Email, error) {
	
	response := &entities.Email{
		Id:        email.Id,
		EmailFrom: email.EmailFrom,
		EmailTo:   email.EmailTo,
		Subject:   email.Subject,
		Html:      email.Html,
	}




	
	return response, nil
}
