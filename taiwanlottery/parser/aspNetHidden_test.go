package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseViewstate(t *testing.T) {
	content, err := ioutil.ReadFile("item_test.html")
	if err != nil {
		t.Error("readfile error %s", err)
	}
	a := string(content)
	_ = a
	viewstat := ParseViewstate(content)
	if viewstat == "" {
		t.Errorf("viewstat is null %s", viewstat)
	}
	fmt.Printf(viewstat)

}
func TestParseEventvalidation(t *testing.T) {

}
func TestParseDropDownList1(t *testing.T) {
	content, err := ioutil.ReadFile("item_test.html")
	if err != nil {
		t.Error("readfile error %s", err)
	}
	a := string(content)
	_ = a
	viewstat := ParseDropDownList1(content)
	if viewstat == "" {
		t.Errorf("viewstat is null %s", viewstat)
	}
	fmt.Printf(viewstat)
}
func TestParseHistory(t *testing.T) {
	content, err := ioutil.ReadFile("item_test.html")
	if err != nil {
		t.Error("readfile error %s", err)
	}
	a := string(content)
	_ = a
	viewstat := ParseDropYear(content)
	if viewstat == "" {
		t.Errorf("viewstat is null %s", viewstat)
	}
	fmt.Printf(viewstat)
}
