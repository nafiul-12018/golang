func getCnt(nums *[]int, val int) int64 {
    i:= 0

    var ans int64 = 0
    j := len(*nums) -1
    for i < j {
        //fmt.Println(i, j)
        if i< j-1 && ((*nums)[i+1] + (*nums)[j] <= val) {
            i++
            continue
        }
        if ((*nums)[i] + (*nums)[j] <= val) {
            ans += int64(i+1)
            //fmt.Println("--", i,j, ans)
        }
        j--
    } 
    if j > 0 {
       ans += int64((j*(j+1))/2)
    }
    return ans
}

func countFairPairs(nums []int, lower int, upper int) int64 {
    sort.Ints(nums)
    //fmt.Println(nums)
    up := getCnt(&nums, upper)
    dw := getCnt(&nums, lower-1)

    return up-dw
}
