package handler

import (
	"github.com/llatrelle/toolkit/api/app"
	"github.com/llatrelle/toolkit/api/middleware"
	"github.com/llatrelle/toolkit/model"
	"net/http"
)

type GenericHandler struct {
	Model  model.Modeler
	Models []model.Modeler
}

func (g *GenericHandler) Get(pk string, w http.ResponseWriter, r *http.Request) {
	m := g.Model
	modelKeys := g.GetParams(r)
	g.GetParams(r)
	err := model.Get(&m, modelKeys[pk])
	if err != nil {
		app.Error(w, err, 500)
		return
	}
	app.Success(w, m, nil, http.StatusOK)
}

func (g *GenericHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	l := g.Models
	err := model.GetAll(g.Model, &l)
	if err != nil {
		app.Error(w, err, 500)
		return
	}
	app.Success(w, l, nil, 200)
	return

}

func (g *GenericHandler) Create(w http.ResponseWriter, r *http.Request) {
	rp := middleware.ContextParams(r)
	if rp.RequestBody.Data == nil {
		app.Error(w, "invalid request", 400)
		return
	}
	payload := rp.RequestBody.Data.(map[string]interface{})
	m := g.Model

	app.Success(w, payload, nil, http.StatusOK)
}

func (g *GenericHandler) Delete(pk string, w http.ResponseWriter, r *http.Request) {

	modelKeys := g.GetParams(r)
	_, err := model.Delete(g.Model, modelKeys[pk])
	if err != nil {
		app.Error(w, err, 500)
		return
	}

	app.Success(w, nil, nil, http.StatusOK)
}

func (g *GenericHandler) Update(pk string, w http.ResponseWriter, r *http.Request) {

}

func (g *GenericHandler) GetParams(r *http.Request) map[string][]string {

	rp := middleware.ContextParams(r)
	urlParamsKeys := rp.Chi.URLParams.Keys
	urlQueryParams := rp.QueryParams
	var modelKeys map[string][]string
	modelKeys = make(map[string][]string)

	for _, b := range urlParamsKeys {

		if b != "*" {

			modelKeys[b] = append(modelKeys[b], rp.Chi.URLParam(b))

		}
	}
	for k, v := range urlQueryParams {

		modelKeys[k] = v
		//		fmt.Printf("K: %#v", k)
		//		fmt.Printf("V: %#v", v)

	}
	//	fmt.Printf("modelKeys %v", modelKeys)
	return modelKeys
}
