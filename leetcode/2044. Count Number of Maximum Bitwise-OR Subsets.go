package leetcode

func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	m := (1 << n)
	mx := 0
	for _, v := range nums {
		mx |= v
	}
	var ans int
	for i := 1; i < m; i++ {
		cur := 0
		for j := 0; j < n; j++ {
			if (i & (1 << j)) > 0 {
				//tmp = append(tmp, nums[j])
				cur |= nums[j]
			}
		}
		if cur == mx {
			ans++
		}
	}
	return ans

}
