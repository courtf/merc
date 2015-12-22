# merc
Some mercator calculations

```go
// convert between (lon, lat) and mercator (x, y)
world := merc.Earth
x, y := world.LonLatToMerc(-122.5, 45.5)

const (
    zoom = 2
    tileSize = 256
)

px, py := world.MercToPixel(x, y, zoom, tileSize)

tx, ty := world.PixelToTile(px, py, tileSize)
tx, ty = world.LonLatToTile(lon, lat, zoom, tileSize)
```
