// game logic

package game

import (
    "math/rand"

    "github.com/leaf-node/lets-make-salad/src/maps"
    "github.com/leaf-node/lets-make-salad/src/items"
    "github.com/leaf-node/lets-make-salad/src/beings"
)

type World struct {
    Tiles *maps.TileMap
    Beings beings.BeingsMap
    Items items.ItemsMap
}


func Init (seed string, gridSize int32, noiseScale float64) *World {

    tiles := maps.GenerateMap(seed, gridSize, noiseScale)
    beings := beings.Init()
    items := items.New()

    world := &World{tiles, beings, items}

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

    if world.Tiles.GetTile(x, y) == "." {
        if world.Items.GetItem(x, y) == "" {
            world.Items.AddItem(item, x, y)
        }
    }


    var species string

    x = int32(rand.Intn(int(world.Tiles.Width)))
    y = int32(rand.Intn(int(world.Tiles.Height)))

    if rand.Intn(2) == 1 {
        species = "dwarf"
    } else {
        species = "hobbit"
    }

    if world.Tiles.GetTile(x, y) == "." {
        world.Beings.AddBeing(x, y, "Jo", species, "other")
    }

    // todo

}

