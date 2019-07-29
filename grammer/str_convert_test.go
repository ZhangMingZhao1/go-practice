package main

import (
	"strconv"
	"testing"
)

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	//注意转数字时后返回两个值，需要去接收判断
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
