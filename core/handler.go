package core

/*
Handler is a layer responsible for processing http requests
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	*Service
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func NewHandler(s *Service) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) GetServices(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	units, err := h.ListUnits()
	if err != nil {
		//Logging????
		fmt.Print(err)
	}
	// Using identation for testing purposes, since it is hard to read json, containing services
	jUnits, err := json.MarshalIndent(units, "", "    ")
	if err != nil {
		LogInfo("Failed to... BlahBLAH")
	}
	w.Write(jUnits)
	log.Printf("Request took: %s", time.Since(start))
}

func (h *Handler) EnableUnits(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var units EnableUnits
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&units); err != nil {
		fmt.Print("Failed to decode json data")
	}
	for _, svc := range units.Services {
		fmt.Print(svc)
	}

	log.Printf("Request took: %s", time.Since(start))
}
