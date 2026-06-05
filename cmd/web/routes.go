package main

import (
	"net/http"

	"github.com/DAlba-sudo/verb"
)

func Routes(s *verb.Server) {
	_, err := s.Register(http.MethodGet, "",
		"./html/base.html",
		"./html/pages/index.html")

	if err != nil {
		panic(err)
	}
}
