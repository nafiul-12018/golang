
type Ch struct {
    Chr rune
    Cnt int
}

type PQ []*Ch

func(p *PQ) Len() int {
    return len(*p)
}

func(p PQ) Less(i, j int) bool {
    return p[i].Cnt > p[j].Cnt
}

func(p *PQ) Swap(i, j int) {
    (*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}


func(p *PQ) Push(data interface{}) {
    item := data.(*Ch)
    *p = append(*p, item) 
}

func(p *PQ) Pop() interface{} {
    n := len(*p)
    data := (*p)[n-1]
    *p = (*p)[:n-1]
    return data
}


func longestDiverseString(a int, b int, c int) string {
    pq := make(PQ, 0)
    if a > 0 {
        heap.Push(&pq, &Ch{
            Chr : 'a',
            Cnt : a,
        })
    }

    if b > 0 {
        heap.Push(&pq, &Ch{
            Chr : 'b',
            Cnt : b,
        })
    }

    
    if c > 0 {
        heap.Push(&pq, &Ch{
            Chr : 'c',
            Cnt : c,
        })
    }

    var ans string
    var arr []rune
    for pq.Len() > 0{
        top := heap.Pop(&pq).(*Ch)
        n := len(arr)
        if  n > 1 && arr[n-1] == arr[n-2] && arr[n-1] == top.Chr {
            if pq.Len() == 0 {
                break
            }

            top2 := heap.Pop(&pq).(*Ch)

            top2.Cnt--;
            arr = append(arr, top2.Chr)

            if(top2.Cnt > 0) {
                heap.Push(&pq, top2)
            }

            heap.Push(&pq, top)
        } else {
            arr = append(arr, top.Chr)

            top.Cnt--;
            if top.Cnt > 0 {
                heap.Push(&pq, top)
            }
        }

        ans = string(arr)
    }
    return ans 
}
