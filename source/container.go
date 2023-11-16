package source

import "strconv"

type Source struct {
	Country, Name, City, Key string

	Channels int
}

func newSource(match []string) (p *Source) {
	if len(match) != 0 {
		p = &Source{Country: match[1], Name: match[2], City: match[3], Key: match[4]}
		p.Channels, _ = strconv.Atoi(match[5])
	}
	return
}
