package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/headzoo/surf.v1"
)

func main() {
	browser := surf.NewBrowser() // Creates browser
	url := "https://www.instagram.com/explore/locations/213480180/"
	err := browser.Open(url)
	if err != nil {
		panic(err)
	}
	jsVar := browser.Dom().Find("script").Eq(1).Text()
	jsJSON := jsVar[21 : len(jsVar)-1]

	m := Instagram{}
	json.Unmarshal([]byte(jsJSON), &m)
	nodeLen := m.EntryData.LocationsPage[0].Location.TopPosts.Nodes
	posts := []Post{}
	top := TopPosts{posts}
	for _, v := range nodeLen {
		top.AddItem(v)
	}
	out, _ := json.Marshal(top)
	fmt.Println(string(out)) // JSON result
}
