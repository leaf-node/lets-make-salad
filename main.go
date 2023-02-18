// starting point

package main

import (
    "os"
    "log"

    "github.com/leaf-node/lets-make-salad/src/game"
    "github.com/leaf-node/lets-make-salad/src/draw"

)

var gridSize = 39
var noiseScale = 0.1


func main() {

    if len(os.Args) != 2 {
        log.Fatal("lets-make-salad takes one arbitrary seed argument.\n")
    }

    world := game.Init(os.Args[1], gridSize, noiseScale)

    startLoop(world)
}

func startLoop (world *game.World) {

    game.Update(world)
    draw.Draw(world)

}

