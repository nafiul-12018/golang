func longestSquareStreak(nums []int) int {
    sort.Ints(nums)
    ans := 0
    mp := make(map[int]int)
    for _, val := range nums {
        mp[val] = 1
    }

    for _, val := range nums {
        v, ok := mp[val*val]
        if ok {
            if mp[val] >= v {
                mp[val*val] = mp[val]+1 
            }
        }
        
        if mp[val] > ans {
            ans = mp[val]
        }
    }
    if ans < 2 {
        ans = -1
    }
    //fmt.Println(mp)
    return ans
}
