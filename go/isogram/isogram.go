package isogram

import "strings"

func IsIsogram(inp string) bool {
    count := make(map[int32]int32)
    inp = strings.ToLower(inp)
    inp = strings.ReplaceAll(inp, "-", "")
    inp = strings.ReplaceAll(inp, " ", "")

    for _,c := range inp {
        if count[c]++; count[c] > 1 {
            return false
        }
    }

    return true
}
