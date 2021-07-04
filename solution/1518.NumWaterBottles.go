package solution

func NumWaterBottles(numBottles int, numExchange int) int {
	res := numBottles
	for numBottles >= numExchange {
		ex := numBottles / numExchange
		res += ex
		numBottles = ex + numBottles%numExchange
	}
	return res
}
