// Copyright (c) 2015, The Tony Authors.
// All rights reserved.
//
// Author: Rentong Zhang <rentongzh@gmail.com>

package base

import (
	"testing"
)

type UrlTest struct {
	base     string
	ref      string
	expected string
}

func (s UrlTest) String() string {
	ret := ""
	ret += ("base:" + s.base + "\t")
	ret += ("ref:" + s.ref + "\t")
	ret += ("expected:" + s.expected)
	return ret
}

var urlTestSet = []UrlTest{
	UrlTest{
		base:     "http://index.bitauto.com/guanzhu/zhongdaxingche/changsha/",
		ref:      "/guanzhu/s2354/",
		expected: "http://index.bitauto.com/guanzhu/s2354/",
	},
	UrlTest{
		base:     "http://index.bitauto.com/guanzhu/zhongdaxingche/changsha",
		ref:      "./index.html",
		expected: "http://index.bitauto.com/guanzhu/zhongdaxingche/index.html",
	},
	UrlTest{
		base:     "http://index.bitauto.com/guanzhu/zhongdaxingche/changsha/",
		ref:      "./index.html",
		expected: "http://index.bitauto.com/guanzhu/zhongdaxingche/changsha/index.html",
	},
	UrlTest{
		base:     "http://index.bitauto.com/guanzhu/zhongdaxingche/changsha/",
		ref:      "../index.html",
		expected: "http://index.bitauto.com/guanzhu/zhongdaxingche/index.html",
	},
}

func TestUrlParse(t *testing.T) {
	for _, item := range urlTestSet {
		ret := Parse(item.base, item.ref)
		if ret != item.expected {
			t.Error("not get expected url:", item.String())
		}
	}
}
