type PQ []int
func(p *PQ) Len() int {
    return len(*p)
}

func(p *PQ) Pop() interface{} {
    n := len(*p)
    data := (*p)[n-1]
    *p = (*p)[:n-1]
    return data
}

func(p *PQ) Push(data interface{}) {
    item := data.(int)
    *p = append(*p, item)
}

func(p PQ) Swap(i, j int) {
    p[i],p[j] = p[j],p[i]
}


func(p *PQ) Less(i, j int) bool {
    return (*p)[i] > (*p)[j]
}

func maxKelements(nums []int, k int) int64 {
    pq := make(PQ, 0)
    for _, data := range nums {
        heap.Push(&pq, data)
    }
    var ans int64 = 0
    cnt := 0
    for cnt < k {
        cnt++
        cur := heap.Pop(&pq).(int)
        //fmt.Println(cur)
        ans += int64(cur)
        var sc int = 0
        if cur % 3 > 0 {
            sc++
        }

        sc += (cur/3)
        //fmt.Println("   ", sc)
        heap.Push(&pq, sc)
        //fmt.Println(pq)
    }

    return ans
}
