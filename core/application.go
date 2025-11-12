/*
Application is a layer responsible for dependency injection and an actual execution loop
*/
package core

import (
	"context"
	"net/http"
	"os"

	"github.com/coreos/go-systemd/v22/dbus"
)

type Config struct {
}

type Application struct {
	*Config
	ShouldClose chan os.Signal
	Connection  *dbus.Conn
	Context     context.Context
	Service     Service
	Handler     Handler
	Server      HttpServer
}

func NewApplication() *Application {
	LogInit()
	ctx := context.Background()
	LogInfo("Started. Connecting to dbus...")
	conn, err := dbus.NewSystemConnectionContext(ctx)
	if err != nil {
		LogError("Failed to connect to dbus!")
		LogError("Try starting the application in privileged mode...", "err", err)
		os.Exit(1)
	}
	LogInfo("Connected. Initializing services...")
	App := &Application{Connection: conn, Context: ctx}
	service := NewService(App)
	App.Service = *service
	handler := NewHandler(service)
	App.Handler = *handler
	App.ShouldClose = make(chan os.Signal, 1)
	App.Server = *HttpInit(App.Handler)
	return App
}

func (a *Application) Run() {
	LogInfo("Application initialized successfully! Running on port 8080...")
	http.ListenAndServe("0.0.0.0:8080", a.Server.mux)
}

func (a *Application) Stop() {
	//closing connection, cleaning, etc.
	a.Connection.Close()
	os.Exit(0)
}
