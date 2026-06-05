package main

import (
	"github.com/DAlba-sudo/verb"
)

func main() {
	ParseConfiguration()
	s := verb.CreateServer(conf.ENV_WEB_STATIC_DIR)

	Routes(s)
	if err := s.Serve(conf.ENV_WEB_ADDRESS); err != nil {
		panic(err)
	}
}
