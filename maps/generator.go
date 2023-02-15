// map generator

package maps

import (
    "fmt"
    "math"
    "crypto/sha1"
    "encoding/binary"

    "github.com/ojrac/opensimplex-go"
)

var gridScale float64 = 0.1
var noise opensimplex.Noise

func GenerateMap(seed string, gridSize int) {

    hash := sha1.Sum([]byte(seed))
    seedInt := int64(binary.BigEndian.Uint64(hash[12:]))

    noise = opensimplex.New(seedInt)

    for i := 0; i < gridSize; i++ {
        for j := 0; j < gridSize; j++ {

            tile := generateTile(j, i, gridScale)
            print(tile, tile)
        }
        fmt.Println()
    }
}

func generateTile(j int, i int, scale float64) string {

    height := math.Sqrt(sampleNoise(j, i, scale))
    height *= math.Sqrt(sampleNoise(-j, -i, scale * 3))

    if height <= 0.35 {
        return ":" // swamp
    } else if height < 0.65 {
        return "." // land
    } else {
        return "M" // rock
    }
}

func sampleNoise(i int, j int, scale float64) float64 {

    value := noise.Eval2(float64(j) * scale, float64(i) * scale)
    normalized := (value + 1) / 2
    return normalized
}

