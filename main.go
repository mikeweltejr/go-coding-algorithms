package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Insert(11)
	obj.Insert(10)
	obj.Insert(20)
	obj.Insert(30)
	rand := obj.GetRandom()
	rand2 := obj.GetRandom()

	fmt.Println(rand)
	fmt.Println(rand2)
}
