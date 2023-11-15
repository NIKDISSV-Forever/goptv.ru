package main

import (
	"fmt"
	"github.com/nikdissv-forever/goptv.ru"
)

func main() {
	plist, err := ptv.Plist()
	if err != nil {
		panic(err)
	}
	for _, i := range plist {
		if pl, err := ptv.Pl(i.Pl); err == nil {
			for _, info := range pl.ExtInfos {
				fmt.Println(info)
			}
		}
	}
}
