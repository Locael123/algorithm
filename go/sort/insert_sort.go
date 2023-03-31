package sort

func InsertSort(list []int) int {
	loop := 0
	for i := 1; i < len(list); i++ {
		value := list[i]
		pos := i
		for pos > 0 && list[pos-1] > value {
			list[pos] = list[pos-1]
			pos--
			// single loop end
			loop++
		}
		list[pos] = value
	}
	return loop
}
