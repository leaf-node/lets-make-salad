// draw scenes

package draw

import (
    "github.com/gen2brain/raylib-go/raylib"

    "github.com/leaf-node/lets-make-salad/src/game"
)

var as assets
var view viewport

type viewport struct {
    x int32
    y int32
    gridSize int32
}

type assets struct {
    size int32
    rock rl.Texture2D
    stones rl.Texture2D
    grass rl.Texture2D
    swamp rl.Texture2D
    dirt rl.Texture2D
}

func Init() {

    rl.InitWindow(1280, 640, "Let's Make Salad!")
    rl.SetTargetFPS(60)

    as = assets{}
    as.load()

    view.gridSize = 32
}

func Draw(world *game.World) {

    rl.BeginDrawing()
    rl.ClearBackground(rl.Black)

    height := int32(world.Tiles.Height)
    width := int32(world.Tiles.Width)

    tint := rl.White

    for x := int32(0); x < width ; x++ {
        for y := int32(0); y < height ; y++ {

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
            default:
                tex = as.dirt
            }

            gs := view.gridSize

            source := rl.Rectangle{float32(0), float32(0), float32(as.size), float32(as.size)}
            dest := rl.Rectangle{float32(x * gs), float32(y * gs), float32(gs), float32(gs)}

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

