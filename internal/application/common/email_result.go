package common

import (
	"time"

	"github.com/google/uuid"
)

type EmailResult struct {
	Id        		uuid.UUID 
	EmailFrom 		string 
	EmailTo   		string 
	Subject   		string 
	Html      		string 
	FirstName 		string 
	LastName  		string 
	Message   		string 
	CreatedAt 		time.Time 
	Msg_Response  	string 
}