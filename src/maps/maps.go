
// SPDX-FileCopyrightText: 2023 Andrew Engelbrecht <andrew@sourceflow.dev>
//
// SPDX-License-Identifier: MIT
//
// map generator - this is a part of lets-make-salad

package maps

import (
    "fmt"
    "crypto/sha1"
    "encoding/binary"

    "github.com/ojrac/opensimplex-go"
)


type TileMap struct {
    array []byte
    Width int32
    Height int32
    noiseScale float64
    noise opensimplex.Noise
}


func newEmptyMap(width int32, height int32, noiseScale float64) *TileMap {
    array := make([]byte, width * height)
    return &TileMap{array, width, height, noiseScale, nil}
}

func GenerateMap(seed string, gridSize int32, noiseScale float64) *TileMap {

    hash := sha1.Sum([]byte(seed))
    seedInt := int64(binary.BigEndian.Uint64(hash[12:]))

    tiles := newEmptyMap(gridSize, gridSize, noiseScale)
    tiles.noise = opensimplex.New(seedInt)

    for y := int32(0); y < tiles.Height; y++ {
        for x := int32(0); x < tiles.Width; x++ {
            tiles.generateTile(x, y)
        }
    }

    return tiles
}

func (t TileMap) DebugPrintMap() {
    for y := t.Height - 1; y >= 0; y-- {
        for x := int32(0); x < t.Width; x++ {

            tile := t.GetTile(x, y)
            print(tile, tile)
        }
        fmt.Println()
    }
}

func (t TileMap) generateTile(x, y int32) {

    var tileStr string

    height := t.sampleNoise(x, y, t.noiseScale)
    height *= t.sampleNoise(-x, -y, t.noiseScale * 3)

    if height <= 0.05 {
        tileStr = ":" // swamp
    } else if height <= 0.12 {
        tileStr = "T" // tree
    } else if height < 0.36 {
        tileStr = "." // land
    } else if height < 0.42 {
        tileStr = "r" // rocky land
    } else {
        tileStr = "R" // rock
    }

    t.setTile(x, y, []byte(tileStr)[0])
}

func (t TileMap) sampleNoise(x int32, y int32, scale float64) float64 {

    value := t.noise.Eval2(float64(x) * scale, float64(y) * scale)
    normalized := (value + 1) / 2
    return normalized
}

func (t TileMap) setTile(x int32, y int32, tile byte) {
    t.array[y* t.Width + x] = tile
}

func (t TileMap) GetTile(x int32, y int32) string {
    if x < 0 || y < 0 || x >= t.Width || y >= t.Height {
        return " "
    }
    return string(t.array[y* t.Width + x])
}

