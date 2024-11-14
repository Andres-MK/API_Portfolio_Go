package mapper

import (
	"github.com/Andres-MK/internal/application/common"
	"github.com/Andres-MK/internal/domain/entities"
)

func NewEmailResultFromValidatedEntity(email *entities.ValidatedEmail) *common.EmailResult {
	return NewEmailResultFromEntity(&email.Email)
}

func NewEmailResultFromEntity(email *entities.Email) *common.EmailResult {

	if email == nil {
		return nil
	}
	
	return &common.EmailResult{
		Id:        email.IdEmail,
		EmailFrom: email.EmailFrom,
		EmailTo:   email.EmailTo,
		Subject:   email.Subject,
		Html:      email.Html,
		FirstName: email.FirstName,
		LastName:  email.LastName,
		Message:   email.Message,
		CreatedAt: email.CreatedAt,
		Company:   email.Company,
		ItisRead:  email.ItisRead,
	}
}

