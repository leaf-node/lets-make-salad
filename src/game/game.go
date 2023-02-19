// game logic

package game

import (
    "github.com/leaf-node/lets-make-salad/src/maps"
)

type World struct {
    Tiles *maps.TileMap
}


func Init (seed string, gridSize int32, noiseScale float64) *World {

    tiles := maps.GenerateMap(seed, gridSize, noiseScale)
    world := &World{tiles}

    return world
}

func Update(world *World) {

    // todo

}

