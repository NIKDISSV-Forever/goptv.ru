package pkg

import "strings"

func Concat(s1, s2 string) string {
	var b strings.Builder
	b.Grow(len(s1) + len(s2))
	b.WriteString(s1)
	b.WriteString(s2)
	return b.String()
}
