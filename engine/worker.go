package engine

import (
	"fmt"

	"github.com/andy80038/CrawlerBigLottery/fetcher"
)

func worker(r Request) []AwardsItme {
	doc, err := fetcher.FetchByForm(r.TargetUrl, r.Value)
	if err != nil {
		fmt.Printf("FetchByForm err%s", err)
	}
	return r.Parser(doc)
}
