func maxChunksToSorted(arr []int) int {
        mp := make(map[int]int)

        for idx, val := range arr {
            mp[val] = idx
        }

        pre := -1
        ans := 0
        sort.Ints(arr)
        //fmt.Println(pre, arr)
        mxIdx := -1;
        for idx, val := range arr {
            if mp[val] > mxIdx {
                mxIdx = mp[val]
            }

            if (idx - pre) == (mxIdx - pre) {
                pre = idx
                ans++
            }           
        }
        return ans
}
