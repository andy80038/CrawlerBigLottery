package parser

import (
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/andy80038/CrawlerBigLottery/engine"
)

var classorgRe = ".table_org.td_hm"
var classgreRe = ".table_gre.td_hm"
var periodRe = "td .td_w #Lotto649Control_history_dlQuery_L649_DrawTerm_" //DrawTerm_[0-9]
var openTimeRe = "td .td_w #Lotto649Control_history_dlQuery_L649_DDate_"
var openAwardsRe = "td .td_w #Lotto649Control_history_dlQuery_No"                         //No[0-9]_[0-9]
var specialNumberRe = "td .td_w.font_red14b_center #Lotto649Control_history_dlQuery_SNo_" //SNo_[0-9]
var awardsTotalMoneyRe = "td .td_w #Lotto649Control_history_dlQuery_Total_"               //[0-9]

func Parser(doc *goquery.Document) []engine.AwardsItme {
	itemArr := []engine.AwardsItme{}
	itemA := parserBigLottery(classorgRe, func(i int) int {
		return i * 2
	}, doc)
	itemB := parserBigLottery(classgreRe, func(i int) int {
		return i*2 + 1
	}, doc)

	itemArr = append(itemArr, itemA...)
	itemArr = append(itemArr, itemB...)

	return itemArr
}

func parserBigLottery(classRe string, operation func(int) int, doc *goquery.Document) []engine.AwardsItme {
	itemArr := []engine.AwardsItme{}
	doc.Find(classRe).Each(func(i int, s *goquery.Selection) {
		item := engine.AwardsItme{}
		offset := strconv.Itoa(operation(i))
		periodMatch := periodRe + offset
		period := s.Find(periodMatch).Text()
		item.Period = period
		fmt.Printf("period:%s\n", period)

		openTimeMatch := openTimeRe + offset
		openTime := s.Find(openTimeMatch).Text()
		item.OpenTime = openTime
		fmt.Printf("openTime:%s\n", openTime)

		specialNumberMatch := specialNumberRe + offset
		specialNumber := s.Find(specialNumberMatch).Text()
		item.SpecialNumber = specialNumber
		fmt.Printf("specialNumber:%s\n", specialNumber)

		awardsTotalMoneyMatch := awardsTotalMoneyRe + offset
		awardsTotalMoney := s.Find(awardsTotalMoneyMatch).Text()
		item.AwardsTotalMoney = awardsTotalMoney
		fmt.Printf("awardsTotalMoney:%s\n", awardsTotalMoney)

		for index := 1; index <= 6; index++ {
			openAwardsMatch := openAwardsRe + strconv.Itoa(index) + "_" + offset
			openAwards := s.Find(openAwardsMatch).Text()
			item.AwardsNumber = append(item.AwardsNumber, openAwards)
		}
		fmt.Printf("AwardsNumber:%s\n", item.AwardsNumber)

		itemArr = append(itemArr, item)
	})
	return itemArr
}
