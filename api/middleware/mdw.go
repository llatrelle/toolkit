package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/llatrelle/toolkit/api/app"
	"io/ioutil"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
)

// The original work was derived from Goji's middleware, source:
// https://github.com/zenazn/goji/tree/master/web/middleware

// Recoverer is a middleware that recovers from panics, logs the panic (and a
// backtrace), and returns a HTTP 500 (Internal Server Error) status if
// possible. Recoverer prints a request ID if one is provided.
//
// Alternatively, look at https://github.com/pressly/lg middleware pkgs.
func RecovererMDW(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {

				fmt.Fprintf(os.Stderr, "Panic: %+v\n", rvr)
				debug.PrintStack()
				app.Error(w, rvr.(error).Error())
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

//ParseParamsMDW es un middleware que obtiene los parámetros de la solicitud y los almacena en el contexto
func ParseParamsMDW(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		params := ContextParams(r)
		ctx := r.Context()
		//		params := requestParams{}
		var err error

		//params.Search = parseSearch(r.URL.Query().Get("q"))

		offset := r.URL.Query().Get("offset")
		if offset == "" {
			offset = r.URL.Query().Get("Offset")
		}
		offset = strings.TrimSpace(offset)

		limit := r.URL.Query().Get("limit")
		if limit == "" {
			limit = r.URL.Query().Get("Limit")
		}
		limit = strings.TrimSpace(limit)

		//		params.Limit = limitParam{}
		//		params.Limit.Offset, _ = strconv.ParseInt(offset, 10, 64)
		//		params.Limit.Count, err = strconv.ParseInt(limit, 10, 64)
		//		if err != nil {
		//			params.Limit.Count = -1
		//		}

		//		sl := &sortList{}
		//		err = sl.Parse(r.URL.Query().Get("sort"))

		//		params.Sort = sl

		//vuelvo a guardar los datos como json
		params.RequestBody.DataJSON, err = json.Marshal(params.RequestBody.Data)
		if err != nil {
			app.Fail(w, err.Error())
			return
		}

		//parseo el body
		r.Body = http.MaxBytesReader(w, r.Body, 20971520)
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			app.Fail(w, err.Error())
			return

		}
		//parseoArchivos
		if r.MultipartForm != nil && r.MultipartForm.File != nil {
			r.ParseMultipartForm(10 << 20)

			fls := r.MultipartForm.File["files"]

			for _, file := range fls {

				f, er := file.Open()
				defer f.Close()

				if er != nil {
					fmt.Printf("Error al abrir archivo %v", file.Filename)
					continue
				}
				tmpFile := fileData{}
				// tmpFile.Header = file.Header
				tmpFile.Filename = file.Filename
				tmpFile.Size = file.Size
				tmpFile.ContentType = file.Header.Get("Content-Type")
				tmpFile.File = f

				params.Files = append(params.Files, tmpFile)

			}
		}

		json.Unmarshal(b, &params.RequestBody.Data)

		//		fmt.Printf("data recibido: %#v\n", string(b))
		//		fmt.Printf("json recibido: %#v\n", params.RequestBody.DataJSON)
		params.QueryParams = r.URL.Query()
		params.Chi = chi.RouteContext(ctx)

		ctx = context.WithValue(r.Context(), "RequestParams", params)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

//Devuelve los parámetros del request actual
func ContextParams(r *http.Request) *requestParams {
	ctx := r.Context()
	v := ctx.Value("RequestParams")
	if v == nil {
		return &requestParams{}
	}
	return v.(*requestParams)
}
