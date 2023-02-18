// game loop

package loop

import (
    "github.com/leaf-node/lets-make-salad/src/game"
)


func update(world *game.World) {

}

func draw(world *game.World) {

    world.Tiles.PrintMap()
}

func StartLoop (world *game.World) {

    update(world)
    draw(world)

}

