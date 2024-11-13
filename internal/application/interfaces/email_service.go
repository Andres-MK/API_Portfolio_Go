package interfaces

import "github.com/Andres-MK/internal/application/command"

type EmailService interface {
	SendEmail(emailCommand *command.CreateEmailCommand) (*command.CreateEmailCommandResult, error)
}
