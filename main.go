package main

import (
	"time"

	"github.com/andy80038/CrawlerBigLottery/biglottery/parser"
	"github.com/andy80038/CrawlerBigLottery/config"
	"github.com/andy80038/CrawlerBigLottery/engine"
	lottoParser "github.com/andy80038/CrawlerBigLottery/superLotto638/parser"
)

func main() {

	simp := engine.SimpleEngine{
		Url:    config.BiglotteryUrl,
		Parser: parser.Parser,
		Origin: config.Biglottery,
	}
	//simp.Run()
	_ = simp

	lotto638 := engine.SimpleEngine{
		Url:    config.SuperLotto638Url,
		Parser: lottoParser.Parser,
		Origin: config.SuperLotto638,
	}
	lotto638.Run()

	time.Sleep(50 * time.Second)
}
