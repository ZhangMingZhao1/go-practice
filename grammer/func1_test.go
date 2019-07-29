package main

import (
	"fmt"
	"time"
)

//函数式编程，重用函数，装饰者模式
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}
