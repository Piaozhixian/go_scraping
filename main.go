package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://piaozhixian.github.io/")
	if err != nil {
		panic("取得失敗")
	}
	h2 := doc.Find("h2")
	h2.Each(func(i int, s *goquery.Selection) {
		fmt.Printf(s.Text())
	})
}

