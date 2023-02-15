
package main

import (
    "os"
    "fmt"

    "github.com/leaf-node/lets-make-salad/maps"
)

var gridSize = 39
var noiseScale = 0.1

func main() {

    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "lets-make-salad takes one arbitrary seed argument.\n")

    } else {
        tiles := maps.GenerateMap(os.Args[1], gridSize, noiseScale)
        tiles.PrintMap()
    }
}

