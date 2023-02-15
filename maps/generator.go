// map generator

package maps

import (
    "fmt"
    "math"
    "crypto/sha1"
    "encoding/binary"

    "github.com/ojrac/opensimplex-go"
)


type tileMap struct {
    array []byte
    width int
    height int
    noiseScale float64
    noise opensimplex.Noise
}


func newEmptyMap(width int, height int, noiseScale float64) tileMap {
    array := make([]byte, width * height)
    return tileMap{array, width, height, noiseScale, nil}
}

func GenerateMap(seed string, gridSize int, noiseScale float64) tileMap {

    hash := sha1.Sum([]byte(seed))
    seedInt := int64(binary.BigEndian.Uint64(hash[12:]))

    tiles := newEmptyMap(gridSize, gridSize, noiseScale)
    tiles.noise = opensimplex.New(seedInt)

    for y := 0; y < tiles.height; y++ {
        for x := 0; x < tiles.width; x++ {
            tiles.generateTile(x, y)
        }
    }

    return tiles
}

func (t tileMap) PrintMap() {
    for y := t.height - 1; y >= 0; y-- {
        for x := 0; x < t.width; x++ {

            tile := t.getTile(x, y)
            print(tile, tile)
        }
        fmt.Println()
    }
}

func (t tileMap) generateTile(x int, y int) {

    var tileStr string

    height := math.Sqrt(t.sampleNoise(x, y, t.noiseScale))
    height *= math.Sqrt(t.sampleNoise(-x, -y, t.noiseScale * 3))

    if height <= 0.35 {
        tileStr = ":" // swamp
    } else if height < 0.60 {
        tileStr = "." // land
    } else if height < 0.65 {
        tileStr = "r" // rocky land
    } else {
        tileStr = "R" // rock
    }

    t.setTile(x, y, []byte(tileStr)[0])
}

func (t tileMap) sampleNoise(x int, y int, scale float64) float64 {

    value := t.noise.Eval2(float64(x) * scale, float64(y) * scale)
    normalized := (value + 1) / 2
    return normalized
}

func (t tileMap) setTile(x int, y int, tile byte) {
    t.array[y* t.width + x] = tile
}

func (t tileMap) getTile(x int, y int) string {
    return string(t.array[y* t.width + x])
}

