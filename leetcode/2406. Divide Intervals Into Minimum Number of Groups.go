type PQ [][]int

func (p *PQ) Len() int {
    return len(*p)
}
func (p PQ) Less(i, j int) bool {
    return p[i][0] < p[j][0]
}

func (p PQ) Swap(i, j int) {
    p[i],p[j] = p[j], p[i]
}

func (p *PQ) Push(data interface{}) {
    item := data.([]int)
    *p = append(*p, item);
}

func (p *PQ) Pop() interface{} {
    n := len(*p)
    data := (*p)[n-1]
    *p = (*p)[:n-1]
    return data
}

func (p *PQ) Top() interface{} {
    data := (*p)[0]
    return data
}


func minGroups(intervals [][]int) int {
    var data [][]int
    for _, item := range intervals {
        data = append(data, []int{item[0], 1})
        data = append(data, []int{item[1]+1, -1})
    }
    sort.Slice(data, func(i, j int) bool {
        if data[i][0] == data[j][0] {
            return data[i][1] < data[j][1]
        }  
        return data[i][0] < data[j][0]
    })


    //fmt.Println(intervals)

    grp := 0
    ans := 0
    for _, item := range data {
       grp += item[1]
       if ans < grp {
          ans = grp;
       } 
    }

    return ans
}
