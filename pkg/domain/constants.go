package domain

const (
	InfluenceWealth    = "Wealth"
	InfluencePower     = "Power"
	InfluenceFame      = "Fame"
	InfluenceOpulence  = "Opulence"
	InfluenceDominance = "Dominance"
	InfluencePrestige  = "Prestige"

	InfluenceWealthID    = 1
	InfluencePowerID     = 2
	InfluenceFameID      = 3
	InfluenceOpulenceID  = 4
	InfluenceDominanceID = 5
	InfluencePrestigeID  = 6
)

var (
	influenceMap = map[string]int{
		InfluenceWealth:    InfluenceWealthID,
		InfluencePower:     InfluencePowerID,
		InfluenceFame:      InfluenceFameID,
		InfluenceOpulence:  InfluenceOpulenceID,
		InfluenceDominance: InfluenceDominanceID,
		InfluencePrestige:  InfluencePrestigeID,
	}
)

func GetInfluenceID(influenceName string) int {
	res, exist := influenceMap[influenceName]
	if !exist {
		return 0
	}

	return res
}
