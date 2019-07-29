package main

import (
	"fmt"
	"testing"
)

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	//key不存在时仍会返回零值，如果用if会判断false，不能通过返回nil来判断元素是否存在
	// m1[3] = 0
	if v, ok := m1[3]; ok {
		fmt.Println("key 3's value is", v, ok)
	} else {
		fmt.Println("key 3 is not existing.")
		t.Log("logggg")
	}
}

func TestTraveMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
