package main

import (
	"net/http"

	"github.com/DAlba-sudo/verb"
)

func Routes(s *verb.Server) {
	Pages(s)
	API(s)
}

func Pages(s *verb.Server) {
	_, err := s.Register(http.MethodGet, "",
		"./html/base.html",
		"./html/pages/index.html")

	_, err = s.Register(http.MethodGet, "/htmx/product/create",
		"./html/forms/product-create.html")

	if err != nil {
		panic(err)
	}
}

func API(s *verb.Server) {

}
