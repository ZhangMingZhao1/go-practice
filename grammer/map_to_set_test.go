package main

import "testing"

func TestMapForSet(t *testing.T) {
	//go没有内置set，map模仿set就是key为int value为bool
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	mySet[n] = false
	t.Log(len(mySet))
	delete(mySet, 1)
	t.Log(len(mySet))

}
