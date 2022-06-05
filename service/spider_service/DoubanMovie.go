package spider_service

import (
	"example.com/m/v2/models"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"log"
	"regexp"
	"strings"
	"time"
)

func DouBanMovie() {
	startUrl := "https://movie.douban.com/top250"
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36"),
	)
	// 设置抓取频率限制
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 500 * time.Millisecond,
		Parallelism: 12,
	})
	fmt.Println("312312")
	//解析列表页
	c.OnHTML(".grid_view", func(e *colly.HTMLElement) {
		e.DOM.Find("li").Each(func(i int, s *goquery.Selection) {
			href, found := s.Find("div.hd > a").Attr("href")
			if found {
				fmt.Printf("FOUND %v-> %s\n", i, href)
				ParseDetail(c, href)
			}
		})
	})
	// 查找下一页
	c.OnHTML("div.paginator > span.next", func(element *colly.HTMLElement) {
		href, found := element.DOM.Find("a").Attr("href")
		if found {
			element.Request.Visit(element.Request.AbsoluteURL(href))
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("Host", "movie.douban.com")
		r.Headers.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("response", len(response.Body))
	})
	//错误响应
	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error %s: %v\n", r.Request.URL, err)
	})
	c.Visit(startUrl)
}

type MovieDetailStruct struct {
	Title    string `json:"title",des:"电影名称"`
	Time     string `json:"time'",des:"上映时间"`
	Duration string `json:"duration",des:"播放时长"`
	Director string `json:"director",des:"导演"`
	FilmType string `json:"filmType",des:"类型"`
	Address  string `json:"address",des:"制片国家/地区"`
	Language string `json:"language",des:"语言"`
	Des      string `json:"des",des:"描述"`
}

func ParseDetail(collector *colly.Collector, href string) {
	c := collector.Clone()
	//解析详情页面
	c.OnHTML("body", func(e *colly.HTMLElement) {
		title := e.DOM.Find("#content>h1 span:first-child").Text()
		re := regexp.MustCompile("\\((.*?)\\)")
		timeTemp := e.DOM.Find("#content>h1 span:last-child").Text()
		time := re.FindStringSubmatch(timeTemp)
		duration := e.DOM.Find("#content #info span[property='v:runtime']").Text()
		director := e.DOM.Find("#content #info .attrs a[rel='v:directedBy']").Text()
		var _typeList []string
		e.DOM.Find("#content #info span[property='v:genre']").Each(func(i int, s *goquery.Selection) {
			_typeList = append(_typeList, s.Text())
		})
		filmType := strings.Join(_typeList, ",")
		des := e.DOM.Find("#link-report span[property='v:summary']").Text()
		re = regexp.MustCompile(`[\s\p{Zs}]{1,}`)
		des = re.ReplaceAllString(des, "")
		text := e.DOM.Find("#content #info ").Text()
		re = regexp.MustCompile("语言:(.*)")
		_language := re.FindStringSubmatch(text)
		re2 := regexp.MustCompile("制片国家/地区:(.*)")
		_address := re2.FindStringSubmatch(text)
		var language string
		if len(_language) > 1 {
			language = _language[1]
		} else {
			language = ""
		}
		var address string
		if len(_address) > 1 {
			address = _address[1]
		} else {
			address = ""
		}
		movie := models.DoubanMovie{
			Title:    title,
			Time:     time[1],
			Duration: duration,
			Director: director,
			FilmType: filmType,
			Address:  address,
			Language: language,
			Des:      des,
		}
		_, err := movie.AddMovie()
		if err != nil {
			log.Fatal("添加失败", title)
			return
		}
		fmt.Println("添加成功", title)
	})
	//错误响应
	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error %s: %v\n", r.Request.URL, err)
	})
	c.Visit(href)
}
