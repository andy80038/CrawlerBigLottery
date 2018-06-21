package engine

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/andy80038/CrawlerBigLottery/fetcher"
	taiwanParser "github.com/andy80038/CrawlerBigLottery/taiwanlottery/parser"
)

type SimpleEngine struct {
	Url    string
	Origin string
	Parser func(doc *goquery.Document) []AwardsItme
}

func (s *SimpleEngine) Run() {
	content, err := fetcher.FetchToByte(s.Url)
	if err != nil {
		panic(err)
	}

	// value := taiwanParser.CreatePostFromValue(content, s.Origin, "radYM", "105", "8")
	// request := Request{
	// 	TargetUrl: s.Url,
	// 	Value:     value,
	// 	Parser:    s.Parser,
	// }
	// worker(request)

	for i := 103; i < 108; i++ {

		for j := 1; j < 13; j++ {

			value := taiwanParser.CreatePostFromValue(content, s.Origin, "radYM", strconv.Itoa(i), strconv.Itoa(j))
			request := Request{
				TargetUrl: s.Url,
				Value:     value,
				Parser:    s.Parser,
			}
			worker(request)
		}

	}
}
