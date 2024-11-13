package api

import (
	"fmt"

	"github.com/Andres-MK/internal/domain/entities"
	"github.com/Andres-MK/internal/domain/repositories"
	"github.com/resend/resend-go/v2"
)

type ApiEmailRepository struct {
}

func NewApiEmailRepository() repositories.ApiEmailRepository {
	return &ApiEmailRepository{}
}

func (api *ApiEmailRepository) ResendEmail(email *entities.ValidatedEmail) (error) {
	paramsEmail := &entities.Email{
		EmailFrom: email.EmailFrom,
		EmailTo:   email.EmailTo,
		Subject:   email.Subject,
		Html:      email.Html,
	}

	apiKey := "re_Wwz9HJej_14twm9cr5yWsaRz8AHdmAuhh"

    client := resend.NewClient(apiKey)

    params := &resend.SendEmailRequest{
        From:    paramsEmail.EmailFrom,
        To:      []string{paramsEmail.EmailTo},
        Html:    paramsEmail.Html,
        Subject: paramsEmail.Subject,
        Cc:      []string{""},
        Bcc:     []string{""},
        ReplyTo: "",
    }
    
    response, err := client.Emails.Send(params)
    fmt.Println(response)
    if err != nil {
        return err
    }

	return err

}