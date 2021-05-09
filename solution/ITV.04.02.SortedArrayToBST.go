package solution

func construct(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	if l == r {
		return &TreeNode{nums[l], nil, nil}
	}
	mid := l + (r-l)/2
	left := construct(nums, l, mid-1)
	right := construct(nums, mid+1, r)
	return &TreeNode{nums[mid], left, right}
}

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return construct(nums, 0, len(nums)-1)
}
