package source

func NewProvider(provider string) *Source {
	return newSource(providerRegex.FindStringSubmatch(provider))
}

func NewProviders(s []string) (providers Sources) {
	providers = make(Sources, 0, len(s))
	for _, html := range s {
		if plist := NewProvider(html); plist != nil {
			providers = append(providers, plist)
		}
	}
	return
}
