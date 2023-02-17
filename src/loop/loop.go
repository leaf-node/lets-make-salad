// game loop

package loop

import (

    "github.com/leaf-node/lets-make-salad/src/maps"
)


func StartLoop (seed string, gridSize int, noiseScale float64) {

    tiles := maps.GenerateMap(seed, gridSize, noiseScale)
    tiles.PrintMap()
}

