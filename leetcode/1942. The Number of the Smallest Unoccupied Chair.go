/*

*/
type Slot struct {
    Start int
    End int
    Id int
    ChairId int
}


type PQ []*Slot

func(q PQ) Len() int {
    return len(q)
}

func(q PQ) Less (i, j int) bool {
    if q[i].End == q[j].End {
        return q[i].Start > q[j].Start
    }
    return q[i].End < q[j].End
}

func(q *PQ) Push (x interface{}) {
    item := x.(*Slot)
    *q = append(*q, item)
    //*pq = append(*pq, item)
}

func(q *PQ) Pop () any {
    n := len(*q)
    item := (*q)[n-1]
    *q = (*q)[:n-1]
    return item
}

func(q *PQ) Top () any {
   // n := len(*q)
    item := (*q)[0]
    return item
}

func(q *PQ) Swap (i, j int) {
    //q[i],q[j] = q[j], q[i]
    (*q)[i], (*q)[j] = (*q)[j], (*q)[i]    
}


type PQ2 []int

func(q PQ2) Len() int {
    return len(q)
}

func(q PQ2) Less (i, j int) bool {
    return q[i] < q[j]
}

func(q *PQ2) Push (x interface{}) {
    item := x.(int)
    *q = append(*q, item)
    //*pq = append(*pq, item)
}

func(q *PQ2) Pop () any {
    n := len(*q)
    item := (*q)[n-1]
    *q = (*q)[:n-1]
    return item
}

func(q *PQ2) Swap (i, j int) {
    //q[i],q[j] = q[j], q[i]
    (*q)[i], (*q)[j] = (*q)[j], (*q)[i]    
}


func smallestChair(times [][]int, targetFriend int) int {
    var arr []*Slot
    for idx, data := range times {
        item := &Slot{
            Start : data[0],
            End : data[1],
            Id : idx,
        }
        arr = append(arr, item)        
    }

    sort.Slice(arr, func(i,j int) bool {
        if arr[i].Start == arr[j].Start {
            return arr[i].Id > arr[i].Id
        }

        return arr[i].Start <= arr[j].Start 
    })

    

    q := make(PQ, 0)
    q2 := make(PQ2, 0)
    var cur int = 0
    for _, data := range arr {
        //heap.Push(&q, data)
        for q.Len() > 0 && q.Top().(*Slot).End <= data.Start  {
            it := heap.Pop(&q).(*Slot)
            heap.Push(&q2, it.ChairId)
        }
        id := cur
        if q2.Len() > 0 {
            id = heap.Pop(&q2).(int)
        } else {
            cur++;
        }

        if data.Id == targetFriend {
            return id;
        }

        data.ChairId = id

        heap.Push(&q, data);
        
    }

    return 0
}
