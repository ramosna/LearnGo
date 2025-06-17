package iteration

import "strings"

func Repeat(letter string, times int) string {
	var result strings.Builder
	for i := 0; i < times; i++ {
		result.WriteString(letter)
	}
	return result.String()
}
