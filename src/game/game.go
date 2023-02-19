// game logic

package game

import (
    "github.com/leaf-node/lets-make-salad/src/maps"
    "github.com/leaf-node/lets-make-salad/src/items"
)

type World struct {
    Tiles *maps.TileMap
    Items items.ItemsMap
}


func Init (seed string, gridSize int32, noiseScale float64) *World {

    tiles := maps.GenerateMap(seed, gridSize, noiseScale)
    items := items.New()

    world := &World{tiles, items}

    return world
}

func Update(world *World) {

    // todo

}

