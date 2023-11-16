package ptv

import (
	"github.com/nikdissv-forever/goptv.ru/internal"
	"github.com/nikdissv-forever/goptv.ru/m3u"
	"github.com/nikdissv-forever/goptv.ru/pkg"
	"github.com/nikdissv-forever/goptv.ru/source"
	"strings"
)

func searchFunc[T any](post func([]string) T) func(string) (T, error) {
	return func(query string) (T, error) {
		lines, err := internal.Search(query)
		return post(lines), err
	}
}

// Functions which no automatically insert prefix at the start of query
var (
	QueryProvider = searchFunc(source.NewProviders)
	QueryPlist    = searchFunc(source.NewPlists)
	QueryCh       = searchFunc(m3u.NewChannels)
	QueryPl       = searchFunc(source.NewPl)
)

const (
	plist    = "plist"
	ch       = "ch:"
	pl       = "pl:"
	provider = "provider"
)

// Functions which automatically insert prefix at the start of query

func Provider(name string) (source.Source, error) { return QueryProvider(provider) }
func Plist() (source.Source, error)               { return QueryPlist(plist) }

func Ch(name string) ([]*m3u.Channel, error) { return QueryCh(pkg.Concat(ch, name)) }
func Pl(name string) (*source.Pl, error)     { return QueryPl(pkg.Concat(pl, name)) }

func Search(q string) (any, error) {
	if strings.EqualFold(q, plist) {
		return QueryPlist(q)
	} else if strings.EqualFold(q[:len(ch)], ch) {
		return QueryCh(q)
	} else if strings.EqualFold(q[:len(pl)], pl) {
		return QueryPl(q)
	}
	return QueryProvider(q)
}
