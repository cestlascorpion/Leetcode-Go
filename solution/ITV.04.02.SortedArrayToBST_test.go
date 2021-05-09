package solution

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestSortedArrayToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	tree := SortedArrayToBST(nums)
	ch := make(chan int)

	wg.Add(2)

	go func() {
		defer wg.Done()
		LDR(tree, ch)
	}()

	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()
	wg.Wait()
}
