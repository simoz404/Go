func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	for i, v := range nums {
		for i1, v1 := range nums {
			if i1 != i && v == v1 {
				break
			} else if i1 == len(nums)-1 {
				return v
			}
		}
	}
	return 0
}