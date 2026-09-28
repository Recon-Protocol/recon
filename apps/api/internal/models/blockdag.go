package models

type BlockTip struct {
	Hash      string `json:"hash"`
	BlueScore uint64 `json:"blue_score"`
	Timestamp int64  `json:"timestamp"`
	TxCount   int    `json:"tx_count"`
	IsBlue    bool   `json:"is_blue"`
}

type BlockDAG struct {
	Tips        []BlockTip `json:"tips"`
	TipCount    int        `json:"tip_count"`
	VirtualDAA  uint64     `json:"virtual_daa_score"`
	Status      string     `json:"status"`
	Service     string     `json:"service"`
	APIVersion  string     `json:"api_version"`
}