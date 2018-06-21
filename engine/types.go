package engine

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type AwardsItme struct {
	//期號
	Period string
	//開獎時間
	OpenTime string
	//銷售金額
	SalesMoney int
	//獎金總額
	AwardsTotalMoney string
	//獎號
	AwardsNumber []string
	//特別號
	SpecialNumber string
}

type ParseItem struct {
	//解析網址
	TargetUrl      string
	Viewstate      string
	Eventalidation string
	//查詢的年
	Year string
	//查詢的月
	Month string
}
type Request struct {
	TargetUrl string
	Value     url.Values
	Parser    func(doc *goquery.Document) []AwardsItme
}
