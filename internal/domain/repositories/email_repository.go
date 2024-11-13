package repositories

import "github.com/Andres-MK/internal/domain/entities"

type EmailRepository interface {
	Create(email *entities.ValidatedEmail) (*entities.Email, error)
	
}

type ApiEmailRepository interface {
	ResendEmail(email *entities.ValidatedEmail) (error)

}