package solution

import "sort"

func MaximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	sum := 0
	for i := range boxTypes {
		num := boxTypes[i][0]
		size := boxTypes[i][1]
		if truckSize >= num {
			sum += size * num
			truckSize -= num
		} else {
			sum += size * truckSize
			break
		}
	}
	return sum
}
