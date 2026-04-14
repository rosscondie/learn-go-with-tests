package iteration

import "strings"

func Repeat(character string, count int) string {
	var repeated strings.Builder

	// modern Go (Go 1.22+):
	for range count {
		repeated.WriteString(character)
	}

	// classic loop version:
	// for i := 0; i < count; i++ {
	// 	repeated.WriteString(character)
	// }

	return repeated.String()
}
