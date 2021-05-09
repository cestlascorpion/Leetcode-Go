package solution

func SumOddLengthSubArrays(arr []int) int {
	sum := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		sum += arr[i] * ((i/2+1)*((n-i-1)/2+1) + ((i+1)/2)*((n-i)/2))
	}
	return sum
}
