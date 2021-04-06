package solution

func CountMatches(items [][]string, ruleKey, ruleValue string) int {
	res := 0
	index := -1
	switch ruleKey {
	case "type":
		index = 0
	case "color":
		index = 1
	case "name":
		index = 2
	}
	if index == -1 {
		return res
	}
	for i := range items {
		if items[i][index] == ruleValue {
			res++
		}
	}
	return res
}
