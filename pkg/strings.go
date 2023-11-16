package pkg

import (
	"fmt"
	"strings"
)

func Concat(s1, s2 string) string {
	var b strings.Builder
	b.Grow(len(s1) + len(s2))
	b.WriteString(s1)
	b.WriteString(s2)
	return b.String()
}
func NewlinesJoin[T fmt.Stringer](s []T) string {
	r := make([]string, len(s))
	for i, source := range s {
		r[i] = source.String()
	}
	return strings.Join(r, "\n")
}
