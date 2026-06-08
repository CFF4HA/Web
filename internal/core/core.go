package core

import (
	"context"
	"log"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/CFF4HA/Web/internal/types"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var (
	Logger = slog.New(slog.NewJSONHandler(log.Writer(), nil))

	DB         *gorm.DB
	LLMGateway *openai.Client
)

func InitDatabase(connectionUrl string) {
	if connectionUrl == "" {
		Logger.Warn("Database connection URL is empty")
		return
	}

	db, err := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		Logger.Error("Failed to connect to database", "error", err)
		return
	}
	DB = db

	// Enable the extensions for UUID.
	tx := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	if tx.Error != nil {
		Logger.Error("Failed to create uuid-ossp extension", "error", tx.Error)
		return
	}

	// Migrate the database types to the database instance.
	if err := DB.AutoMigrate(&types.Model{}, &types.CommonResource{}, &types.Label{}); err != nil {
		Logger.Error("Failed to auto-migrate database schema", "error", err)
	}

	Logger.Info("Database connection established successfully")
}

func InitLLMGateway(gatewayUrl string, apikey string, model string) {
	if gatewayUrl == "" || apikey == "" {
		return
	}

	client := openai.NewClient(
		option.WithBaseURL(gatewayUrl),
		option.WithHeader("x-bf-vk", apikey),
	)

	// WARNING: This is potentially unsafe.
	LLMGateway := &openai.Client{}
	*LLMGateway = client

	ctx := context.Background()
	models, err := LLMGateway.Models.List(ctx) // Test connection
	if err != nil {
		Logger.Error("Failed to test LLM gateway connection", "error", err)
		return
	}

	Logger.Info("LLM gateway connection established successfully", "available_models", len(models.Data), "data", models.RawJSON())
}
