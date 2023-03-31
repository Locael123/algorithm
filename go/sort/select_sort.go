package sort

func SelectSort(list []int) int {
	loop := 0
	for i := 0; i < len(list); i++ {
		min := i
		for j := i + 1; j < len(list); j++ {
			if list[min] > list[j] {
				min = j
			}
		}
		if min != i {
			list[i], list[min] = list[min], list[i]
		}
		// single loop end
		loop++
	}
	return loop
}
