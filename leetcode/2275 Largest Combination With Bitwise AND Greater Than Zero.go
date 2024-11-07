func largestCombination(candidates []int) int {
    var bit [32]int
    ans := 0
    for _, v:= range candidates {
        idx := 0
        for v > 0 {
            if (v & 1) == 1 {
                bit[idx]++
                if bit[idx] > ans {
                    ans = bit[idx]
                }
            }

            v = v>>1
            idx++
        }
    }

    return ans
}
