package interfaces

import (
	"github.com/Andres-MK/internal/application/command"
	"github.com/Andres-MK/internal/application/query"
)

type EmailService interface {
	SendEmail(emailCommand *command.CreateEmailCommand) (*command.CreateEmailCommandResult, error)
	
	GetAllEmail() (*query.ListEmailsQueryResult, error)
}
