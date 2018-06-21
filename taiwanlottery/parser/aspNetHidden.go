package parser

import (
	"fmt"
	"net/url"
	"regexp"
)

var viewstateRe = `<input type="hidden" name="__VIEWSTATE" id="__VIEWSTATE" value="([^"]+)"`
var eventvalidationRe = `<input type="hidden" name="__EVENTVALIDATION" id="__EVENTVALIDATION" value="([^"]*)"`
var dropDownList1Re = `<select name="([^$]+\$DropDownList1)"`
var historyRe = `type="radio" name="([^$]+\$chk)"`
var dropYear = ` <select name="([^$]+\$dropYear)"`
var dropMonth = ` <select name="([^$]+\$dropMonth)"`
var btnSubmitRe = `<input type="submit" name="([^$]+\$btnSubmit)"`

func CreatePostFromValue(contene []byte, listIndex string, historychk string, year string, month string) url.Values {

	value := url.Values{
		"__VIEWSTATE":       {ParsePostValue(contene, viewstateRe)},
		"__EVENTVALIDATION": {ParsePostValue(contene, eventvalidationRe)},
		//"__VIEWSTATEENCRYPTED":                   {""}, //38need
		//"__EVENTTARGET":                          {""}, //38need
		ParsePostValue(contene, dropDownList1Re): {listIndex},
		ParsePostValue(contene, historyRe):       {historychk},
		ParsePostValue(contene, dropYear):        {year},
		ParsePostValue(contene, dropMonth):       {month},
		ParsePostValue(contene, btnSubmitRe):     {"查詢"},
	}
	return value
}

func ParsePostValue(content []byte, strRe string) string {
	re := regexp.MustCompile(strRe)
	matches := re.FindAllSubmatch(content, -1)
	for _, v := range matches {
		//fmt.Printf("%s\n\n", v)
		return string(v[1])
	}
	return ""
}

func ParseDropYear(content []byte) string {
	re := regexp.MustCompile(dropYear)
	matches := re.FindAllSubmatch(content, -1)
	for _, v := range matches {
		fmt.Printf("%s\n\n", v)
		//fmt.Printf("V1:%s\n\n", v[1])
		return string(v[1])
	}
	return ""
}

func ParseHistory(content []byte) string {
	re := regexp.MustCompile(historyRe)
	matches := re.FindAllSubmatch(content, -1)
	for _, v := range matches {
		fmt.Printf("%s\n\n", v)
		//fmt.Printf("V1:%s\n\n", v[1])
		return string(v[1])
	}
	return ""
}

func ParseDropDownList1(content []byte) string {
	re := regexp.MustCompile(dropDownList1Re)
	matches := re.FindAllSubmatch(content, -1)
	for _, v := range matches {
		fmt.Printf("%s\n\n", v)
		//fmt.Printf("V1:%s\n\n", v[1])
		return string(v[1])
	}
	return ""
}

func ParseViewstate(content []byte) string {
	re := regexp.MustCompile(viewstateRe)
	matches := re.FindAllSubmatch(content, -1)
	for _, v := range matches {
		return string(v[1])
	}
	return ""
}
func ParseEventvalidation(content []byte) string {
	re := regexp.MustCompile(eventvalidationRe)
	matches := re.FindAllSubmatch(content, -1)
	for _, v := range matches {

		return string(v[1])
	}
	return ""
}
