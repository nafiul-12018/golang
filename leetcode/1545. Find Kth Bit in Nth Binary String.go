/*
1.  1
3.  2
7.  3
15. 4 
31
63
127
0111001 1 0110001
4 10

*/

func solve(n, k int) byte {
       var data []rune
    data = append(data, '0')
    n--
    for n > 0 {
       m := len(data) 
       data = append(data, '1')
       for j := m-1; j>=0; j-- {
            if data[j] == '0' {
                data = append(data, '1')
            } else {
                data = append(data, '0')
            }
       }
       //fmt.Println(string(data))

       n--
    }
    //fmt.Println(data)
    return byte(data[k-1])
}

func recursive(n, k int) bool {
    if n == 1 {
        return false
    }
    m := 1<<n
    if m/2 == k {
        return true
    }
    var ret bool
    if k < (m/2) {
        ret = recursive(n-1, k)
    } else {
        ret = !recursive(n-1, m-k)

    }
    return ret
}

func findKthBit(n int, k int) byte {
  
    //fmt.Println(data)
    var ans rune = '0'
    if recursive(n, k) {
        ans = '1'
    }
    return byte(ans)
}
