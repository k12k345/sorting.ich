package selectionsort

// SmallestElementIndex gibt den Index des kleinsten Elements im Array arr ab Index start zurück.
func SmallestElementIndex(arr []int, start int) int {
	for i := 0; i < len(arr)-1; i++ {

		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			temp := arr[i]
			arr[i] = arr[minIndex]
			arr[minIndex] = temp
		}
	}

	return 0
}

// SelectionSort sortiert die übergebene Liste mittels des Selection-Sort-Algorithmus.
func SelectionSort(arr []int) {
	SmallestElementIndex(arr, 0)
}
