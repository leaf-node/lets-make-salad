
// SPDX-FileCopyrightText: 2023 Andrew Engelbrecht <andrew@sourceflow.dev>
//
// SPDX-License-Identifier: MIT
//
// items rules - this is a part of lets-make-salad

package items

type coord struct {
    x int32
    y int32
}

type ItemsMap map[coord]string


func New() ItemsMap {

    return make(ItemsMap)
}

// add item, return false if it could not be added
func (i ItemsMap) AddItem(item string, x int32, y int32) bool {

    if i.GetItem(x, y) != "" {
        return false
    }
    i[coord{x, y}] = item

    return true
}

func (i ItemsMap) GetItem(x int32, y int32) string {

    return i[coord{x, y}]
}

func (i ItemsMap) RemoveItem(item string, x int32, y int32) {

    delete(i, coord{x, y})
}

