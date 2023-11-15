package m3u

import (
	"regexp"
	"strconv"
	"strings"
)

type Extinf struct {
	Raw          string
	Duration, ID int
	Group        string
	Name         string
	URL          string
}

type ExtInfos []*Extinf

func (e *Extinf) String() string {
	var b strings.Builder
	b.Grow(len(e.Raw) + 1 + len(e.URL))
	b.WriteString(e.Raw)
	b.WriteByte('\n')
	b.WriteString(e.URL)
	return b.String()
}

var extinfRegex = regexp.MustCompile(`#EXTINF:(-?\d+)\s*tvch-id="(\d+)"\s*group-title="(.+)",\s*(.+)\s*`)

func NewExtinf(extinf, url string) (r *Extinf) {
	if match := extinfRegex.FindStringSubmatch(extinf); len(match) != 0 {
		r = &Extinf{Raw: match[0]}
		r.Duration, _ = strconv.Atoi(match[1])
		r.ID, _ = strconv.Atoi(match[2])
		r.Group = match[3]
		r.Name = match[4]
	}
	if url != "" {
		if r == nil {
			r = &Extinf{}
		}
		r.URL = url
	}
	return
}
func NewExtInfos(extinf []string) (sequence ExtInfos) {
	const step = 3
	sequence = make(ExtInfos, 0, len(extinf)/step)
	for i := 0; i < len(extinf)-step; i += step {
		if info := NewExtinf(extinf[i], extinf[i+1]); info != nil {
			sequence = append(sequence, info)
		}
	}
	return
}
