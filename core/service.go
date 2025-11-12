/*
Service is a layer which is responsible for scraping data,
using dbus, provided by Application class
*/
package core

import (
	"fmt"

	"github.com/coreos/go-systemd/v22/dbus"
)

type Service struct {
	*Application
	ch chan string
}

func NewService(a *Application) *Service {
	ch := make(chan string)
	return &Service{Application: a, ch: ch}
}

func (s *Service) ListAllUnits() (units []dbus.UnitStatus, err error) {
	units, err = s.Connection.ListUnitsContext(s.Context)
	if err != nil {
		return
	}
	return
}

func (s *Service) RestartUnitByName(name string) {
	status, err := s.Connection.RestartUnitContext(s.Context, name, "replace", s.ch)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", status)
}
