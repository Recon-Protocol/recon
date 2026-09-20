package services

import (
	"github.com/Recon-Protocol/recon/apps/api/internal/models"
)

const (
	maxNetworkSecurity  = 35
	maxSupplyState      = 25
	maxEmissionPressure = 20
	maxProtocolEvent    = 20
)

func GetIntelligence() (models.Intelligence, error) {
	analytics, err := GetAnalytics()
	if err != nil {
		return models.Intelligence{}, err
	}

	signals := []string{}

	// NETWORK SECURITY (35%)
	const hashrateTarget = 400.0
	networkRaw := analytics.HashratePHS / hashrateTarget
	if networkRaw > 1.0 {
		networkRaw = 1.0
	}
	networkScore := int(networkRaw * float64(maxNetworkSecurity))

	if analytics.HashratePHS >= 300 {
		signals = append(signals, "Hashrate above 300 PH/s")
	}
	if analytics.HashratePHS >= 350 {
		signals = append(signals, "Hashrate approaching ATH range")
	}

	// SUPPLY STATE (25%)
	supplyScore := 0
	if analytics.PercentMined >= 90.0 {
		supplyRaw := (analytics.PercentMined - 90.0) / 10.0
		if supplyRaw > 1.0 {
			supplyRaw = 1.0
		}
		supplyScore = int(supplyRaw * float64(maxSupplyState))
	}

	if analytics.PercentMined >= 96 {
		signals = append(signals, "Over 96% of supply mined")
	}

	// EMISSION PRESSURE (20%)
	const (
		emissionFloor   = uint64(1_000_000_000 * 1e8)
		emissionCeiling = uint64(3_000_000_000 * 1e8)
	)
	emissionScore := 0
	remaining := float64(analytics.RemainingSupply)
	if analytics.RemainingSupply <= emissionFloor {
		emissionScore = maxEmissionPressure
	} else if analytics.RemainingSupply < emissionCeiling {
		emissionRaw := 1.0 - (remaining-float64(emissionFloor))/(float64(emissionCeiling)-float64(emissionFloor))
		emissionScore = int(emissionRaw * float64(maxEmissionPressure))
	}

	if analytics.RemainingSupply < emissionFloor {
		signals = append(signals, "Remaining supply below 1B KAS")
	}

	// PROTOCOL EVENT (20%)
	protocolScore := 0
	days := analytics.DaysToReduction
	switch {
	case days <= 7:
		protocolScore = maxProtocolEvent
		signals = append(signals, "Block reward reduction within 7 days")
	case days <= 14:
		protocolScore = int(float64(maxProtocolEvent) * 0.75)
		signals = append(signals, "Block reward reduction within 14 days")
	case days <= 30:
		protocolScore = int(float64(maxProtocolEvent) * 0.50)
		signals = append(signals, "Block reward reduction within 30 days")
	case days <= 60:
		protocolScore = int(float64(maxProtocolEvent) * 0.25)
	}

	// TOTAL
	totalScore := networkScore + supplyScore + emissionScore + protocolScore
	grade, trend := gradeFromScore(totalScore)

	intel := models.Intelligence{
		Score:      totalScore,
		Grade:      grade,
		Trend:      trend,
		Signals:    signals,
		Status:     analytics.Status,
		Service:    analytics.Service,
		APIVersion: analytics.APIVersion,
	}

	intel.Components.NetworkSecurity = models.IntelligenceComponent{Score: networkScore, Max: maxNetworkSecurity}
	intel.Components.SupplyState = models.IntelligenceComponent{Score: supplyScore, Max: maxSupplyState}
	intel.Components.EmissionPressure = models.IntelligenceComponent{Score: emissionScore, Max: maxEmissionPressure}
	intel.Components.ProtocolEvent = models.IntelligenceComponent{Score: protocolScore, Max: maxProtocolEvent}

	return intel, nil
}

func gradeFromScore(score int) (string, string) {
	switch {
	case score >= 90:
		return "A+", "Strong"
	case score >= 75:
		return "A", "Healthy"
	case score >= 60:
		return "B", "Stable"
	case score >= 40:
		return "C", "Moderate"
	default:
		return "D", "Weak"
	}
}
