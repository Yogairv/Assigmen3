package helpers

import "math/rand"

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func CheckStatusWater(value int) string {
	var status string
	if value < 5 {
		status = "baik"
	} else if (value >= 6) && (value <= 8) {
		status = "siaga"
	} else if value > 8 {
		status = "buruk"
	}

	return status
}

func CheckStatusWind(value int) string {
	var status string
	if value <= 6 {
		status = "baik"
	} else if (value >= 7) && (value <= 15) {
		status = "siaga"
	} else if value > 15 {
		status = "buruk"
	}

	return status
}
