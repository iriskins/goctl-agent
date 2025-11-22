package core

/*
Application is a layer responsible for dependency injection and an actual execution loop
*/

import (
	"context"
	"net/http"
	"os"
	"sync"

	"github.com/coreos/go-systemd/v22/dbus"
)

type Config struct {
	Addr string
	Port string
}

func NewConfig(a string, p string) *Config {
	return &Config{Addr: a, Port: p}
}

type Application struct {
	*Config
	ShouldClose chan os.Signal
	Connection  *dbus.Conn
	Context     context.Context
	Service     *Service
	Handler     *Handler
	Server      HttpServer
	mu          sync.RWMutex
}

func NewApplication() *Application {
	NewLogger()
	ctx := context.Background()
	LogInfo("Started. Connecting to dbus...")
	conn, err := dbus.NewSystemConnectionContext(ctx)
	if err != nil {
		LogError("Failed to connect to dbus!")
		LogError("Try starting the application in privileged mode...", "err", err)
		os.Exit(1)
	}
	LogInfo("Connected. Initializing services...")
	App := &Application{Connection: conn, Context: ctx, ShouldClose: make(chan os.Signal, 1)}
	service := NewService(App.Connection, App.Context, &App.mu)
	App.Service = service
	App.Service.Initialize()
	handler := NewHandler(service)
	App.Handler = handler
	App.Server = *HttpInit(*App.Handler)
	return App
}

func (a *Application) Run() {
	a.Server.GET("/services", a.Handler.GetServices)
	a.Server.POST("/services/enable", a.Handler.EnableUnits)
	LogInfo("Application initialized successfully! Running on port 8080...")
	http.ListenAndServe("0.0.0.0:8080", a.Server.mux)
}

func (a *Application) Stop() {
	a.Connection.Close()
	LogInfo("Connection to dbus has been successfully closed, exiting...")
	os.Exit(0)
}
