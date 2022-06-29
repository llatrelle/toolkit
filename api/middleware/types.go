package middleware

import (
	"github.com/go-chi/chi"
	"mime/multipart"
	"net/url"
)

type requestParams struct {
	//Contiene la lista de campos para el sort del resultado
	//Sort *sortList

	//Contiene el texto para realización de search (parámetro Get q)
	Search *searchParam

	//Contiene los parámetros de offset y limit del request
	//Limit limitParam

	//contiene los datos parseados del  body recibido via POST/PUT
	RequestBody requestBody

	//Realiza la búsqueda de un parámetro Get recibido en el request
	QueryParams url.Values

	//Referencia al Contexto del ruteador chi
	Chi *chi.Context

	//Contiene la información del usuario obtenida del token jwt
	//UserInfo userInfo

	//Permisos permiso

	//Contiene los archivos y la informacion de los mismos del form-data
	Files []fileData
}
type fileData struct {
	Filename string

	File        multipart.File
	Size        int64
	ContentType string
}

type requestBody struct {
	Data     interface{} `json:"data"`
	DataJSON []byte      `json:"-"`
	//	Params 	 interface{} 	`json:"params"`
}

type searchParam struct {
	Str  string
	Type string
}
