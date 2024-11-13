package api

import (
	"encoding/json"
	"net/http"

	"github.com/Andres-MK/internal/api/dto/mapper"
	"github.com/Andres-MK/internal/api/dto/request"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	emailCommand, err := sendEmailRequest.ToCreateEmailCommand()
	if err != nil {
		http.Error(w, "Objeto no válido", http.StatusBadRequest)
        return
	}

	result, err := ec.service.SendEmail(emailCommand)
	if err != nil {
		http.Error(w, "Error al enviar el email", http.StatusInternalServerError)
		return
	}

	response := mapper.ToEmailResponse(result.Result)
	
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    err = json.NewEncoder(w).Encode(response)
    if err != nil {
        http.Error(w, "Error al enviar la respuesta", http.StatusInternalServerError)
    }
}