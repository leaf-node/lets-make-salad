// starting point

package main

import (
    "os"
    "log"

    "github.com/gen2brain/raylib-go/raylib"

    "github.com/leaf-node/lets-make-salad/src/game"
    "github.com/leaf-node/lets-make-salad/src/draw"
    "github.com/leaf-node/lets-make-salad/src/util"
)

var gridSize = int32(250)
var noiseScale = 0.1
var height = int32(704)
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

    world := game.Init(seed, int(gridSize), noiseScale)

    startLoop(world)
}

func startLoop (world *game.World) {

    draw.Init(width, height, gridSize, gridSize)

    for !rl.WindowShouldClose() {
        handleInput()
        game.Update(world)
        draw.Draw(world)
    }

    rl.CloseWindow()
}

func handleInput() {

    accel := float32(0.03)

    goleft := false
    goright := false
    goup := false
    godown := false

    if rl.IsKeyDown(rl.KeyLeft)  { goleft = true }
    if rl.IsKeyDown(rl.KeyRight) { goright = true }
    if rl.IsKeyDown(rl.KeyUp)    { goup = true }
    if rl.IsKeyDown(rl.KeyDown)  { godown = true }

    if goleft && ! goright {
        draw.View.VelX -= accel
    } else if goright && !goleft {
        draw.View.VelX += accel
    } else {

        oldSign := util.Sign(draw.View.VelX)
        newVel := draw.View.VelX - oldSign * 3 * accel
        if util.Sign(newVel) != oldSign {
            newVel = 0
        }
        draw.View.VelX = newVel
    }

    if godown && !goup {
        draw.View.VelY -= accel
    } else if goup && !godown {
        draw.View.VelY += accel
    } else {

        oldSign := util.Sign(draw.View.VelY)
        newVel := draw.View.VelY - oldSign * 3 * accel
        if util.Sign(newVel) != oldSign {
            newVel = 0
        }
        draw.View.VelY = newVel
    }
}
