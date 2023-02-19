// items rules

package items

type coord struct {
    x int32
    y int32
}

type ItemsMap map[coord]string


func New() ItemsMap {

    return make(ItemsMap)
}

func (i ItemsMap) AddItem(item string, x int32, y int32) {

    i[coord{x, y}] = item
}

func (i ItemsMap) GetItem(x int32, y int32) string {

    return i[coord{x, y}]
}

func (i ItemsMap) RemoveItem(item string, x int32, y int32) {

    delete(i, coord{x, y})
}

