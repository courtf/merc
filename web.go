package merc

import "math"

/* Web Mercator Specific Calcs */

// Thanks: http://www.maptiler.org/google-maps-coordinates-tile-bounds-projection/

// meters per pixel, at the provided zoom and tile size
func (w World) Resolution(zoom uint, tileSize int) float64 {
	return w.Circumference / float64(tileSize*(1<<zoom))
}

func (w World) MercToPixels(x, y float64, zoom uint, tileSize int) (px, py float64) {
	metersPerPixel := w.Resolution(zoom, tileSize)
	shift := w.Circumference / 2

	px = (x + shift) / metersPerPixel
	py = (y + shift) / metersPerPixel
	return
}

func (w World) MercToTile(x, y float64, zoom uint, tileSize int) (tx, ty int) {
	px, py := w.MercToPixels(x, y, zoom, tileSize)

	tx = int(math.Ceil(px/float64(tileSize)) - 1)
	ty = int(math.Ceil(py/float64(tileSize)) - 1)
	return
}
