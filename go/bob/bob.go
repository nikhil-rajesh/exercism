// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
    "strings"
    "unicode"    
)

func IsEmpty(s string) bool {
    return s == ""
}

func IsQuestion(s string) bool {
    return strings.HasSuffix(s, "?")
}

func IsUpper(s string) bool {
    return ContainsLetter(s) && strings.ToUpper(s) == s
}

func ContainsLetter(s string) bool {
    for _, c := range s  {
        if unicode.IsLetter(c) {
            return true
        }
    }
    return false
}

func Hey(remark string) string {
    r := strings.TrimSpace(remark)

    switch {
        case IsEmpty(r):
            return "Fine. Be that way!"
        case IsUpper(r) && IsQuestion(r):
            return "Calm down, I know what I'm doing!"
        case IsUpper(r):
            return "Whoa, chill out!"
        case IsQuestion(r):
            return "Sure."
        default:
            return "Whatever."
    }
}
