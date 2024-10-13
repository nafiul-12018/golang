type PQ [][]int

func (p *PQ) Len() int {
    return len(*p)
}

func (p *PQ) Swap(i, j int) {
    (*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PQ) Less(i, j int) bool {
    return (*p)[i][0] < (*p)[j][0]
}

func (p *PQ) Pop() interface{} {
    n := len(*p)
    item := (*p)[n-1]
    *p = (*p)[: n-1]
    return item
}

func (p *PQ) Push(data interface{}) {
    item := data.([]int)
    *p = append(*p, item)
}

func smallestRange(nums [][]int) []int {
    pq := make(PQ, 0)
    mx := -10000005
    for idx, data := range nums {
        heap.Push(&pq, []int{data[0], idx, 0})
        if mx < data[0] {
            mx = data[0]
        }
    }
    //fmt.Println(mx)
    ans := []int{0,0}
    dif := 20000004
    for pq.Len() > 0 {
        front := heap.Pop(&pq).([]int)
        //fmt.Println(front)
        cur := mx - front[0]
        //fmt.Println(cur, mx, front[0])
        if dif > cur {
            dif = cur
            ans = []int{front[0], mx}
        }
        n := len(nums[front[1]])
        if front[2] == n-1 {
            break
        }
        if mx < nums[front[1]][front[2]+1] {
            mx = nums[front[1]][front[2]+1]
        }
        heap.Push(&pq, []int{nums[front[1]][front[2]+1], front[1], front[2]+1})

    }
    return ans 
}
