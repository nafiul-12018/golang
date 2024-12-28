/*
 1.  2       1      2.     6.     7.      5.    1
    [[3 0]  [3 0]  [3 0]  [8 3]  [13 4] [13 4] [13 4]]
    [[3 0]  [3 1]  [3 2]  [8 3]  [-1 4] [12 5] [6 6]]
    [[12 5] [12 5] [12 5] [12 5] [12 5] [12 5] [6 6]]
                     |
                     3   
  1.  2       1      2.     6.     7.      5.    1
    [[3 0]  [3 1]  [3 2]  [8 3]  [13 4] [12 5] [6 6]]
    [[3 0]  [3 0]  [3 0]  [8 3]  [13 4] [13 4] [13 4]]
    [[13 4] [13 4] [13 4] [13 4] [13 4] [12 5] [6 6]]
  l
*/

func maxSumOfThreeSubarrays(nums []int, k int) []int {
    sum := 0;
    n := len(nums)
    prefixSum := make([]int, n)
    for idx, val := range nums {
        sum += val
        prefixSum[idx] = sum
    }

    left := make([][]int, n)
    right := make([][]int, n)
    val := 0
    pos := -1
    for i := 0; i<n; i++ {
        if i-k+1 >= 0 {
            cur := prefixSum[i]
            if i-k >= 0 {
                cur -= prefixSum[i-k]
            }
            if val < cur {
                val = cur
                pos = i-k+1
            }
        }
        left[i] = []int{val, pos}
    }

    val = -1
    for i := n-1; i>=0; i-- {
        if i + k - 1 < n {
            cur := prefixSum[i+k-1]
            if i > 0 {
                cur -= prefixSum[i-1]
            }
            if val <= cur {
                val = cur
                pos = i
            }
        }
        right[i] = []int{val, pos}
    }
    var ans []int

    fmt.Println(left)
    fmt.Println(right)
    val = -1
    for i:= k; i<n-k; i++ {
        l := left[i-1][0]
        m := prefixSum[i+k-1] - prefixSum[i-1]
        r := right[i+k][0] 
        if val < l+m+r {
            val = l+m+r
            ans = []int{left[i-1][1], i, right[i+k][1]}
        }
    }
    

    return ans
}
