// Copyright (c) 2014, The Tony Authors.
// All rights reserved.
//
// Author: Rentong Zhang <rentongzh@gmail.com>

package base

import (
	"sync"
	"time"
)

type Counter struct {
	data      map[string]int64
	startTime string
	lock      *sync.RWMutex
}

func NewCounter() *Counter {
	return &Counter{
		data:      make(map[string]int64),
		startTime: time.Now().Format("2000-01-01 01:01:01"),
		lock:      &sync.RWMutex{},
	}
}

func (self *Counter) Incr(key string, val int64) {
	self.lock.Lock()
	defer self.lock.Unlock()
	_, ok := self.data[key]
	if !ok {
		self.data[key] = val
	} else {
		self.data[key] += val
	}
}

func (self *Counter) Get(key string) int64 {
	self.lock.RLock()
	defer self.lock.RUnlock()
	val, ok := self.data[key]
	if !ok {
		return 0
	}
	return val
}

func (self *Counter) Set(key string, val int64) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.data[key] = val
}

func (self *Counter) Stat() map[string]int64 {
	self.lock.RLock()
	defer self.lock.RUnlock()
	return self.data
}
