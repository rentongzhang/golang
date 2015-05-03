// Copyright (c) 2015, The Tony Authors.
// All rights reserved.
//
// Author: Rentong Zhang <rentongzh@gmail.com>

package base

import (
	"strings"
	"testing"
)

func TestGetRawHtml(t *testing.T) {
	crawler := NewCrawler()
	url := "http://www.taobao.com/"
	err, rawHtml := crawler.GetRawHtml(url, 1)
	if err != nil || len(rawHtml) < 1000 {
		t.Error("crawl taobao raw html falied")
	}
	t.Log(rawHtml)
}

func TestGetDomHtml(t *testing.T) {
	crawler := NewCrawler()
	url := "http://www.taobao.com/"
	err, domDoc := crawler.GetDomHtml(url, 1)
	if err != nil || domDoc == nil {
		t.Error("crawl taobao dom html falied")
	}

	if !strings.Contains(domDoc.Find("title").Text(), "淘宝网") {
		t.Error("taobao's domtree not find title ")
	}
	html, _ := domDoc.Html()
	t.Log(html)
}
