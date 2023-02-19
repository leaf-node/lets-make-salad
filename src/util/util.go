package util

func ClampF32(value, min, max float32) float32 {
    if value < min {
        return min
    } else if value > max {
        return max
    } else {
        return value
    }
}

func Sign(value float32) float32 {
    if value > 0 {
        return 1
    } else if value < 0 {
        return -1
    } else {
        return 0
    }
}

