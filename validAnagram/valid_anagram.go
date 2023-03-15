func isAnagram(s string, t string) bool {
    d := make(map[rune]int)
    for _, character := range s {
        if _, ok := d[character]; ok {
            d[character] += 1
        } else {
            d[character] = 1
        }
    }

    for _, character := range t {
        if v, ok := d[character]; ok && v > 0 {
            d[character] -= 1
        } else {
            return false
        }
    }

    remaining := 0
    for _, value := range d {
        remaining += value
    }

    if remaining > 0 {
        return false
    }
    return true

    
}
