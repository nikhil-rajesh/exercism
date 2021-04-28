package bob

import (
    "strings"
    "unicode"    
)

// Check if strings is empty or not
func IsEmpty(s string) bool {
    return s == ""
}

// Check if input is a question
func IsQuestion(s string) bool {
    return strings.HasSuffix(s, "?")
}

// Check if input is in upper case
func IsUpper(s string) bool {
    return ContainsLetter(s) && strings.ToUpper(s) == s
}

// Check if string contains alphabets
func ContainsLetter(s string) bool {
    for _, c := range s  {
        if unicode.IsLetter(c) {
            return true
        }
    }
    return false
}

// Returns responses to input strings
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
