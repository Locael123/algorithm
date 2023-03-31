package sort

func BubbleSort(list []int) int {
	loop := 0
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
			// single loop end
			loop++
		}
	}
	return loop
}
