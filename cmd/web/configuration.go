package main

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

var (
	conf Configuration
)

type Configuration struct {
	ENV_WEB_STATIC_DIR string `env:"ENV_WEB_STATIC_DIR" default:"./static"`
	ENV_WEB_ADDRESS    string `env:"ENV_WEB_ADDRESS" default:":8080"`

	ENV_WEB_POSTGRES_DB_URL string `env:"ENV_WEB_POSTGRES_DB"`
	ENV_WEB_LLM_GATEWAY_URL string `env:"ENV_WEB_LLM_GATEWAY_URL"`
	ENV_WEB_BIFROST_API_KEY string `env:"ENV_WEB_BIFROST_API_KEY"`
	ENV_WEB_LLM_MODEL       string `env:"ENV_WEB_LLM_MODEL" default:"gemini/gemini-1.5-flash"`

	ENV_WEB_RELOAD_TEMPLATES bool `env:"ENV_WEB_RELOAD_TEMPLATES" default:"false"`
}

func ParseConfiguration() {
	ctx := context.Background()

	if err := envconfig.Process(ctx, &conf); err != nil {
		panic(err)
	}
}
