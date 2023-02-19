// draw scenes

package draw

import (
    "github.com/gen2brain/raylib-go/raylib"

    "github.com/leaf-node/lets-make-salad/src/game"
    "github.com/leaf-node/lets-make-salad/src/util"
)

var as assets
var View viewport

type viewport struct {
    X int32
    Y int32
    VelX int32
    VelY int32

    TileSize int32
    ScreenWidth int32
    ScreenHeight int32

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

    View.TileSize = 32
    View.ScreenWidth = width
    View.ScreenHeight = height
    View.mapWidth = mapWidth
    View.mapHeight = mapHeight

    View.X = mapWidth * View.TileSize / 2
    View.Y = mapHeight * View.TileSize / 2
}

func Draw(world *game.World) {

    handleWindowResize()
    moveViewport()

    tint := rl.White

    rl.BeginDrawing()
    rl.ClearBackground(rl.Black)

    ts := float32(View.TileSize)
    sh := float32(View.ScreenHeight)

    bottomX := View.X / View.TileSize
    bottomY := View.Y / View.TileSize
    topX := View.ScreenWidth / View.TileSize + bottomX
    topY := View.ScreenHeight / View.TileSize + bottomY

    source := rl.Rectangle{float32(0), float32(0), float32(as.size), float32(as.size)}
    origin := rl.Vector2{0, 0}
    rotation := float32(0)

    var tex rl.Texture2D

    for x := bottomX ; x <= topX ; x++ {
        for y := bottomY ; y <= topY ; y++ {

            tile := world.Tiles.GetTile(int(x), int(y))

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

            pixelC :=       float32(x)      * ts - float32(View.X)
            pixelR := sh - (float32(y) + 1) * ts + float32(View.Y)

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

    topLimit := View.mapHeight * View.TileSize - View.ScreenHeight
    rightLimit := View.mapWidth * View.TileSize - View.ScreenWidth

    maxVel := View.TileSize * 3

    View.VelX = util.Clamp32(View.VelX, -maxVel, maxVel)
    View.VelY = util.Clamp32(View.VelY, -maxVel, maxVel)

    if View.X <= 0 && View.VelX < 0 {
        View.VelX = 0
    } else if View.X >= rightLimit && View.VelX > 0 {
        View.VelX = 0
    }

    if View.Y <= 0 && View.VelY < 0 {
        View.VelY = 0
    } else if View.Y >= topLimit && View.VelY > 0 {
        View.VelY = 0
    }

    View.X += View.VelX
    View.Y += View.VelY

    View.X = util.Clamp32(View.X, 0, rightLimit)
    View.Y = util.Clamp32(View.Y, 0, topLimit)
}

func AccelerateViewport(goLeft, goRight, goUp, goDown bool) {

    accel := int32(1)

    oldSign := util.Sign32(View.VelX)
    if goLeft && ! goRight && oldSign != 1 {
        View.VelX -= accel
    } else if goRight && !goLeft && oldSign != -1 {
        View.VelX += accel
    } else if View.VelX != 0 {

        // decelerate quickly
        newVel := View.VelX - oldSign * 3 * accel
        if util.Sign32(newVel) != oldSign {
            newVel = 0
        }
        View.VelX = newVel
    }

    oldSign = util.Sign32(View.VelY)
    if goDown && !goUp && oldSign != 1 {
        View.VelY -= accel
    } else if goUp && !goDown && oldSign != -1 {
        View.VelY += accel
    } else if View.VelY != 0 {

        // decelerate quickly
        newVel := View.VelY - oldSign * 3 * accel
        if util.Sign32(newVel) != oldSign {
            newVel = 0
        }
        View.VelY = newVel
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

    View.Y += View.ScreenHeight - newHeight

    View.ScreenWidth = newWidth
    View.ScreenHeight = newHeight
}

