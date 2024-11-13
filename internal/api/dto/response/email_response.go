package response

import "time"

type EmailResponse struct {
	Id        string    `json:"id"`
	EmailFrom string    `json:"email_from"`
	EmailTo   string    `json:"email_to"`
	Subject   string    `json:"subject"`
	Html      string    `json:"html"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type ListEmailsResponse struct {
	Emails []*EmailResponse `json:"Emails"`
}