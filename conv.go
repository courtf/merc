package merc

import "math"

func DegToRad(n float64) float64 {
	return n * math.Pi / 180
}

func RadToDeg(n float64) float64 {
	return n * 180 / math.Pi
}
