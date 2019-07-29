package main

import "testing"

func TestWithFunValue(t *testing.T) {
	//其实就是obj{int:func}类似的键值对，go常用这种映射去实现工厂模式
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))

}
