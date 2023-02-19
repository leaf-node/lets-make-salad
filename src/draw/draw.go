// draw scenes

package draw

import (
    "math"

    "github.com/gen2brain/raylib-go/raylib"

    "github.com/leaf-node/lets-make-salad/src/game"
    "github.com/leaf-node/lets-make-salad/src/util"
)

var as assets
var View viewport

type viewport struct {
    X float32
    Y float32
    VelX float32
    VelY float32

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

    View.X = float32(mapWidth) / 2
    View.Y = float32(mapHeight) / 2
}

func Draw(world *game.World) {

    handleWindowResize()
    moveViewport()

    tint := rl.White

    rl.BeginDrawing()
    rl.ClearBackground(rl.Black)

    ts := float32(View.TileSize)

    bottomX := int32(math.Floor(float64(View.X)))
    bottomY := int32(math.Floor(float64(View.Y)))
    topX := View.ScreenWidth / View.TileSize + bottomX + 1
    topY := View.ScreenHeight / View.TileSize + bottomY + 1

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

            xf := float32(x)
            yf := float32(y)

            dest := rl.Rectangle{((xf - View.X) * ts), (float32(View.ScreenHeight) - (yf - View.Y + 1) * ts), ts, ts}

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

    topmost := float32(View.mapHeight) - float32(View.ScreenHeight) / float32(View.TileSize)
    rightmost := float32(View.mapWidth) - float32(View.ScreenWidth) / float32(View.TileSize)

    View.VelX = util.ClampF32(View.VelX, -3, 3)
    View.VelY = util.ClampF32(View.VelY, -3, 3)

    if View.X <= 0 && View.VelX < 0 {
        View.VelX = 0
    } else if View.X >= rightmost && View.VelX > 0 {
        View.VelX = 0
    }

    if View.Y <= 0 && View.VelY < 0 {
        View.VelY = 0
    } else if View.Y >= topmost && View.VelY > 0 {
        View.VelY = 0
    }

    View.X += View.VelX
    View.Y += View.VelY

    View.X = util.ClampF32(View.X, 0, rightmost)
    View.Y = util.ClampF32(View.Y, 0, topmost)
}

func AccelerateViewport(goLeft, goRight, goUp, goDown bool) {

    accel := float32(0.03)

    if goLeft && ! goRight && util.Sign(View.VelX) != 1 {
        View.VelX -= accel
    } else if goRight && !goLeft && util.Sign(View.VelX) != -1 {
        View.VelX += accel
    } else if View.VelX != 0 {

        // decelerate quickly
        oldSign := util.Sign(View.VelX)
        newVel := View.VelX - oldSign * 3 * accel
        if util.Sign(newVel) != oldSign {
            newVel = 0
        }
        View.VelX = newVel
    }

    if goDown && !goUp && util.Sign(View.VelY) != 1 {
        View.VelY -= accel
    } else if goUp && !goDown && util.Sign(View.VelY) != -1 {
        View.VelY += accel
    } else if View.VelY != 0 {

        // decelerate quickly
        oldSign := util.Sign(View.VelY)
        newVel := View.VelY - oldSign * 3 * accel
        if util.Sign(newVel) != oldSign {
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

    if rl.IsWindowResized() {

        View.ScreenWidth = int32(rl.GetScreenWidth())
        View.ScreenHeight = int32(rl.GetScreenHeight())

    } else if rl.IsWindowFullscreen() {

        monitor := rl.GetCurrentMonitor()
        View.ScreenWidth = int32(rl.GetMonitorWidth(monitor))
        View.ScreenHeight = int32(rl.GetMonitorHeight(monitor))
    }
}

