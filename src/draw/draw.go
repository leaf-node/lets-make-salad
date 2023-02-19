// draw scenes

package draw

import (
    "github.com/gen2brain/raylib-go/raylib"

    "github.com/leaf-node/lets-make-salad/src/game"
    "github.com/leaf-node/lets-make-salad/src/util"
)

var as assets
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

type assets struct {
    size int32
    rock rl.Texture2D
    stones rl.Texture2D
    grass rl.Texture2D
    swamp rl.Texture2D
    dirt rl.Texture2D
}

func Init(width int32, height int32, mapWidth int32, mapHeight int32) {

    rl.SetTraceLog(rl.LogWarning)

    rl.InitWindow(width, height, "Let's Make Salad!")
    rl.SetTargetFPS(60)

    rl.SetWindowState(rl.FlagWindowResizable)

    as = assets{}
    as.load()

    view.tileSize = 32
    view.screenWidth = width
    view.screenHeight = height
    view.mapWidth = mapWidth
    view.mapHeight = mapHeight

    view.x = mapWidth * view.tileSize / 2
    view.y = mapHeight * view.tileSize / 2
}

func Draw(world *game.World) {

    handleWindowResize()
    moveViewport()

    tint := rl.White

    rl.BeginDrawing()
    rl.ClearBackground(rl.Black)

    ts := float32(view.tileSize)
    sh := float32(view.screenHeight)

    bottomX := view.x / view.tileSize
    bottomY := view.y / view.tileSize
    topX := view.screenWidth / view.tileSize + bottomX
    topY := view.screenHeight / view.tileSize + bottomY

    source := rl.Rectangle{float32(0), float32(0), float32(as.size), float32(as.size)}
    origin := rl.Vector2{0, 0}
    rotation := float32(0)

    var tex rl.Texture2D

    for x := bottomX ; x <= topX ; x++ {
        for y := bottomY ; y <= topY ; y++ {

            tile := world.Tiles.GetTile(x, y)

            switch tile {
            case "R":
                tex = as.rock
            case "r":
                tex = as.stones
            case ".":
                tex = as.grass
            case ":":
                tex = as.swamp
            case " ":
                continue
            default:
                tex = as.dirt
            }

            pixelC :=       float32(x)      * ts - float32(view.x)
            pixelR := sh - (float32(y) + 1) * ts + float32(view.y)

            dest := rl.Rectangle{pixelC, pixelR, ts, ts}

            rl.DrawTexturePro(tex, source, dest, origin, rotation, tint)
        }
    }

    rl.EndDrawing()
}

func (as *assets) load() {

    as.size = 16

    as.rock = rl.LoadTexture("img/rock.png")
    as.stones = rl.LoadTexture("img/stones.png")
    as.grass = rl.LoadTexture("img/grass.png")
    as.swamp = rl.LoadTexture("img/swamp.png")
    as.dirt = rl.LoadTexture("img/dirt.png")
}

func (as *assets) unload() {

    rl.UnloadTexture(as.rock)
    rl.UnloadTexture(as.stones)
    rl.UnloadTexture(as.grass)
    rl.UnloadTexture(as.swamp)
    rl.UnloadTexture(as.dirt)
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

