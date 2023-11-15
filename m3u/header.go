package m3u

import "regexp"

type Header struct {
	Country, Provider, City string
}

var headerRegex = regexp.MustCompile(`<img\s+.+src=".+/(\w{2})\.gif">\s*([^,]+),\s*г\.\s*([^<]+)`)

func NewHeader(header string) *Header {
	if match := headerRegex.FindStringSubmatch(header); len(match) != 0 {
		return &Header{Country: match[1], Provider: match[2], City: match[3]}
	}
	return nil
}
