package merc

// This calculation is an arc length.
// Definition: θ = arcLen / radius
// arcLen = θ * radius
// So, we first convert to a radian measure of
// our longitude angle.
func (w World) LongitudeToMercator(lon float64) float64 {
	return DegToRad(lon) * w.Radius
}

// Inverse of LongitudeToMercator.
func (w World) MercatorToLongitude(x float64) float64 {
	return RadToDeg(x / w.Radius)
}
