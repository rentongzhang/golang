// Copyright (c) 2015, The Tony Authors.
// All rights reserved.
//
// Author: Rentong Zhang <rentongzh@gmail.com>

package base

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Crawler struct {
	utf8Converter *Utf8Converter
}

func NewCrawler() *Crawler {
	return &Crawler{
		utf8Converter: NewUtf8Converter(),
	}
}

func (c *Crawler) httpGet(url string) (error, []byte) {
	r, err := http.Get(url)
	if err != nil {
		return err, []byte("")
	}
	if r.Body != nil {
		defer r.Body.Close()
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err, []byte("")
	}
	return nil, c.utf8Converter.ToUTF8(b)
}

func (c *Crawler) GetRawHtml(url string, repeatTimes int) (error, string) {
	if len(url) == 0 {
		return errors.New("invalid url"), ""
	}
	var retErr error
	for i := 0; i < repeatTimes; i++ {
		err, html := c.httpGet(url)
		if err != nil {
			retErr = err
			time.Sleep(1 * time.Second)
			continue
		}
		return nil, string(html)
	}
	return retErr, ""
}

func (c *Crawler) GetDomHtml(url string, repeatTimes int) (error,
	*goquery.Document) {
	if len(url) == 0 {
		return errors.New("invalid url"), nil
	}

	var retErr error
	for i := 0; i < repeatTimes; i++ {
		err, html := c.httpGet(url)
		if err != nil {
			retErr = err
			time.Sleep(1 * time.Second)
			continue
		}
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
		if err != nil {
			retErr = err
			continue
		}
		return nil, doc
	}
	return retErr, nil
}

func (c *Crawler) GetRawHtmlByPost(api string, params url.Values) (error,
	string) {
	if len(api) == 0 {
		return errors.New("invalid api"), ""
	}
	req, err := http.NewRequest("POST", api, strings.NewReader(params.Encode()))
	if err != nil {
		return err, ""
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	hc := http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		return err, ""
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, ""
	}
	return nil, string(c.utf8Converter.ToUTF8(data))
}

func (c *Crawler) GetDomHtmlByPost(api string, params url.Values) (error, *goquery.Document) {
	err, rawHtml := c.GetRawHtmlByPost(api, params)
	if err != nil {
		return err, nil
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawHtml))
	if err != nil {
		return err, nil
	}
	return nil, doc
}
