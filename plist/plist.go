package plist

import (
	"regexp"
	"strconv"
)

type Plist struct {
	Country, Name, City, Pl string

	Channels int
}
type Plists []*Plist

var plistRegex = regexp.MustCompile(`<img\s+.+src=".+/(\w{2})\.gif">\s*([^,]+),\s*г\.\s*([^,]+),\s*плейлист\s+"([^"]+)",\s*ТВ\s*каналов:\s*(\d+)`)

func NewPlist(plist string) (p *Plist) {
	if match := plistRegex.FindStringSubmatch(plist); len(match) != 0 {
		p = &Plist{}
		p.Country = match[1]
		p.Name = match[2]
		p.City = match[3]
		p.Pl = match[4]
		p.Channels, _ = strconv.Atoi(match[5])
	}
	return
}

func NewPlists(s []string) (plists Plists) {
	plists = make(Plists, 0, len(s))
	for _, html := range s {
		if plist := NewPlist(html); plist != nil {
			plists = append(plists, plist)
		}
	}
	return
}
