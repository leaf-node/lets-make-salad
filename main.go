
package main

import (
    "os"
    "fmt"

    "github.com/leaf-node/lets-make-salad/maps"
)

var gridSize = 79

func main() {

    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "lets-make-salad takes one arbitrary seed argument.\n")
    } else {
        maps.GenerateMap(os.Args[1], gridSize)
    }
}

