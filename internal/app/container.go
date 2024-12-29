package app

import (
	"blockchain-parser/config"
	"blockchain-parser/internal/ethereum"
	"blockchain-parser/internal/logger"
	"blockchain-parser/internal/parser"
	"blockchain-parser/internal/storage"
	"context"
	"log"

	"go.uber.org/dig"
)

func InitializeContainer() *dig.Container {
	config.LoadConfig()

	logger.InitLogger()

	c := dig.New()

	c.Provide(storage.NewMockStorage)
	c.Provide(parser.NewParser)
	c.Provide(parser.NewParserHandler)

	c.Provide(func() *ethereum.Client {
		return ethereum.NewClient(config.GlobalConfig.EthereumRPCURL)
	})

	err := c.Invoke(func(p parser.Parser) {
		if err := p.InitialLatesBlock(); err != nil {
			log.Fatalf("failed to start monitoring: %w", err)
		}

		ctx, _ := context.WithCancel(context.Background())

		p.StartMonitoring(ctx)
	})
	if err != nil {
		log.Fatalf("failed to start monitoring: %w", err)
	}

	logger.Log.Info("Container initialized")

	return c
}
