
func getMaximumXor(nums []int, maximumBit int) []int {
    bit := 0
    n := len(nums)
    mx := (1<<maximumBit) - 1 
    for _, val := range nums {
        bit ^= val
    }
    ans := make([]int, 0)
    for i:=0; i<n; i++ {
        cur := mx ^ bit
        ans = append(ans, cur)
        val := nums[n-i-1]

        bit = bit ^ val
    }
    return ans
}
