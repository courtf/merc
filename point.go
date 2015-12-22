package merc

func (w World) LonLatToMerc(lon, lat float64) (x, y float64) {
	x = w.LongitudeToMercator(lon)
	y = w.LatitudeToMercator(lat)
	return
}

func (w World) MercToLonLat(x, y float64) (lon, lat float64) {
	lon = w.MercatorToLongitude(x)
	lat = w.MercatorToLatitude(y)
	return
}
