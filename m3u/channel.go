package m3u

type Channel struct {
	*Header
	*Extinf
}

type Channels []*Channel

func NewChannel(header *Header, m3uInfo *Extinf) *Channel {
	if header == nil || m3uInfo == nil {
		return nil
	}
	return &Channel{Header: header, Extinf: m3uInfo}
}

func NewChannels(s []string) (channels Channels) {
	const step = 5
	channels = make(Channels, 0, len(s)/step)
	for i := 0; i < len(s)-step; i += step {
		if channel := NewChannel(NewHeader(s[i]), NewExtinf(s[i+2], s[i+3])); channel != nil {
			channels = append(channels, channel)
		}
	}
	return
}
