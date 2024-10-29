/*
[
    [3,2,4],
    [2,1,9],
    [1,1,7]
]

*/
func maxMoves(grid [][]int) int {
    n := len(grid)
    m := len(grid[0])
    var dp [][]int = make([][]int, n)

    for i := 0; i<n; i++ {
        dp[i] = make([]int, m)
        dp[i][m-1] = 1
    }
    ans := 0
    for i := m-2; i>= 0; i-- { //col
        for j := 0; j<n; j++ { //row
            var ret int = 0

            if grid[j][i] < grid[j][i+1] {
                ret = dp[j][i+1]
            }
            //fmt.Println(ret)
            if j > 0 && grid[j][i] < grid[j-1][i+1] {
                ret = max(ret, dp[j-1][i+1])
            }
           // fmt.Println(ret)

            if j < n-1 && grid[j][i] < grid[j+1][i+1] {
                ret = max(ret, dp[j+1][i+1])
            }
            //fmt.Println(ret)


            dp[j][i]= ret+1
            if i == 0 {
                ans = max(ans, ret+1)
            }
            //fmt.Println(j, i, dp[j][i])
            
        }
    }
    return ans -1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
} 
