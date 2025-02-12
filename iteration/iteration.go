package iteration

import "strings"

func Repeat(r string, n int) string {
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		sb.WriteString(r)
	}
	return sb.String()
}
