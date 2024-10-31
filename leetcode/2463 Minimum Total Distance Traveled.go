func solve(i, j int, robots *[]int, factory *[][]int, dp *[][]int64) int64 {
    
    if i == len(*robots) {
        return 0
    }

    if j == len(*factory) {
        return 1e12
    }

    if (*dp)[i][j] != -1 {
        return (*dp)[i][j]
    }

    dif := 0
    var ret int64 = 1e12
    ret = min(ret, solve(i, j+1, robots, factory, dp))
    for k:=0; k < (*factory)[j][1] && i+k < len(*robots); k++ {
        dif += abs((*factory)[j][0] - (*robots)[i+k])
        ret = min(ret, int64(dif) + solve(i+k+1, j+1, robots, factory, dp))
    }

    (*dp)[i][j] = ret
    return ret
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func min(x,y int64) int64 {
    if x < y {
        return x
    }
    return y
}

func minimumTotalDistance(robot []int, factory [][]int) int64 {
    n := len(robot)
    m := len(factory)
    sort.Ints(robot)
    sort.Slice(factory, func(i, j int) bool {
        return factory[i][0] < factory[j][0]
    })
    var dp [][]int64 = make([][]int64, n)
    for i := range dp {
        dp[i] = make([]int64, m)
        for j := range dp[i] {
            dp[i][j]=-1
        }
    }

    return solve(0,0, &robot, &factory, &dp)
    
}
