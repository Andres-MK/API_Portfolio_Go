package mapper

import (
	"fmt"

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
		Company:   email.Company,
		ItisRead:  email.ItisRead,
	}
}

func ToEmailListResponse(emails []*common.EmailResult) *response.ListEmailsResponse {
	var responseList []*response.EmailResponse
	for _, email := range emails {
		fmt.Println(email)
		responseList = append(responseList, ToEmailResponse(email))
	}
	return &response.ListEmailsResponse{Emails: responseList}
}


