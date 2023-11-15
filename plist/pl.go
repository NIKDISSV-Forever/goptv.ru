package plist

import "github.com/nikdissv-forever/goptv.ru/m3u"

// Pl One provider m3u
type Pl struct {
	Header *m3u.Header
	m3u.ExtInfos
}

func NewPl(s []string) *Pl {
	if len(s) != 0 {
		if h := m3u.NewHeader(s[0]); h != nil {
			return &Pl{Header: h, ExtInfos: m3u.NewExtInfos(s[1:])}
		}
	}
	return nil
}
