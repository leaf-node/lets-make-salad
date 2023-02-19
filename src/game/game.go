// game logic

package game

import (
    "math/rand"

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

func (world *World) Update() {

    var item string

    x := int32(rand.Intn(int(world.Tiles.Width)))
    y := int32(rand.Intn(int(world.Tiles.Height)))

    if rand.Intn(2) == 1 {
        item = "w"
    } else {
        item = "s"
    }

    switch world.Tiles.GetTile(x, y) {
    case ".":
        if world.Items.GetItem(x, y) == "" {
            world.Items.AddItem(item, x, y)
        }
    }

    // todo

}

