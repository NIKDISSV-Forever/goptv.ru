package internal

import (
	"github.com/nikdissv-forever/goptv.ru/pkg"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const SearchURL = "https://proxytv.ru/iptv/php/srch.php"

var DefaultSearchHeaders = http.Header{
	"Referer": {"https://proxytv.ru/"}, "User-Agent": {"Mozilla/5.0"},
	"Content-Type": {"application/x-www-form-urlencoded"}}

func Search(query string) (r Result, err error) {
	request, err := http.NewRequest("POST", SearchURL, strings.NewReader(pkg.Concat("udpxyaddr=", url.QueryEscape(query))))
	if err != nil {
		return
	}
	request.Header = DefaultSearchHeaders
	do, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer do.Body.Close()
	if do.ContentLength != -1 {
		r = make([]byte, do.ContentLength)
		_, err = do.Body.Read(r)
		return
	}
	return io.ReadAll(do.Body)
}

type Result []byte

func (r Result) Lines() []string {
	div := r.ResultsDiv()
	if hr := strings.LastIndex(div, "<br>"); hr != -1 {
		div = div[:hr]
	}
	return strings.Split(div, "<br>")
}
func (r Result) ResultsDiv() string { return strings.SplitN(string(r), "<hr>", 3)[1] }
