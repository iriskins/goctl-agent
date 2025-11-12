package core

import (
	"net/http"
)

type HttpServer struct {
	mux *http.ServeMux
	Handler
}

func HttpInit(h Handler) *HttpServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.Services)
	return &HttpServer{mux: mux}
}
