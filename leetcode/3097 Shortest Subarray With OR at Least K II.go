
func add(x int, cnt *[]int) {
    idx := 0
    for x > 0 {
        if x&1 == 1 {
            (*cnt)[idx]++
        }

        x = x>>1
        idx++
    }
}

func check(x int, cnt *[]int) bool {
    val := 0
    for i:= 0; i<32; i++ {
        if (*cnt)[i] > 0 {
            val |= 1<<i
        }
    }
    return val>=x
}

func remove(x int, cnt *[]int) {
    idx := 0
    for x > 0 {
        if x&1 == 1 {
            (*cnt)[idx]--
        }

        x = x>>1
        idx++
    }
}



func minimumSubarrayLength(nums []int, k int) int {
    n := len(nums)
    cnt := make([]int, 32)
    i := 0
    ans := -1
    for j:=0; j<n; j++ {
        add(nums[j], &cnt)

        for i<=j && check(k, &cnt) {
            if (j-i+1) < ans || ans == -1 {
                ans = j-i+1
            }

            remove(nums[i], &cnt)
            i++
        }
    }

    return ans
}
