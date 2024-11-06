func canSortArray(nums []int) bool {
    var arr []int
    bit := 0
    mn := nums[0]
    mx := nums[0]
    for _, v := range nums {
        cnt := 0
        m := v
        for v > 0 {
            if v & 1 == 1 {
                cnt++
            }
            v = v>>1
        }

        if bit == cnt {
            if m < mn {
                mn = m
            }

            if mx < m {
                mx = m
            }
        }else {
            if bit != 0 {
                arr = append(arr, mn)
                arr = append(arr, mx)
            }
            mn = m
            mx = m
            bit = cnt
        }

        
    }
    if bit != 0 {
        arr = append(arr, mn)
        arr = append(arr, mx)
    }

    for i:= 1; i<len(arr); i++ {
        if arr[i] < arr[i-1] {
            return false
        }
    }
    //fmt.Println(nums)
    //fmt.Println(arr)
    return true
}
