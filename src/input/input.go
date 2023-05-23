
// SPDX-FileCopyrightText: 2023 Andrew Engelbrecht <andrew@sourceflow.dev>
//
// SPDX-License-Identifier: MIT
//
// input handling - this is a part of lets-make-salad

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

    if rl.IsKeyPressed(rl.KeyMinus) {
        draw.Zoom("out")
    }
    if rl.IsKeyPressed(rl.KeyEqual) {
        draw.Zoom("in")
    }
}

