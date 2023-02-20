// viewport for map

package draw

import (
    "github.com/gen2brain/raylib-go/raylib"

    "github.com/leaf-node/lets-make-salad/src/util"
)

var view viewport

type viewport struct {
    x int32
    y int32
    velX int32
    velY int32

    tileSize int32
    screenWidth int32
    screenHeight int32

    mapWidth int32
    mapHeight int32
}


func moveViewport() {

    mapPixelH := view.mapHeight * view.tileSize
    mapPixelW := view.mapWidth * view.tileSize
    topLimit := mapPixelH - view.screenHeight
    rightLimit := mapPixelW - view.screenWidth

    maxVel := view.tileSize * 3
    view.velX = util.Clamp32(view.velX, -maxVel, maxVel)
    view.velY = util.Clamp32(view.velY, -maxVel, maxVel)

    if view.x <= 0 && view.velX < 0 {
        view.velX = 0
    } else if view.x >= rightLimit && view.velX > 0 {
        view.velX = 0
    }

    if view.y <= 0 && view.velY < 0 {
        view.velY = 0
    } else if view.y >= topLimit && view.velY > 0 {
        view.velY = 0
    }

    view.x += view.velX
    view.y += view.velY

    if view.screenWidth > mapPixelW {
        view.x = -(view.screenWidth - mapPixelW) / 2
    } else {
        view.x = util.Clamp32(view.x, 0, rightLimit)
    }

    if view.screenHeight > mapPixelH {
        view.y = -(view.screenHeight - mapPixelH) / 2
    } else {
        view.y = util.Clamp32(view.y, 0, topLimit)
    }
}

func AccelerateViewport(goLeft, goRight, goUp, goDown bool) {

    accel := int32(1)

    oldSign := util.Sign32(view.velX)
    if goLeft && ! goRight && oldSign != 1 {
        view.velX -= accel
    } else if goRight && !goLeft && oldSign != -1 {
        view.velX += accel
    } else if view.velX != 0 {

        // decelerate quickly
        newVel := view.velX - oldSign * 3 * accel
        if util.Sign32(newVel) != oldSign {
            newVel = 0
        }
        view.velX = newVel
    }

    oldSign = util.Sign32(view.velY)
    if goDown && !goUp && oldSign != 1 {
        view.velY -= accel
    } else if goUp && !goDown && oldSign != -1 {
        view.velY += accel
    } else if view.velY != 0 {

        // decelerate quickly
        newVel := view.velY - oldSign * 3 * accel
        if util.Sign32(newVel) != oldSign {
            newVel = 0
        }
        view.velY = newVel
    }

}

func ResizeWindow(fullscreen bool, maximized bool) {

    isfs := rl.IsWindowFullscreen()

    if fullscreen {
        if !isfs {
            rl.ToggleFullscreen()
        }
    } else {
        if isfs { rl.ToggleFullscreen() }

        if maximized {
            rl.MaximizeWindow()
        } else {
            rl.RestoreWindow()
        }
    }
}

func handleWindowResize() {

    var newWidth int32
    var newHeight int32

    if !rl.IsWindowFullscreen() {
        newWidth = int32(rl.GetScreenWidth())
        newHeight = int32(rl.GetScreenHeight())

    } else {
        monitor := rl.GetCurrentMonitor()
        newWidth = int32(rl.GetMonitorWidth(monitor))
        newHeight = int32(rl.GetMonitorHeight(monitor))
    }

    view.y += view.screenHeight - newHeight

    view.screenWidth = newWidth
    view.screenHeight = newHeight
}

