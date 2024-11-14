package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int					`json:"status"`
	Data        interface{}			`json:"data"`
	Message     string 				`json:"message"`
	contentType string
	respWrite   http.ResponseWriter
}

func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		respWrite:   rw,
		contentType: "application/json",
	}
}

// Método que imprime la respuesta por pantalla
func (r *Response) Send() {
	r.respWrite.Header().Set("Content-Type", r.contentType)
	r.respWrite.WriteHeader(r.Status)
	
	output, _ := json.Marshal(&r)
	fmt.Fprintln(r.respWrite, string(output))
}

// Se envia respuesta al cliente - Éxito
func SendData(rw http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(rw)
	response.Data = data
	response.Send()
}

// Se envia respuesta al cliente - Creado
func (resp *Response) Created() {
	resp.Status = http.StatusCreated
	resp.Message = "OK"
}

func SendCreated(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.Created()
	response.Send()
}

// Se envia respuesta al cliente - No encontrado
func (resp *Response) NotFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Recurso no encontrado"
}

func SendNotFound(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.NotFound()
	response.Send()
}

// Se envia respuesta al cliente - Error
func (resp *Response) Error(message string) {
	resp.Status = http.StatusBadRequest
	resp.Message = message
}

func SendError(rw http.ResponseWriter, message string) {
	response := CreateDefaultResponse(rw)
	response.Error(message)
	response.Send()
}