type Cell struct {
    R,C, Time int
}

type PQ []*Cell

func(p PQ) Len() int {
    return len(p)
}

func (p PQ) Less(i, j int) bool {
    return p[i].Time < p[j].Time
}

func (p PQ) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}

func (p *PQ) Pop() interface{} {
    n := len(*p)
    top := (*p)[n-1]
    *p = (*p)[:n-1]
    return top
}

func (p *PQ) Push(data interface{}) {
    *p = append(*p, data.(*Cell))
} 

func minimumTime(grid [][]int) int {
    n := len(grid)
    m := len(grid[0])

    if grid[0][1] > 1 && grid[1][0] > 1 {
        return -1
    }

    dp := make([][]int, n)
    
    for i := 0; i<n; i++ {
        dp[i] = make([]int, m)
        for j:= 0 ; j<m; j++ {
            dp[i][j] = 1<<32
        }
    }

    dp[0][0] = 0

    pq := make(PQ, 0)

    heap.Push(&pq, &Cell{R:0, C:0, Time:0})

    adjacent := [][]int{{0,1}, {1, 0}, {0, -1}, {-1, 0}}

    for len(pq) > 0 {
        top := heap.Pop(&pq).(*Cell)
        //fmt.Println(top)
        if top.R == n-1 && top.C == m-1 {
            return top.Time
        }

        // if dp[top.R][top.C] != -1 {
        //     continue
        // }
        //dp[top.R][top.C] = top.Time
        for _, adj := range adjacent {
            nr := top.R + adj[0]
            nc := top.C + adj[1]
            nt := top.Time + 1

            if nr < 0 || nr >= n || nc < 0 || nc >= m {
                continue
            }

            if grid[nr][nc] > nt {
                nt = grid[nr][nc]
                if (grid[nr][nc] - top.Time) % 2 == 0 {
                     nt += 1
                }
            }

            if dp[nr][nc] > nt {
                dp[nr][nc] = nt
                heap.Push(&pq, &Cell{R: nr, C: nc, Time: nt})
            }
        }
    }
    return -1
}
