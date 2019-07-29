package main

import (
	"fmt"
	"testing"
)

func deferFunc() {
	fmt.Println("Clean resources")
}
func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println("Clear resources")
	}()
	fmt.Println("Start")
	//defer仍会执行到，return之前会执行的操作，释放资源常用，就算下面语句报错也会执行
	panic("some error")
	fmt.Println("end") //报不可到达的错误
}
