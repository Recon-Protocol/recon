package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Recon-Protocol/recon/apps/api/internal/models"
)

func GetBlockDAG() (models.BlockDAG, error) {
	// Step 1: Get current tip hashes from blockdag info
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get("https://api.kaspa.org/info/blockdag")
	if err != nil {
		return models.BlockDAG{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.BlockDAG{}, err
	}

	var dagInfo struct {
		VirtualDAAScore  string   `json:"virtualDaaScore"`
		VirtualParentHashes []string `json:"virtualParentHashes"`
	}

	if err := json.Unmarshal(body, &dagInfo); err != nil {
		return models.BlockDAG{}, err
	}

	// Step 2: Fetch details for up to 5 tip blocks
	tips := dagInfo.VirtualParentHashes
	if len(tips) > 5 {
		tips = tips[:5]
	}

	var blockTips []models.BlockTip

	for _, hash := range tips {
		tip, err := fetchBlockTip(client, hash)
		if err != nil {
			continue
		}
		blockTips = append(blockTips, tip)
	}

	// Parse virtualDaaScore
	var daaScore uint64
	fmt.Sscanf(dagInfo.VirtualDAAScore, "%d", &daaScore)

	return models.BlockDAG{
		Tips:       blockTips,
		TipCount:   len(blockTips),
		VirtualDAA: daaScore,
		Status:     "online",
		Service:    "recon-api",
		APIVersion: "v1",
	}, nil
}

func fetchBlockTip(client *http.Client, hash string) (models.BlockTip, error) {
	url := fmt.Sprintf("https://api.kaspa.org/blocks/%s?includeTransactions=false", hash)

	resp, err := client.Get(url)
	if err != nil {
		return models.BlockTip{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.BlockTip{}, err
	}

	var block struct {
		Header struct {
			Timestamp string `json:"timestamp"`
		} `json:"header"`
		VerboseData struct {
			Hash           string   `json:"hash"`
			BlueScore      string   `json:"blueScore"`
			TransactionIds []string `json:"transactionIds"`
			IsChainBlock   bool     `json:"isChainBlock"`
		} `json:"verboseData"`
	}

	if err := json.Unmarshal(body, &block); err != nil {
		return models.BlockTip{}, err
	}

	var blueScore uint64
	fmt.Sscanf(block.VerboseData.BlueScore, "%d", &blueScore)

	var timestamp int64
	fmt.Sscanf(block.Header.Timestamp, "%d", &timestamp)

	return models.BlockTip{
		Hash:      hash[:8] + "..." + hash[56:],
		BlueScore: blueScore,
		Timestamp: timestamp / 1000, // ms to seconds
		TxCount:   len(block.VerboseData.TransactionIds),
		IsBlue:    block.VerboseData.IsChainBlock,
	}, nil
}