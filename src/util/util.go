package util

func Clamp32(value, min, max int32) int32 {
    if value < min {
        return min
    } else if value > max {
        return max
    } else {
        return value
    }
}

func Sign32(value int32) int32 {
    if value > 0 {
        return 1
    } else if value < 0 {
        return -1
    } else {
        return 0
    }
}

