package source

func NewPlist(plist string) *Source {
	return newSource(plistRegex.FindStringSubmatch(plist))
}

func NewPlists(s []string) (plists Sources) {
	plists = make(Sources, 0, len(s))
	for _, html := range s {
		if plist := NewPlist(html); plist != nil {
			plists = append(plists, plist)
		}
	}
	return
}
