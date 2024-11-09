
func minEnd(n int, x int) int64 {
    var ans int64 = 0
    var data []bool
    
    for x > 0 {
        if x & 1 == 1 {
            data = append(data, true)
        }else {
            data = append(data, false)
        }

        x = x>>1
    }
    fmt.Println(data)
    idx := 0
    n--
    for n > 0 {
        if len(data) == idx {
            data = append(data, false)
        }

        if data[idx] == true {
            idx++
            continue
        }

        if n & 1 == 1 {
           data[idx] = true
        }else {
           data[idx] = false
        }
        n = n>>1
        idx++
    }    
    for i:= 0; i<len(data); i++ {
        if data[i] == true {
            ans |= (1<<i)
        }
    }
    return ans   
}
