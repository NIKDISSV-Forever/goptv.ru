package m3u

import (
	"regexp"
	"strings"
)

type Header struct{ Country, Provider, City string }

func (h *Header) String() string { return EXTM3UHeader(h.Provider) }

func EXTM3UHeader(author string) string {
	const start = `#EXTM3U list-autor="`
	var b strings.Builder
	b.Grow(len(start) + len(author) + 1)
	b.WriteString(start)
	b.WriteString(author)
	b.WriteByte('"')
	return b.String()
}

var headerRegex = regexp.MustCompile(`<img\s+.+src=".+/(\w{2})\.gif">\s*([^,]+),\s*г\.\s*([^<]+)`)

func NewHeader(header string) *Header {
	if match := headerRegex.FindStringSubmatch(header); len(match) != 0 {
		return &Header{Country: match[1], Provider: match[2], City: match[3]}
	}
	return nil
}
