// draw scenes

package draw

import (
    "log"
    "fmt"
    "strings"
    "path/filepath"
    "io/ioutil"

    "github.com/gen2brain/raylib-go/raylib"

    "github.com/leaf-node/lets-make-salad/src/game"
)

// 'view' global and 'viewport' type are declared in viewport.go

type texture struct {
    tex rl.Texture2D
    size int32
}

type textureMap map[string]texture

var textures textureMap


func Init(width int32, height int32, mapWidth int32, mapHeight int32) {

    rl.SetTraceLog(rl.LogWarning)

    rl.InitWindow(width, height, "Let's Make Salad!")
    rl.SetTargetFPS(60)

    rl.SetWindowState(rl.FlagWindowResizable)

    textures = textureMap{}
    textures.load()

    initView(width, height, mapWidth, mapHeight)
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

    origin := rl.Vector2{0, 0}
    rotation := float32(0)

    var tex texture
    var texName string

    for x := bottomX ; x <= topX ; x++ {
        for y := bottomY ; y <= topY ; y++ {

            tile := world.Tiles.GetTile(x, y)

            switch tile {
            case "R":
                texName = "rock"
            case "r":
                texName = "stones"
            case ".":
                texName = "grass"
            case "T":
                texName = "tree"
            case ":":
                texName = "swamp"
            case " ":
                continue
            default:
                texName = "unknownTile"
            }

            if _, ok := textures[texName] ; !ok {
                texName = "unknownTile"
            }
            tex = textures[texName]

            pixelC :=       float32(x)      * ts - float32(view.x)
            pixelR := sh - (float32(y) + 1) * ts + float32(view.y)

            dest   := rl.Rectangle{pixelC, pixelR, ts, ts}
            source := rl.Rectangle{float32(0), float32(0), float32(tex.size), float32(tex.size)}

            rl.DrawTexturePro(tex.tex, source, dest, origin, rotation, tint)


            item := world.Items.GetItem(x, y)

            switch item {
            case "w":
                texName = "wood"
            case "s":
                texName = "stoneBricks"
            case "":
                texName = "NONE"
            default:
                texName = "unknownItem"
            }

            if texName != "NONE" {
                if _, ok := textures[texName] ; !ok {
                    texName = "unknownItem"
                }
                tex = textures[texName]

                rl.DrawTexturePro(tex.tex, source, dest, origin, rotation, tint)
            }


            being := world.Beings.GetBeing(x, y)

            switch being.Species {
            case "dwarf":
                texName = "dwarf"
            case "hobbit":
                texName = "hobbit"
            case "":
                continue
            default:
                texName = "unknownBeing"
            }

            if _, ok := textures[texName] ; !ok {
                texName = "unknownBeing"
            }
            tex = textures[texName]

            rl.DrawTexturePro(tex.tex, source, dest, origin, rotation, tint)
        }
    }

    rl.EndDrawing()
}

func (t textureMap) load() {

    spritesDir := "img/16x16/"

    files, err := ioutil.ReadDir(spritesDir)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        name := file.Name()

        fullPath := fmt.Sprintf("%s%s", spritesDir, name)
        shortName := strings.TrimSuffix(name, filepath.Ext(name))

        tex := rl.LoadTexture(fullPath)
        t[shortName] = texture{tex, 16}
    }

    if _, ok := t["unknownTile"] ; !ok  {
        log.Fatal("Missing 'unknownTile.png' image")
    }
    if _, ok := t["unknownItem"] ; !ok  {
        log.Fatal("Missing 'unknownItem.png' image")
    }
}

func (t textureMap) unload() {

    for name, _ := range t {
        rl.UnloadTexture(t[name].tex)
        delete(t, name)
    }
}

