package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Andres-MK/internal/api"
	"github.com/Andres-MK/internal/api/dto/request"
	"github.com/Andres-MK/internal/api/dto/response"
	"github.com/Andres-MK/internal/application/command"
	"github.com/Andres-MK/internal/application/query"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type MockEmailService struct{
}

func (m *MockEmailService) GetAllEmail() (*query.ListEmailsQueryResult, error) {
	return nil, nil
}

func (m *MockEmailService) SendEmail(*command.CreateEmailCommand) (*command.CreateEmailCommandResult, error){
	return nil, nil
}

func TestCreateEmailController(t *testing.T) {
	email := request.CreateEmailRequest{
			EmailFrom: 	"onboarding@resend.dev",
			EmailTo: 	"andres.mkern@gmail.com",
			Subject:	"Esta es una prueba",
			Html: 		"<strong>Este es una html </strong>",
			FirstName: 	"Holaaaaa",
			LastName: 	"Aguero",
			Message: 	"Hola que tal",
			Company: 	"SPA",
	}

	jsonData, _ := json.Marshal(email)
	req, err := http.NewRequest("POST", "/email", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	mockService := &MockEmailService{}

	rr := httptest.NewRecorder()
	r := mux.NewRouter() 
	
	api.NewEmailController(r, mockService) 
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Se esperaba un código de estado 201")

	var responseBody response.Response
	err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatal(err)
	}

	expectedData := response.Response {
		Status: 201,
		Data: nil,
		Message: "OK",
	  }

	assert.Equal(t, expectedData, responseBody)
}