package mapper

import (
	"github.com/Andres-MK/internal/application/common"
	"github.com/Andres-MK/internal/domain/entities"
)

func NewEmailResultFromValidatedEntity(email *entities.ValidatedEmail) *common.EmailResult {
	return NewEmailResultFromEntity(&email.Email)
}

func NewEmailResultFromEntity(email *entities.Email) *common.EmailResult {
	return &common.EmailResult{
		Id:        email.Id,
		EmailFrom: email.EmailFrom,
		EmailTo:   email.EmailTo,
		Subject:   email.Subject,
		Html:      email.Html,
		FirstName: email.FirstName,
		LastName:  email.LastName,
		Message:   email.Message,
		CreatedAt: email.CreatedAt,
	}
}

