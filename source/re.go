package source

import "regexp"

var plistRegex = regexp.MustCompile(`<img\s+.+src=".+/(\w{2})\.gif">\s*([^,]+),\s*г\.\s*([^,]+),\s*плейлист\s+"([^"]+)",\s*ТВ\s*каналов:\s*(\d+)`)
var providerRegex = regexp.MustCompile(`<img\s+.+src=".+/(\w{2})\.gif">\s*([^,]+),\s*г\.\s*([^,]+),\s*робот\s+"([^"]+)",\s*ТВ\s*каналов:\s*(\d+)`)
