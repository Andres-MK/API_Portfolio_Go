package mapper

import (
	"github.com/Andres-MK/internal/api/dto/response"
	"github.com/Andres-MK/internal/application/common"
)

func ToEmailResponse(email *common.EmailResult) *response.EmailResponse {
	return &response.EmailResponse{
		Id:        email.Id.String(),
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

