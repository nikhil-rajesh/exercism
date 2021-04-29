package acronym

import s "strings"

// Returns the abbreviation of the input string
func Abbreviate(inp string) string {
	var out string

	// Format Input String
	inp = s.ReplaceAll(inp, "-", " ")
	inp = s.ReplaceAll(inp, "_", " ")
	inp = s.ToUpper(inp)

	for _, word := range s.Split(inp, " ") {
		if word != "" {
			out += string(word[0])
		}
	}

	return out
}
