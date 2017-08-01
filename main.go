package main

import (
  "fmt"
  "gopkg.in/headzoo/surf.v1"
  "github.com/PuerkitoBio/goquery"
)

func main() {
  browser := surf.NewBrowser() // Creates browser
  url := "https://www.instagram.com/explore/locations/213480180/washington-district-of-columbia/"
  err := browser.Open(url)
  if err != nil {
    panic(err)
  }
  browser.Dom().Find("script").Each(func(idx int, s *goquery.Selection) {
    fmt.Println(idx, s.Text())
  })
}
