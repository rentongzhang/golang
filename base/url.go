// Copyright (c) 2015, The Tony Authors.
// All rights reserved.
//
// Author: Rentong Zhang <rentongzh@gmail.com>

package base

import (
	"net/url"
)

func Parse(base, link string) string {
	iUrl, err := url.Parse(base)
	if err != nil {
		return link
	}

	ref, err := url.Parse(link)
	if err != nil {
		return link
	}

	return iUrl.ResolveReference(ref).String()
}
