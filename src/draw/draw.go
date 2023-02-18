// draw scenes

package draw

import (
    "github.com/gen2brain/raylib-go/raylib"

    "github.com/leaf-node/lets-make-salad/src/game"
)

var spriteSize = int32(16)


type assets struct {
    rock rl.Texture2D
    stones rl.Texture2D
    grass rl.Texture2D
    swamp rl.Texture2D
    dirt rl.Texture2D
}

func Init() assets {

    rl.InitWindow(1280, 640, "Let's Make Salad!")
    rl.SetTargetFPS(60)

    as := assets{}
    as.load()

    return as
}

func Draw(world *game.World, as assets) {

    rl.BeginDrawing()
    rl.ClearBackground(rl.Black)

    height := int32(world.Tiles.Height)
    width := int32(world.Tiles.Width)

    tint := rl.White

    for x := int32(0); x < width ; x++ {
        for y := int32(0); y < height ; y++ {

            tile := world.Tiles.GetTile(int(x), int(y))

            x1 := x * spriteSize
            y1 := y * spriteSize

            switch tile {
            case "R":
                rl.DrawTexture(as.rock, x1, y1, tint)
            case "r":
                rl.DrawTexture(as.stones, x1, y1, tint)
            case ".":
                rl.DrawTexture(as.grass, x1, y1, tint)
            case ":":
                rl.DrawTexture(as.swamp, x1, y1, tint)
            default:
                rl.DrawTexture(as.dirt, x1, y1, tint)
            }
        }
    }

    rl.EndDrawing()
}

func (as *assets) load() {

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

