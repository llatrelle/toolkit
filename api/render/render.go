// Package app El paquete application contiene todos los métodos y recursos comunes a la aplicación
package render

import (
	"github.com/unrolled/render"
	"io"
	"toolkit/logger"
)

func init() {
	renderer = render.New()

}

var renderer *render.Render

// Copia del body de response
var Rendered *appResponse

// Estructura de datos para respuestas al cliente
type appResponse struct {
	Data     interface{} `json:"data"`
	Status   string      `json:"status"`
	Code     int         `json:"code"`
	Message  interface{} `json:"message,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
}

// jSend Renderiza la respuesta en un formato deseado
func jSend(w io.Writer, status string, cod int, message interface{}, data interface{}, metadata interface{}) {
	response := appResponse{
		Status:   status,
		Code:     cod,
		Message:  message,
		Data:     data,
		Metadata: metadata,
		//Errors:   errors,
	}
	Rendered = &response
	renderer.JSON(w, cod, response)

}

// Error Devuelve un mensaje de Error al Cliente
func Error(w io.Writer, message interface{}, cod ...int) {
	if cod == nil {
		cod = append(cod, 500)
	}

	logger.Trace("render", "error", message)
	jSend(w, "error", cod[0], message, nil, nil)
}

// Fail Devuelve un mensaje de Falla al Cliente
func Fail(w io.Writer, data interface{}, cod ...int) {
	if cod == nil {
		cod = append(cod, 400)
	}
	if d, ok := data.(string); ok {
		logger.Trace("render", "fail", data)
		jSend(w, "fail", cod[0], d, nil, nil)
	} else {
		jSend(w, "fail", cod[0], "", data, nil)
	}
}

// Success Devuelve un mensaje  Success al Cliente
func Success(w io.Writer, data interface{}, metadata interface{}, cod ...int) {
	if cod == nil {
		cod = append(cod, 200)
	}
	jSend(w, "success", cod[0], "", data, metadata)
}

// SuccessWithWarning Normalmente un warning es un mensaje de Ok, con un aviso
func SuccessWithWarning(w io.Writer, data interface{}, warnings interface{}, metadata interface{}, errors interface{}, cod ...int) {
	if cod == nil {
		cod = append(cod, 200)
	}
	jSend(w, "warning", cod[0], warnings, data, metadata)
}
