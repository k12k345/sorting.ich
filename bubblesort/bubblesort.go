package bubblesort

// BubbleUp geht einmal durch das Array arr und tauscht benachbarte Elemente,
// wenn das linke größer ist als das rechte.
// Gibt true zurück, wenn mindestens ein Tausch stattgefunden hat.
func BubbleUp(arr []int) bool {

	for i := 0; i < len(arr)-1; i++ {

for i := 0; i < len(arr)-1; i++ {
        if arr[i] > arr[i+1] {
            tausch := arr[i]
            arr[i] = arr[i+1]
            arr[i+1] = tausch

            return true
        }
    }
    return false
}


// BubbleSort sortiert die übergebene Liste mittels des Bubble-Sort-Algorithmus.
func BubbleSort(arr []int) {
	if !BubbleUp(arr) {
		return
	}

}

//sortiert nicht
