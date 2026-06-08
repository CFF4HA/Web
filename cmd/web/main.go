package main

import (
	"github.com/CFF4HA/Web/internal/core"
	"github.com/DAlba-sudo/verb"
)

func main() {
	ParseConfiguration()
	InitializeAuxiliaryServices()

	s := verb.CreateServer(conf.ENV_WEB_STATIC_DIR)
	if conf.ENV_WEB_RELOAD_TEMPLATES {
		s.Options.ReloadTemplates = true
	}

	Routes(s)
	if err := s.Serve(conf.ENV_WEB_ADDRESS); err != nil {
		panic(err)
	}
}

func InitializeAuxiliaryServices() {
	core.InitDatabase(conf.ENV_WEB_POSTGRES_DB_URL)
	core.InitLLMGateway(conf.ENV_WEB_LLM_GATEWAY_URL, conf.ENV_WEB_BIFROST_API_KEY, conf.ENV_WEB_LLM_MODEL)
}
