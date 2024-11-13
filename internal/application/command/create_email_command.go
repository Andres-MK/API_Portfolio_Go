package command

import (
	"github.com/Andres-MK/internal/application/common"
	"github.com/google/uuid"
)

type CreateEmailCommand struct {
	// TODO: Implement idempotency key

	Id        uuid.UUID 
	EmailFrom string 
	EmailTo   string 
	Subject   string 
	Html      string 
	FirstName string 
	LastName  string 
	Message   string 
}

type CreateEmailCommandResult struct {
	Result *common.EmailResult
}
