// input handling

package input

import (
    "github.com/gen2brain/raylib-go/raylib"

    "github.com/leaf-node/lets-make-salad/src/draw"
)

func HandleInput () {

    // scroll viewport
    goLeft := false
    goRight := false
    goUp := false
    goDown := false

    if rl.IsKeyDown(rl.KeyLeft)  { goLeft = true }
    if rl.IsKeyDown(rl.KeyRight) { goRight = true }
    if rl.IsKeyDown(rl.KeyUp)    { goUp = true }
    if rl.IsKeyDown(rl.KeyDown)  { goDown = true }

    draw.AccelerateViewport(goLeft, goRight, goUp, goDown)

    // resize window
    if rl.IsKeyPressed(rl.KeyF) {
        if rl.IsWindowFullscreen() {
            draw.ResizeWindow(false, true)
        } else {
            draw.ResizeWindow(true, false)
        }
    }
    if rl.IsKeyPressed(rl.KeyM) {
        if !rl.IsWindowMaximized() {
            draw.ResizeWindow(false, true)
        } else {
            draw.ResizeWindow(false, false)
        }
    }
}

