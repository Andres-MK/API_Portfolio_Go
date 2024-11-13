package request

import "github.com/Andres-MK/internal/application/command"

type CreateEmailRequest struct {
	EmailFrom string `json:"email_from"`
	EmailTo   string `json:"email_to"`
	Subject   string `json:"subjet"`
	Html      string `json:"html"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Message   string `json:"message"`
}

func (req *CreateEmailRequest) ToCreateEmailCommand() (*command.CreateEmailCommand, error) {

	return &command.CreateEmailCommand{
		EmailFrom: req.EmailFrom,
		EmailTo:   req.EmailTo,
		Subject:   req.Subject,
		Html:      req.Html,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Message:   req.Message,
	}, nil
}