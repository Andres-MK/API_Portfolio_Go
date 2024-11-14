package api

import (
	"encoding/json"
	"net/http"

	"github.com/Andres-MK/internal/api/dto/mapper"
	"github.com/Andres-MK/internal/api/dto/request"
	"github.com/Andres-MK/internal/api/dto/response"
	"github.com/Andres-MK/internal/application/interfaces"
	"github.com/gorilla/mux"
)

type EmailController struct {
	service interfaces.EmailService
}

func NewEmailController(r *mux.Router, services interfaces.EmailService) *EmailController {
	controller := &EmailController{
		service: services,
	}

	r.HandleFunc("/email", controller.CreateEmailController).Methods("POST")	
	

	return controller
	
}

func (ec *EmailController) CreateEmailController(w http.ResponseWriter, r *http.Request){
	var sendEmailRequest request.CreateEmailRequest

	if err := json.NewDecoder(r.Body).Decode(&sendEmailRequest); err != nil {
		response.SendError(w, err.Error())
		return 
	}

	emailCommand, err := sendEmailRequest.ToCreateEmailCommand()
	if err != nil {
		response.SendError(w, err.Error())
        return
	}

	_, err = ec.service.SendEmail(emailCommand)
	if err != nil {
		response.SendError(w, err.Error())
		return
	}

	response.SendCreated(w)	
}

func (ec *EmailController) GetAllEmailController(w http.ResponseWriter, r *http.Request) {
	emails, err := ec.service.GetAllEmail()
	if err != nil {
		response.SendNotFound(w)
		return
	}

	resp := mapper.ToEmailListResponse(emails.Result)
	response.SendData(w, resp)	
}