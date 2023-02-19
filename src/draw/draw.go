// draw scenes

package draw

import (
    "github.com/gen2brain/raylib-go/raylib"

    "github.com/leaf-node/lets-make-salad/src/game"
)

var as assets
var View viewport

type viewport struct {
    X int32
    Y int32
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

    as = assets{}
    as.load()

    View.TileSize = 32
    View.ScreenWidth = width
    View.ScreenHeight = height
    View.mapWidth = mapWidth
    View.mapHeight = mapHeight

    View.X = mapWidth / 2
    View.Y = mapHeight / 2
}

func Draw(world *game.World) {

    clampViewport()

    tint := rl.White

    rl.BeginDrawing()
    rl.ClearBackground(rl.Black)

    for x := View.X ; x <= View.ScreenWidth / View.TileSize + View.X ; x++ {
        for y := View.Y ; y <= View.ScreenHeight / View.TileSize + View.Y ; y++ {

            var tex rl.Texture2D

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

            gs := View.TileSize

            source := rl.Rectangle{float32(0), float32(0), float32(as.size), float32(as.size)}
            dest := rl.Rectangle{float32((x - View.X) * gs), float32(View.ScreenHeight - ((y - View.Y + 1) * gs)), float32(gs), float32(gs)}

            origin := rl.Vector2{0, 0}
            rotation := float32(0)

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

func clampViewport() {

    topmostPos := View.mapHeight - View.ScreenHeight / View.TileSize
    rightmostPos := View.mapWidth - View.ScreenWidth / View.TileSize

    if View.X < 0 { View.X = 0 }
    if View.Y < 0 { View.Y = 0 }

    if View.X > rightmostPos { View.X = rightmostPos }
    if View.Y > topmostPos { View.Y = topmostPos }
}
