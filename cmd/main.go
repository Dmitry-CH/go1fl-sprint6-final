package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	l := log.New(log.Default().Writer(), "serv ", log.LstdFlags|log.Lshortfile)

	s := server.New(l)

	err := s.Http.ListenAndServe()
	if err != nil {
		s.Logger.Fatal(err)
	}
}
