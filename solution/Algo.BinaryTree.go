package solution

import (
	"fmt"
	"sync"
)

func ldr(node *TreeNode, ch chan<- int) {
	if node.Left != nil {
		ldr(node.Left, ch)
	}
	ch <- node.Val
	if node.Right != nil {
		ldr(node.Right, ch)
	}
}

func LDR(node *TreeNode, ch chan<- int) {
	defer close(ch)
	ldr(node, ch)
}

func LDR2(node *TreeNode) {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		LDR(node, ch)
	}()

	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()
	wg.Wait()
}
