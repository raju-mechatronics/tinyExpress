package te

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Application struct {
	router
	server  *http.Server
	config  AppConfig
	handler http.Handler
}

func App(config ...AppConfig) *Application {
	serverConfig := AppConfig{
		Port: 4000,
		Host: "127.0.0.1",
		AllowedMethod: []RequestType{
			RequestTypeAny,
			RequestTypeConnect,
			RequestTypeDelete,
			RequestTypeGet,
			RequestTypeHead,
			RequestTypeOptions,
			RequestTypePatch,
			RequestTypePost,
			RequestTypePut,
			RequestTypeTrace,
		},
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

	app := &Application{
		config: serverConfig,
		server: &http.Server{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("Processing request", r.URL.Path, " ", time.Now())
				time.Sleep(5 * time.Second)
				fmt.Println("Processing request", r.URL.Path, " ", time.Now())
				fmt.Fprintln(w, "Response for", r.URL.Path)
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
