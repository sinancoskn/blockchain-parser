package ethereum

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	url string
}

func NewClient(url string) *Client {
	return &Client{
		url: url,
	}
}

func (c *Client) GetFirstAccount() (string, error) {
	result, err := c.Request("eth_accounts", nil)
	if err != nil {
		return "", fmt.Errorf("failed to fetch accounts: %w", err)
	}

	var accounts []string
	if err := json.Unmarshal(result, &accounts); err != nil {
		return "", fmt.Errorf("failed to unmarshal accounts: %w", err)
	}

	return accounts[0], nil
}

func (c *Client) GetBlockByNumber(blockNumber string) (json.RawMessage, error) {
	return c.Request("eth_getBlockByNumber", []interface{}{blockNumber, true})
}

func (c *Client) GetCurrentBlock() (string, error) {
	result, err := c.Request("eth_blockNumber", nil)
	if err != nil {
		return "", fmt.Errorf("failed to get the current block number: %w", err)
	}

	var blockNumber string
	if err := json.Unmarshal(result, &blockNumber); err != nil {
		return "", fmt.Errorf("failed to unmarshal block number: %w", err)
	}

	return blockNumber, nil
}

func (c *Client) GetTransactions(address string) ([]Transaction, error) {
	currentBlock, err := c.GetCurrentBlock()
	if err != nil {
		return nil, fmt.Errorf("failed to get the current block: %w", err)
	}

	const blocksToCheck = 10
	transactions := []Transaction{}

	for i := 0; i < blocksToCheck; i++ {
		blockNumber := fmt.Sprintf("0x%x", hexToInt(currentBlock)-i)
		blockData, err := c.GetBlockByNumber(blockNumber)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch block %s: %w", blockNumber, err)
		}

		var block Block
		if err := json.Unmarshal(blockData, &block); err != nil {
			return nil, fmt.Errorf("failed to unmarshal block: %w", err)
		}

		for _, tx := range block.Transactions {
			if tx.From == address || tx.To == address {
				transactions = append(transactions, tx)
			}
		}
	}

	return transactions, nil
}

func (c *Client) SendTransaction(from, to, value string) (string, error) {
	payload := map[string]interface{}{
		"from":     from,
		"to":       to,
		"value":    value,
		"password": "password",
	}

	result, err := c.Request("eth_sendTransaction", []interface{}{payload})
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %w", err)
	}

	var txHash string
	if err := json.Unmarshal(result, &txHash); err != nil {
		return "", fmt.Errorf("failed to unmarshal transaction hash: %w", err)
	}

	return txHash, nil
}

func (c *Client) Request(method string, params []interface{}) (json.RawMessage, error) {
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  method,
		"params":  params,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(c.url, "application/json", bytes.NewReader(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var result struct {
		Result json.RawMessage `json:"result"`
		Error  *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("JSON-RPC error: %d - %s", result.Error.Code, result.Error.Message)
	}

	return result.Result, nil
}

func hexToInt(hexStr string) int {
	var result int
	fmt.Sscanf(hexStr, "0x%x", &result)
	return result
}
