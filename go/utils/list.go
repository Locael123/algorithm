package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateList(length int) []int {
	list := make([]int, length)
	rand.Seed(time.Now().Unix())

	for i := range list {
		list[i] = rand.Intn(length)
	}

	return list
}

func GetRegularList() []int {
	return []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
}

func ShowList(list []int, sortFunc func([]int) int) {
	fmt.Println(list)
	loop := sortFunc(list)
	fmt.Println(list)
	fmt.Println("Loop : ", loop)
}
