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
    GridSize int32
    ScreenWidth int32
    ScreenHeight int32
}

type assets struct {
    size int32
    rock rl.Texture2D
    stones rl.Texture2D
    grass rl.Texture2D
    swamp rl.Texture2D
    dirt rl.Texture2D
}

func Init(width int32, height int32) {

    rl.InitWindow(width, height, "Let's Make Salad!")
    rl.SetTargetFPS(60)

    as = assets{}
    as.load()

    View.GridSize = 32
    View.ScreenWidth = width
    View.ScreenHeight = height
}

func Draw(world *game.World) {

    rl.BeginDrawing()
    rl.ClearBackground(rl.Black)

    tint := rl.White

    for x := View.X ; x <= View.X + View.ScreenWidth / View.GridSize ; x++ {
        for y := View.Y ; y <= View.Y + View.ScreenHeight / View.GridSize ; y++ {

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

            gs := View.GridSize

            source := rl.Rectangle{float32(0), float32(0), float32(as.size), float32(as.size)}
            dest := rl.Rectangle{float32(x * gs), float32(View.ScreenHeight - ((y + 1) * gs)), float32(gs), float32(gs)}

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

