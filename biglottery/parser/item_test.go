package parser

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func TestParser(t *testing.T) {
	content, err := ioutil.ReadFile("item_test.html")
	if err != nil {
		t.Errorf("readFile error : %s", err)
	}
	contentStr := string(content)
	node, err := html.Parse(strings.NewReader(contentStr))

	doc := goquery.NewDocumentFromNode(node)
	if err != nil {
		t.Errorf("%s", err)
	}
	item := Parser(doc)
	if len(item) < 0 {
		t.Errorf("GetItme Error  len is 0 ")
	}
	fmt.Printf("%+v", item)
}
