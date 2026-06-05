package core

import (
	"log"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var (
	Logger = slog.New(slog.NewJSONHandler(log.Writer(), nil))

	DB         *gorm.DB
	LLMGateway *openai.Client
)

func InitDatabase(connectionUrl string) {
	db, err := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{})
	if err != nil {
		return
	}

	DB = db
}

func InitLLMGateway(gatewayUrl string, apikey string) {
	client := openai.NewClient(
		option.WithBaseURL(gatewayUrl),
		option.WithHeader("x-bf-vk", apikey),
	)
}
