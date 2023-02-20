// draw scenes

package draw

import (
    "github.com/gen2brain/raylib-go/raylib"

    "github.com/leaf-node/lets-make-salad/src/game"
)

// 'view' global and 'viewport' type are declared in viewport.go




type assets struct {
    size int32
    rock rl.Texture2D
    stones rl.Texture2D
    grass rl.Texture2D
    swamp rl.Texture2D
    dirt rl.Texture2D
    tree rl.Texture2D
    wood rl.Texture2D
    stoneBricks rl.Texture2D
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
    topX := view.screenWidth / view.tileSize + bottomX + 1
    topY := view.screenHeight / view.tileSize + bottomY + 1

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
            case "T":
                tex = as.tree
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


            item := world.Items.GetItem(x, y)

            switch item {
            case "w":
                tex = as.wood
            case "s":
                tex = as.stoneBricks
            default:
                continue
            }

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
    as.tree = rl.LoadTexture("img/tree.png")
    as.wood = rl.LoadTexture("img/wood.png")
    as.stoneBricks = rl.LoadTexture("img/stoneBricks.png")
}

func (as *assets) unload() {

    rl.UnloadTexture(as.rock)
    rl.UnloadTexture(as.stones)
    rl.UnloadTexture(as.grass)
    rl.UnloadTexture(as.swamp)
    rl.UnloadTexture(as.dirt)
    rl.UnloadTexture(as.tree)
    rl.UnloadTexture(as.wood)
    rl.UnloadTexture(as.stoneBricks)
}

