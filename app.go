package te

import (
	"fmt"
	"net/http"
	"strconv"
)

type Application struct {
	//Router
	server              *http.Server
	config              AppConfig
	handler             http.Handler
	beforeSendListeners []func(req *Request, res *Response)
}

func App(config ...AppConfig) *Application {
	serverConfig := AppConfig{
		Port: 4000,
		Host: "127.0.0.1",
	}

	if len(config) > 0 {
		if config[0].Port != 0 {
			serverConfig.Port = config[0].Port
		}
		if config[0].Host != "" {
			serverConfig.Host = config[0].Host
		}
		if len(config[0].AllowedMethod) > 0 {
			serverConfig.AllowedMethod = config[0].AllowedMethod
		}
	}

	fmt.Println(serverConfig)

	var app *Application
	app = &Application{
		config: serverConfig,
		server: &http.Server{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				handle(app, w, r)
			}),
		},
	}

	fmt.Print(app.server.Addr)

	return app
}

func (app *Application) Listen(port ...int) error {
	if len(port) > 0 {
		app.config.Port = port[0]
	}

	addr := app.config.Host + ":" + strconv.Itoa(app.config.Port)
	app.server.Addr = addr

	fmt.Println("Listening on", app.server.Addr)
	return app.server.ListenAndServe()
}

func (app *Application) Close() error {
	return app.server.Close()
}

func (app *Application) GetServer() *http.Server {
	return app.server
}

func (app *Application) GetConfig() *AppConfig {
	return &app.config
}

func (app *Application) GetHandler() *http.Handler {
	return &app.handler
}

func (app *Application) OnBeforeSend(listener func(req *Request, res *Response)) {
	app.beforeSendListeners = append(app.beforeSendListeners, listener)
}

func (app *Application) BeforeSend(req *Request, res *Response) {
	for _, listener := range app.beforeSendListeners {
		listener(req, res)
	}
}

func handle(app *Application, w http.ResponseWriter, r *http.Request) {
	req, err := newRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := NewResponse(&w)

	//app.Resolve(req, res)

	app.BeforeSend(req, res)

	// send the response
	if !res.resolved {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// write the response Header
	for key, values := range *res.header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// write the response bod
	w.Write(*res.body)

}
