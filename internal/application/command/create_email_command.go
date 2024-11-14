package command

import (
	"github.com/Andres-MK/internal/application/common"
)

type CreateEmailCommand struct {
	// TODO: Implement idempotency key

	EmailFrom string 
	EmailTo   string 
	Subject   string 
	Html      string 
	FirstName string 
	LastName  string 
	Message   string
	Company   string 
}

type CreateEmailCommandResult struct {
	Result *common.EmailResult
}
