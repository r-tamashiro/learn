package main
// https://qiita.com/ryurock/items/b0466ad144f5e0555a95

import (
    //"net/http"
    //"io/ioutil"
    "fmt"
    "github.com/PuerkitoBio/goquery"
)


func main() {
    doc, err := goquery.NewDocument("https://ja.wikipedia.org/wiki/Linux")
    if err != nil {
        panic(err)
    }
    doc.Find("title").Each(func(_ int, s *goquery.Selection) {
        fmt.Println(s.Text())
        //url, _ := s.Attr("href")
        //fmt.Println(url)
    })
}
