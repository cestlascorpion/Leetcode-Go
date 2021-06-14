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

func FromTopToBottom(root *TreeNode) [][]*TreeNode {
	res := make([][]*TreeNode, 0)
	if root == nil {
		return res
	}
	level := make([]*TreeNode, 0, 1)
	level = append(level, root)
	for len(level) != 0 {
		nextLevel := make([]*TreeNode, 0, len(level)*2)
		for i := range level {
			if level[i].Left != nil {
				nextLevel = append(nextLevel, level[i].Left)
			}
			if level[i].Right != nil {
				nextLevel = append(nextLevel, level[i].Right)
			}
		}
		res = append(res, level)
		level = nextLevel
	}
	return res
}
