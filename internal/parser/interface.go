package parser

import (
	"blockchain-parser/internal/ethereum"
	"context"
)

type Parser interface {
	StartMonitoring(context.Context)
	InitialLatesBlock() error
	GetCurrentBlock() int
	Subscribe(address string) bool
	Unsubscribe(address string) bool
	GetTransactions(address string) []ethereum.Transaction
}
