package main

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	Set map[int]int
}

var ints []int

func Constructor() RandomizedSet {
	var r RandomizedSet
	r.Set = make(map[int]int)
	ints = []int{}

	return r
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.Set[val]

	if ok == true {
		return false
	}

	ints = append(ints, val)
	this.Set[val] = len(ints) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.Set[val]

	if ok == false {
		return false
	}

	this.Set[ints[len(ints)-1]] = index
	ints[index] = ints[len(ints)-1]
	ints = ints[:len(ints)-1]
	delete(this.Set, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	if len(ints) == 0 {
		return ints[0]
	}

	s2 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s2)
	return ints[r.Intn(len(ints))]
}
