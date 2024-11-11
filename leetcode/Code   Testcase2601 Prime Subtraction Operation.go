func primeSubOperation(nums []int) bool {
   mx := 0
   for _, v := range nums {
        if v > mx {
            mx = v
        }
   }
   var prime []int
   mask := make([]int, mx+1) 
   for i := 2 ; i <= mx; i++ {
        if mask[i] == 1 {
            continue
        }

        prime = append(prime, i)

        for j := i+i; j <= mx; j+= i {
            mask[j] = 1
        }
   }

    pre := -1
    for _, val := range nums {
        v := binsearch(&prime, val)

       // fmt.Println(val, v)
        if v == -1 {
            if pre >= val {
                return false
            }
            pre = val
            continue
        }

        for v >=0 && pre >= val-prime[v] {
            v--
        }

        if v >=0 {
            pre = val-prime[v]
        } else if pre < val {
            pre = val
        } else {
            return false
        }
    }
    return true
}

func binsearch(data *[]int, val int) int {
    ans := -1

    low := 0
    high := len(*data) - 1
    //fmt.Println(low, high,"---")
    for low <= high {
        md := (low + high) / 2
        if (*data)[md] < val {
            ans = md
            low = md + 1
        } else {
            high = md - 1
        }   
    }
    return ans
}
