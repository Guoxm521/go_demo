package spider_service

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func Baidu() {
	fmt.Println("执行")
	//colly.AllowedDomains("www.baidu.com"),
	c := colly.NewCollector()
	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//	link := e.Attr("href")
	//	fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	//	c.Visit(e.Request.AbsoluteURL(link))
	//})
	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Response %s: %d bytes\n", r.Request.URL, len(r.Body))
		fmt.Printf(string(r.Body))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error %s: %v\n", r.Request.URL, err)
	})

	c.Visit("http://www.baidu.com/")
}
