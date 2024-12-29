package parser

import (
	"blockchain-parser/internal/ethereum"
	"blockchain-parser/internal/storage"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type parser struct {
	client  *ethereum.Client
	storage storage.Storage
}

const (
	subscriptionPrefix = "subscription:"
	transactionPrefix  = "transactions:"
	lastBlockKey       = "lastParsedBlock"
)

func NewParser(client *ethereum.Client, storage storage.Storage) Parser {
	return &parser{
		client:  client,
		storage: storage,
	}
}

func (p *parser) InitialLatesBlock() error {
	currentBlock, err := p.client.GetCurrentBlock()
	if err != nil {
		return err
	}

	currentBlockInt := hexToInt(currentBlock)

	p.storage.Set(lastBlockKey, currentBlockInt)

	return nil
}

func (p *parser) GetCurrentBlock() int {
	value, exists := p.storage.Get(lastBlockKey)
	if !exists {
		return 0
	}

	if blockNumber, ok := value.(int); ok {
		return blockNumber
	}

	return 0
}

func (p *parser) Subscribe(address string) bool {
	key := subscriptionPrefix + address
	_, exists := p.storage.Get(key)
	if exists {
		return false
	}

	p.storage.Set(key, true)
	return true
}

func (p *parser) Unsubscribe(address string) bool {
	key := subscriptionPrefix + address
	_, exists := p.storage.Get(key)
	if !exists {
		return false
	}

	p.storage.Delete(key)
	return true
}

func (p *parser) GetTransactions(address string) []ethereum.Transaction {
	key := transactionPrefix + address
	data, _ := p.storage.Get(key)
	if transactions, ok := data.([]ethereum.Transaction); ok {
		return transactions
	}

	return nil
}

func (p *parser) ProcessBlocks() error {
	currentBlockHex, err := p.client.GetCurrentBlock()
	if err != nil {
		return fmt.Errorf("failed to get current block: %w", err)
	}

	currentBlock := hexToInt(currentBlockHex)

	lastParsedBlock := p.GetCurrentBlock()
	for i := lastParsedBlock + 1; i <= currentBlock; i++ {
		blockHex := fmt.Sprintf("0x%x", i)
		blockData, err := p.client.GetBlockByNumber(blockHex)
		if err != nil {
			return fmt.Errorf("failed to fetch block %s: %w", blockHex, err)
		}

		var block ethereum.Block
		if err := json.Unmarshal(blockData, &block); err != nil {
			return fmt.Errorf("failed to unmarshal block: %w", err)
		}

		p.processTransactions(block.Transactions)

		p.storage.Set(lastBlockKey, i)
	}

	return nil
}

func (p *parser) processTransactions(transactions []ethereum.Transaction) {
	for _, tx := range transactions {
		fromSubscribed := p.isSubscribed(tx.From)
		toSubscribed := p.isSubscribed(tx.To)

		if fromSubscribed {
			p.addTransaction(tx.From, tx)
			p.notify(tx.From, tx)
		}
		if toSubscribed {
			p.addTransaction(tx.To, tx)
			p.notify(tx.To, tx)
		}
	}
}

func (p *parser) isSubscribed(address string) bool {
	key := subscriptionPrefix + address
	_, exists := p.storage.Get(key)
	return exists
}

func (p *parser) addTransaction(address string, tx ethereum.Transaction) {
	key := transactionPrefix + address
	data, _ := p.storage.Get(key)

	var txList []ethereum.Transaction
	if data != nil {
		txList = data.([]ethereum.Transaction)
	}

	txList = append(txList, tx)
	p.storage.Set(key, txList)
}

func (p *parser) notify(address string, tx ethereum.Transaction) {
	// TODO: Implement notification
	fmt.Printf("New transaction for %s: %+v\n", address, tx)
}

func (p *parser) StartMonitoring(ctx context.Context) {
	fmt.Println("StartMonitoring invoked")
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Shutting down monitoring...")
				return
			default:
				err := p.ProcessBlocks()
				if err != nil {
					fmt.Printf("Error processing blocks: %v\n", err)
				}

				fmt.Println("Running monitoring...")

				time.Sleep(5 * time.Second)
			}
		}
	}()
}

func hexToInt(hexStr string) int {
	var result int
	fmt.Sscanf(hexStr, "0x%x", &result)
	return result
}
