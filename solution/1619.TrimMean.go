package solution

func TrimMean(arr []int) float64 {
	k := len(arr) / 20
	// fmt.Println(k)
	// big K
	for i := 0; i < k; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	// fmt.Println(arr)
	for i := 0; i < k; i++ {
		for j := 0; j < len(arr)-i-1-k; j++ {
			if arr[j] < arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	// fmt.Println(arr)
	sum := 0
	for i := 0; i < len(arr)-k*2; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(len(arr)-k*2)
}
