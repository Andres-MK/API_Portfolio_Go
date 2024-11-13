package services

import (
	"github.com/Andres-MK/internal/application/command"
	"github.com/Andres-MK/internal/application/interfaces"
	"github.com/Andres-MK/internal/application/mapper"
	"github.com/Andres-MK/internal/domain/entities"
	"github.com/Andres-MK/internal/domain/repositories"
)

type EmailService struct {
	emailRepository repositories.EmailRepository
	apiEmailRepository repositories.ApiEmailRepository
}

func NewEmailService(emailRepository repositories.EmailRepository, 
	apiEmailRepository repositories.ApiEmailRepository,
	) interfaces.EmailService {
	return &EmailService{
		emailRepository: emailRepository,
		apiEmailRepository: apiEmailRepository,
	}
}

func (s *EmailService) SendEmail(emailCommand *command.CreateEmailCommand) (*command.CreateEmailCommandResult, error) {
	var newEmail = entities.NewEmail(
		emailCommand.EmailFrom,
		emailCommand.EmailTo,
		emailCommand.Subject,
		emailCommand.Html,
		emailCommand.FirstName,
		emailCommand.LastName,
		emailCommand.Message,
	)

	validateEmail, err := entities.NewValidatedEmail(newEmail)
	if err != nil {
		return nil, err
	}

	// _, err = s.emailRepository.Create(validateEmail)
	// if err != nil {
	// 	return nil, err
	// }

	err = s.apiEmailRepository.ResendEmail(validateEmail)
	if err != nil {
		return nil, err
	}

	result := command.CreateEmailCommandResult{
		Result: mapper.NewEmailResultFromValidatedEntity(validateEmail),
	}

	return &result, nil
	
}

