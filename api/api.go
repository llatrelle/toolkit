package api

import "net/http"

type Api struct {
	port   string
	router http.Handler
}

func NewApi() *Api {
	app := new(Api)
	return app
}

func (app *Api) SetPort(port string) {
	app.port = port

}

func (app *Api) SetRouter(router http.Handler) {
	app.router = router

}

func (app *Api) Serve() error {
	// init minimal values
	app.init()
	err := http.ListenAndServe(":"+app.port, app.router)
	return err

}

//init set minimal values to serve
func (app *Api) init() {
	// default port
	if app.port == "" {
		app.port = "8080"
	}
	//default router
	if app.router == nil {
		app.router = defaultRoutes()
	}
}
