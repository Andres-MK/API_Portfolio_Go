package entities

type ValidatedEmail struct {
	Email
	isValidated bool
}

func (vp *ValidatedEmail) IsValid() bool {
	return vp.isValidated
}

func NewValidatedEmail(email *Email) (*ValidatedEmail, error) {
	if err := email.validate(); err != nil {
		return nil, err
	}

	return &ValidatedEmail{
		Email:       *email,
		isValidated: true,
	}, nil
}