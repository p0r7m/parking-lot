package payment

func CalculateCost(costType string, startDuration int64, endDuration int64, rateCard HourlyRateCard) int64 {
	if costType != "hourly" {
		return -1
	}

	duration := endDuration - startDuration

	fare := int64(0)

	for _, card := range rateCard.cards {
		if duration >= card.minDuration && duration <= card.maxDuration {
			return card.cost
		}
		if card.cost > fare {
			fare = card.cost
		}
	}
	return fare
}