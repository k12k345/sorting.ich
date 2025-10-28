package insertionsort

// MoveLeft bewegt das Element an Stelle i so lange nach links,
// bis es an der richtigen Stelle im bereits sortierten Teil der Liste ist.
func MoveLeft(arr []int, i int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Verschiebe Elemente, die größer als key sind, um eine Position nach rechts
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]

		}
		arr[j+1] = key
	}
}

// InsertionSort sortiert die übergebene Liste mittels des Insertion-Sort-Algorithmus.
func InsertionSort(arr []int) {
	MoveLeft(arr, len(arr))
}
