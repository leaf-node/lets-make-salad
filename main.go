// starting point

package main

import (
    "os"
    "log"

    "github.com/gen2brain/raylib-go/raylib"

    "github.com/leaf-node/lets-make-salad/src/game"
    "github.com/leaf-node/lets-make-salad/src/draw"
)

var gridSize = 250
var noiseScale = 0.1
var height = int32(720)
var width = int32(1280)

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

    draw.Init(width, height)

    for !rl.WindowShouldClose() {
        handleInput()
        game.Update(world)
        draw.Draw(world)
    }

    rl.CloseWindow()
}

func handleInput() {

    stepSize := int32(2)

    if (rl.IsKeyDown(rl.KeyDown)) {
        draw.View.Y -= stepSize
    }
    if (rl.IsKeyDown(rl.KeyUp)) {
        draw.View.Y += stepSize
    }
    if (rl.IsKeyDown(rl.KeyLeft)) {
        draw.View.X -= stepSize
    }
    if (rl.IsKeyDown(rl.KeyRight)) {
        draw.View.X += stepSize
    }
}
