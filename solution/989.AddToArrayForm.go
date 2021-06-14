package solution

func array2Integer(num []int) int {
	integer := 0
	index := 0
	for {
		integer += num[index]
		if index != len(num)-1 {
			integer *= 10
			index++
		} else {
			break
		}
	}
	return integer
}

func integer2Array(num int) []int {
	array := make([]int, 0)
	if num == 0 {
		return []int{0}
	}
	for num > 0 {
		array = append(array, num%10)
		num /= 10
	}
	for i, j := 0, len(array)-1; i < j; {
		array[i], array[j] = array[j], array[i]
		i++
		j--
	}
	return array
}

func AddToArrayFormWithOutOverFlow(num []int, k int) []int {
	return integer2Array(array2Integer(num) + k)
}

func arrayAdd(a, b []int) []int {
	sum := make([]int, 0)
	for i, j, o := len(a)-1, len(b)-1, 0; ; {
		x, y := 0, 0
		if i >= 0 {
			x = a[i]
		}
		if j >= 0 {
			y = b[j]
		}
		z := x + y + o
		o = z / 10
		z = z % 10
		sum = append(sum, z)
		i--
		j--
		if i < 0 && j < 0 {
			if o != 0 {
				sum = append(sum, o)
			}
			break
		}
	}
	for i, j := 0, len(sum)-1; i < j; {
		sum[i], sum[j] = sum[j], sum[i]
		i++
		j--
	}
	return sum
}

func AddToArrayForm(num []int, k int) []int {
	return arrayAdd(num, integer2Array(k))
}
