package kaspa

type networkResponse struct {
	NetworkName     string  `json:"networkName"`
	BlockCount      uint64  `json:"blockCount,string"`
	HeaderCount     uint64  `json:"headerCount,string"`
	Difficulty      float64 `json:"difficulty"`
	VirtualDAAScore uint64  `json:"virtualDaaScore,string"`
}
