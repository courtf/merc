package merc

import "math"

const (
	earthRadius        = 6378137
	earthCircumference = 2 * earthRadius * math.Pi
)

type World struct {
	Radius        float64
	Circumference float64
}

var Earth = World{
	Radius:        earthRadius,
	Circumference: earthCircumference,
}

func NewWorld(circumference float64) World {
	return World{
		Radius:        circumference / (2 * math.Pi),
		Circumference: circumference,
	}
}

func NewWorldWithRadius(radius float64) World {
	return World{
		Radius:        radius,
		Circumference: 2 * math.Pi * radius,
	}
}
