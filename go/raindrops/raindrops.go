package raindrops

import "strconv"

var sounds = map[int]string {
    3: "Pling",
    5: "Plang",
    7: "Plong",
}

func Convert(num int) string {
    out := ""

    for fact, sound := range sounds {
        if num%fact == 0 {
            out += sound
        }
    }

    if out == "" {
        return strconv.Itoa(num)
    }

    return out
}
