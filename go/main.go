package main

import (
	"fmt"
	"go_algorithm/sort"
	"go_algorithm/utils"
)

func main() {
	// This is main, replace it
	l := utils.GenerateList(10)
	utils.ShowList(l, sort.InsertSort)

	// ---------------------
	fmt.Println("---------------------")

	// regular
	regularList := utils.GetRegularList()
	utils.ShowList(regularList, sort.InsertSort)
}
