package luhn

import (
    s "strings"
    uc "unicode"
)

// Check if valid Luhn number or not
func Valid(inp string) bool {
    inp = s.ReplaceAll(inp, " ", "")
    if s.IndexFunc(inp, IsNotDigit) >= 0 || len(inp) < 2 {
        return false
    }

    count := 0
    for i,c := range inp {
        var temp int
        if (len(inp) - i)%2 == 0 {
            temp = 2*int(c - '0')
            if temp > 9 {
                temp -= 9
            }
        } else {
            temp = int(c - '0')
        }
        count += temp
    }

    return count%10 == 0
}

func IsNotDigit(c rune) bool {
    return !uc.IsDigit(c)
}
