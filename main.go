package main

import (
  "fmt"
  "encoding/json"
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
  jsJson := jsVar[21:len(jsVar)-1]

  m := Instagram{}
  json.Unmarshal([]byte(jsJson), &m)
  node_len := m.EntryData.LocationsPage[0].Location.TopPosts.Nodes
  posts := []Post{}
  top := TopPosts{posts}
  for _,v := range node_len {
    top.AddItem(v)
  }
  out, _ := json.Marshal(top)
  fmt.Println(string(out)) // JSON result
}
