package services

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"

	"github.com/Recon-Protocol/recon/apps/api/internal/models"
)

func GetReport() (models.Report, error) {

	analytics, err := GetAnalytics()
	if err != nil {
		return models.Report{}, err
	}

	network, err := GetNetwork()
	if err != nil {
		return models.Report{}, err
	}

	// ── DATE / CALENDAR WEEK ─────────────────────────────────────────────
	now := time.Now()
	_, cw := now.ISOWeek()
	date := now.Format("02.01.2006")
	calendarWeek := fmt.Sprintf("CW%d", cw)

	// ── EXCHANGE FLOWS (kaspa-lens.com) ──────────────────────────────────
	exchangeFlows, err := fetchExchangeFlows()
	if err != nil {
		exchangeFlows = []models.ExchangeFlow{}
	}

	return models.Report{
		// Header
		ReportNumber: "—",
		Date:         date,
		CalendarWeek: calendarWeek,
		Network:      network.Network,

		// Network
		Hashrate:      fmt.Sprintf("%.2f PH/s", analytics.HashratePHS),
		DAAScore:      fmt.Sprintf("%d", network.VirtualDAAScore),
		NetworkStatus: "Healthy",

		// Mining
		BlockReward:   "₭" + analytics.BlockReward,
		NextReduction: fmt.Sprintf("in ~%.1f days", analytics.DaysToReduction),

		// Supply
		CirculatingSupply: fmt.Sprintf("%.2fB KAS", float64(analytics.CirculatingSupply)/1e17),
		PercentMined:      fmt.Sprintf("%.2f%%", analytics.PercentMined),

		// Exchange Flows
		ExchangeFlows: exchangeFlows,

		// Manual fields — filled by analyst
		TPS:         "—",
		BPS:         "—",
		Nodes:       "—",
		MinerCount:  "—",
		PoolRevenue: "—",
		L2Activity:  "—",

		// Metadata
		Status:     analytics.Status,
		Service:    analytics.Service,
		APIVersion: analytics.APIVersion,
	}, nil
}

// fetchExchangeFlows pulls named exchange addresses from kaspa-lens.com
func fetchExchangeFlows() ([]models.ExchangeFlow, error) {
	req, err := http.NewRequest("GET", "https://kaspa-lens.com/public-api/public/web/wallet/addresses", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Conversion-Currency", "USD")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var raw []struct {
		Address struct {
			Name string `json:"name"`
		} `json:"address"`
		CurrentBalanceKas float64 `json:"currentBalanceKas"`
		Change24h         *struct {
			ChangeBalance float64 `json:"changeBalance"`
		} `json:"change24h"`
		Change7d *struct {
			ChangeBalance float64 `json:"changeBalance"`
		} `json:"change7d"`
		Change30d *struct {
			ChangeBalance float64 `json:"changeBalance"`
		} `json:"change30d"`
	}

	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	// Only include named exchanges
	exchanges := map[string]bool{
		"MEXC": true, "Bybit": true, "Gate.io": true,
		"KuCoin": true, "Kraken": true, "Bitget": true,
	}

	seen := map[string]bool{}
	flows := []models.ExchangeFlow{}

	for _, entry := range raw {
		name := entry.Address.Name
		if !exchanges[name] || seen[name] {
			continue
		}
		seen[name] = true

		flow := models.ExchangeFlow{
			Name:      name,
			BalanceKas: math.Round(entry.CurrentBalanceKas),
		}
		if entry.Change24h != nil {
			flow.Change24h = math.Round(entry.Change24h.ChangeBalance)
		}
		if entry.Change7d != nil {
			flow.Change7d = math.Round(entry.Change7d.ChangeBalance)
		}
		if entry.Change30d != nil {
			flow.Change30d = math.Round(entry.Change30d.ChangeBalance)
		}
		flows = append(flows, flow)
	}

	return flows, nil
}