func resultsArray(nums []int, k int) []int {
    
    if k == 1 {
        return nums
    }
    pre := 0
    n := len(nums)
    ans := make([]int, n - k+ 1)

    for idx, val := range nums {
        if idx == 0 {
            continue
        }

        if nums[idx-1] != (nums[idx]-1) {
            pre = idx
            continue
        } else if idx-pre+1 == k && pre <= (n-k) {
            ans[pre] = val
            //ans = append(ans, val)
            pre++
        }

    }

    for i:= 0; i<len(ans); i++ {
        if ans[i] == 0 {
            ans[i] = -1
        }
    }
    return ans
}
