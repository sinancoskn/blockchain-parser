package integration

import (
	"blockchain-parser/internal/ethereum"
	"blockchain-parser/internal/parser"
	"testing"
	"time"
)

func init() {
	SetupTest()
}

func TestParser_SubscribeAndWaitNotification(t *testing.T) {
	err := Container.Invoke(func(parser parser.Parser, client *ethereum.Client) {
		success := parser.Subscribe(TEST_ADDRESS_2)
		if !success {
			t.Fatalf("Failed to subscribe to address: %s", TEST_ADDRESS_2)
		}

		_, err := client.SendTransaction(TEST_ADDRESS_1, TEST_ADDRESS_2, "0xde0b6b3a7640000")
		if err != nil {
			t.Fatalf("Failed to send transaction: %v", err)
		}

		var transactions []ethereum.Transaction
		for i := 0; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)

			transactions = parser.GetTransactions(TEST_ADDRESS_2)
			if len(transactions) == 1 {
				break
			}
		}

		if len(transactions) != 1 {
			t.Fatalf("Failed to get transaction for address %s, expected 1 but got %d", TEST_ADDRESS_2, len(transactions))
		}
	})

	if err != nil {
		t.Fatalf("Failed to invoke for TestParser_SubscribeAndWaitNotification: %v", err)
	}
}

func TestParser_Unsubscribe(t *testing.T) {
	err := Container.Invoke(func(parser parser.Parser) {
		parser.Subscribe(TEST_ADDRESS_1)

		success := parser.Unsubscribe(TEST_ADDRESS_1)
		if !success {
			t.Fatalf("Failed to unsubscribe address: %s", TEST_ADDRESS_1)
		}

		success = parser.Unsubscribe(TEST_ADDRESS_1)
		if success {
			t.Fatalf("Unsubscribing non-subscribed address should fail: %s", TEST_ADDRESS_1)
		}
	})

	if err != nil {
		t.Fatalf("Failed to invoke for TestParser_Unsubscribe: %v", err)
	}
}
