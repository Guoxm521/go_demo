module example.com/m/v2

go 1.17

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-ini/ini v1.66.4
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.10.1
	github.com/jinzhu/gorm v1.9.16
	github.com/unknwon/com v1.0.1
)

require (
	github.com/PuerkitoBio/goquery v1.8.0 // indirect
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/antchfx/htmlquery v1.2.5 // indirect
	github.com/antchfx/xmlquery v1.3.11 // indirect
	github.com/antchfx/xpath v1.2.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocolly/colly/v2 v2.1.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/temoto/robotstxt v1.1.2 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.0.0-20220331220935-ae2d96664a29 // indirect
	golang.org/x/net v0.0.0-20220531201128-c960675eff93 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	pkg/setting => ../pkg/setting
	pkg/util => ../pkg/util
)
