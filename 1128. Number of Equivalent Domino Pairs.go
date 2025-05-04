func numEquivDominoPairs(dominoes [][]int) int {
    mp:= make([]int, 100)
    ans := 0
    for _, val := range dominoes {
        if val[0] > val[1] {
            val[0], val[1] = val[1], val[0]
        }
        //hash := fmt.Sprintf("%d:%d", val[0], val[1])
        hash := val[0]*10 + val[1]
        cnt := mp[hash]

        if cnt > 0 {
            ans += cnt
        }

        mp[hash]++;
    }

    return ans;
}
