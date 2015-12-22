package merc

import "math"

// Definiton: y'(θ)= radius * secθ
// radius * ln[tan(pi/2 + θ/2)]
// This is expounded here:
// https://en.wikipedia.org/wiki/Mercator_projection#Derivation_of_the_Mercator_projection
func (w World) LatitudeToMercator(lat float64) float64 {
	return w.Radius * math.Log(math.Tan(math.Pi/4+DegToRad(lat)/2))
}

// Inverse of the LatitudeToMercator.
func (w World) MercatorToLatitude(y float64) float64 {
	return RadToDeg(2*math.Atan(math.Exp(y/w.Radius)) - (math.Pi / 2))
}
