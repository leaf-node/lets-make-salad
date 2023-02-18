// starting point

package main

import (
    "os"
    "log"

    "github.com/gen2brain/raylib-go/raylib"

    "github.com/leaf-node/lets-make-salad/src/game"
    "github.com/leaf-node/lets-make-salad/src/draw"
)

var gridSize = 39
var noiseScale = 0.1


func main() {

    if len(os.Args) > 2 {
        log.Fatal("lets-make-salad takes one optional, arbitrary seed argument.\n")
    }

    var seed string
    if len(os.Args) == 2 {
        seed = os.Args[1]
    } else {
        seed = "letsmakesalad"
    }

    world := game.Init(seed, gridSize, noiseScale)

    startLoop(world)
}

func startLoop (world *game.World) {

    draw.Init()

    for !rl.WindowShouldClose() {
        game.Update(world)
        draw.Draw(world,)
    }

    rl.CloseWindow()
}

