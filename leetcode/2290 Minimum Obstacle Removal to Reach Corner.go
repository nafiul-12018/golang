type Cell struct {
    R, C, Cost int
}

type List []*Cell

func (l List) Len() int {
    return len(l)
}

func (l List) Swap(i, j int) {
    l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
    return l[i].Cost < l[j].Cost
}

func (l *List) Pop() interface{} {
    n := len(*l)
    top := (*l)[n-1]
    *l = (*l)[:n-1]
    return top
}

func (l *List) Push(x interface{}) {
    data := x.(*Cell)
    *l = append(*l, data)
}

func minimumObstacles(grid [][]int) int {
    n := len(grid)
    m := len(grid[0])
    //fmt.Println(n, m)
    list := make(List,0)
    heap.Push(&list, &Cell{R:0, C:0, Cost:0,})

    adjacent := [][]int{{0,1}, {1,0}, {0, -1}, {-1, 0}}
    
    mp := make([][]int, n)
    for i:=0; i<n; i++ {
        mp[i] = make([]int, m)
        for j:=0; j<m; j++ {
            mp[i][j] = n*m
        }
    }

    mp[0][0]=0
    //mp["0#0"] = 0
    for len(list) > 0 {
        item := heap.Pop(&list).(*Cell)

        if item.R == n-1 && item.C == m-1 {
            return item.Cost
        }
        for _, adj := range adjacent {
            nr := item.R + adj[0]
            nc := item.C + adj[1]
            nCost := item.Cost

            if nr < 0 || nr >= n || nc < 0 || nc >= m {
                continue
            }

            if grid[nr][nc] == 1 {
                nCost++
            }

            //key := fmt.Sprintf("%d#%d", nr, nc)

            //val, ok := mp[key]

            if mp[nr][nc] > nCost {
                mp[nr][nc] = nCost
            } else {
                continue
            }

            heap.Push(&list, &Cell{R:nr, C:nc, Cost: nCost,})
        }

    }
    return 0
}
