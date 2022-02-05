/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/17 7:41 上午
# @File : util.go
# @Description :
# @Attention :
*/
package watcher

import (
	"fmt"
	"sync"
)

func merge2(a, b <-chan IData) <-chan IData {
	c := make(chan IData, 1)
	go func() {
		defer close(c)
		for nil != a || nil != b {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				c <- v
			}
		}
	}()

	return c
}

func mergeN(chans ...<-chan IData) <-chan IData {
	r := make(chan IData, 1)
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(chans))
		for _, c := range chans {
			go func(c <-chan IData) {
				for v := range c {
					r <- v
				}
				wg.Done()
			}(c)
		}
		wg.Wait()
		close(r)
	}()

	return r
}

var intFunc = func(v IData) {
	fmt.Println(v)
}

func asChan(vs ...IData) <-chan IData {
	c := make(chan IData)
	go func() {
		for _, v := range vs {
			c <- v
			// time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func calc1Counts(v int) int {
	r := 0
	for v > 0 {
		r++
		v = v & (v - 1)
	}
	return r
}
