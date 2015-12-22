package merc

import "math"

/* Web Mercator Specific Calcs */

// Thanks: http://www.maptiler.org/google-maps-coordinates-tile-bounds-projection/

// funcs here produce coordinates with the origin in the top letf, so pixel and tile
// coordinates are in OSM/Google format, not TMS.

func (w World) PixelCircumference(zoom uint, tileSize int) int {
	return tileSize * (1 << zoom)
}

// meters per pixel, at the provided zoom and tile size
func (w World) Resolution(zoom uint, tileSize int) float64 {
	return w.Circumference / float64(w.PixelCircumference(zoom, tileSize))
}

func (w World) MercToPixel(x, y float64, zoom uint, tileSize int) (px, py float64) {
	metersPerPixel := w.Resolution(zoom, tileSize)

	// this shift moves the cordinate system to [0, Circumference]
	// in both dimensions
	shift := w.Circumference / 2

	px = (x + shift) / metersPerPixel
	py = (y + shift) / metersPerPixel

	// move origin to top left
	py = float64(w.PixelCircumference(zoom, tileSize)) - py

	return
}

func (w World) MercToTile(x, y float64, zoom uint, tileSize int) (tx, ty int) {
	px, py := w.MercToPixel(x, y, zoom, tileSize)
	tx, ty = w.PixelToTile(px, py, tileSize)
	return
}

func (w World) PixelToTile(px, py float64, tileSize int) (tx, ty int) {
	tx = int(math.Floor(px / float64(tileSize)))
	ty = int(math.Floor(py / float64(tileSize)))
	return
}

func (w World) LonLatToPixel(lon, lat float64, zoom uint, tileSize int) (px, py float64) {
	x, y := w.LonLatToMerc(lon, lat)
	return w.MercToPixel(x, y, zoom, tileSize)
}

func (w World) LonLatToTile(lon, lat float64, zoom uint, tileSize int) (tx, ty int) {
	px, py := w.LonLatToPixel(lon, lat, zoom, tileSize)
	return w.PixelToTile(px, py, tileSize)
}
