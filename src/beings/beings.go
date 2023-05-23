
// SPDX-FileCopyrightText: 2023 Andrew Engelbrecht <andrew@sourceflow.dev>
//
// SPDX-License-Identifier: MIT
//
// beings - this is a part of lets-make-salad

package beings

import (
    "log"
    "fmt"
)

type Being struct {
    Name string
    Species string
    Gender string
    Location Coord
}

type Coord struct {
    x int32
    y int32
}

type BeingsMap map[Coord]Being


func Init() BeingsMap {

    beings := make(BeingsMap)
    return beings
}

// add being, return false if it could not be added
func (b BeingsMap) AddBeing(x int32, y int32, name string, species string, gender string) bool {

    if species != "dwarf" && species != "hobbit" {
        log.Fatal(fmt.Sprintf("invalid species: %s", species))
    }

    if b.GetBeing(x, y).Species != "" {
        return false
    }

    coord := Coord{x, y}
    being := Being{name, species, gender, coord}

    b[coord] = being

    return true
}

func (b BeingsMap) GetBeing(x int32, y int32) Being {
    return b[Coord{x, y}]
}

