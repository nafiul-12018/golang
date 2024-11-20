func maximumSubarraySum(nums []int, k int) int64 {
    mp := map[int]int{}
    var sum int64 = 0
    var ans int64 = 0

    for idx,val := range nums {
        mp[val]++
        sum += int64(val)
        if (idx-k) >= 0 {
            sum -= int64(nums[idx-k])
            mp[nums[idx-k]]--

            if mp[nums[idx-k]] == 0 {
                delete(mp, nums[idx-k])
            }
        }

        if len(mp) == k {
            //fmt.Println(idx, mp, sum)
            if ans < sum {
                ans = sum
            }
        }   
    }
    return ans
}
