package ptv

import (
	"github.com/nikdissv-forever/goptv.ru/internal"
	"github.com/nikdissv-forever/goptv.ru/m3u"
	"github.com/nikdissv-forever/goptv.ru/pkg"
	"github.com/nikdissv-forever/goptv.ru/plist"
)

func Ch(query string) (m3u.Channels, error) {
	r, err := internal.Search(pkg.Concat("ch:", query))
	if err != nil {
		return nil, err
	}
	return m3u.NewChannels(r.Lines()), nil
}

func Plist() (plist.Plists, error) {
	r, err := internal.Search("plist")
	if err != nil {
		return nil, err
	}
	return plist.NewPlists(r.Lines()), nil
}
func Pl(query string) (*plist.Pl, error) {
	r, err := internal.Search(pkg.Concat("pl:", query))
	if err != nil {
		return nil, err
	}
	return plist.NewPl(r.Lines()), nil
}
