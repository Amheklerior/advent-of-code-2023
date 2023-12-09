package day6

type Race struct {
	totTime, record int
}

func getWaysToWin(race Race) int {
	return findHighestChargeTime(race.totTime, race.record) - findLowestChargeTime(race.totTime, race.record) + 1
}

func findLowestChargeTime(totTime, record int) int {
	chargeTime := 1
	distance := simulateDistance(totTime, chargeTime)

	for distance <= record {
		chargeTime++
		distance = simulateDistance(totTime, chargeTime)
	}

	return chargeTime
}

func findHighestChargeTime(totTime, record int) int {
	chargeTime := totTime - 1
	distance := simulateDistance(totTime, chargeTime)

	for distance <= record {
		chargeTime--
		distance = simulateDistance(totTime, chargeTime)
	}

	return chargeTime
}

func simulateDistance(totTime, chargeTime int) int {
	const VELOCITY_CHARGE_PER_MS = 1
	v := VELOCITY_CHARGE_PER_MS * chargeTime
	d := v * (totTime - chargeTime)
	return d
}
