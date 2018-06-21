package parser

import (
	"io/ioutil"
	"strings"
	"testing"

	"fmt"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

//var classorgRe = ".table_org.td_hm"
//var classgreRe = ".table_gre.td_hm"
var numberRe = "td .td_w.font_black14b_center #SuperLotto638Control_history1_dlQuery_NoX_X" //開獎號碼順序[1-6]   排序順序[0-n]
var seRe = "td .td_w.font_red14b_center #SuperLotto638Control_history1_dlQuery_No7_X"       // 排序順序[0-n]
var totRe = "td #SuperLotto638Control_history1_dlQuery_Total_X" //排序順序[0-n]
func TestParser(t *testing.T) {
	content, err := ioutil.ReadFile("item_testHtml.html")
	if err != nil {
		t.Errorf("readFile error : %s", err)
	}
	contentStr := string(content)
	node, err := html.Parse(strings.NewReader(contentStr))

	doc := goquery.NewDocumentFromNode(node)
	if err != nil {
		t.Errorf("%s", err)
	}
	//doc.Find(".table_org.td_hm td .td_w.font_black14b_center #SuperLotto638Control_history1_dlQuery_No2_0").
	//	Each(func(i int, s *goquery.Selection){
	//	fmt.Printf(" J:%d  ss:%s \n",i,s.Text())
	//})

	//
	//doc.Find(".table_org.td_hm").Each(func(i int, s *goquery.Selection) {
	//
	//	s.Find("td .td_w.font_black14b_center #SuperLotto638Control_history1_dlQuery_No2_0").Each(func(j int, ss *goquery.Selection){
	//		fmt.Printf(" J:%d  ss:%s \n",j,ss.Text())
	//	})
	//	fmt.Printf("%d\n",i)
	//})

	doc.Find(".table_org.td_hm").Each(func(i int, s *goquery.Selection) {

		n := s.Find("td .td_w.font_black14b_center #SuperLotto638Control_history1_dlQuery_No2_0").Text()
		fmt.Printf("N:%s\n", n)

		sp := s.Find("td .td_w.font_red14b_center #SuperLotto638Control_history1_dlQuery_No7_0").Text()
		fmt.Printf("sp:%s\n", sp)

		totla := s.Find("td #SuperLotto638Control_history1_dlQuery_Total_0").Text()
		fmt.Printf("totla:%s\n", totla)

	})

}
