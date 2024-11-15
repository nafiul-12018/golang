func minimizedMaximum(n int, quantities []int) int {
    
    low := 1
    high := 0
    for _, item := range quantities {
        if high < item {
            high = item
        }
    }
    ans := high
    for low <= high {
        md := (low + high) / 2
        cnt := 0
        for _, item := range quantities {
            cnt += item / md
            if item % md > 0{
                cnt += 1
            } 
        }

        if cnt <= n {
            ans = md
            high = md - 1
        } else {
            low = md + 1
        }
    }
    return ans
}
