func leftmostBuildingQueries(heights []int, queries [][]int) []int {
    for idx, _ := range queries {
        if queries[idx][0] > queries[idx][1] {
            queries[idx][0], queries[idx][1] = queries[idx][1], queries[idx][0]
        }
        queries[idx] = append(queries[idx], idx)
    }

    sort.Slice(queries, func (i, j int) bool {
        return queries[i][1] < queries[j][1]  
    })

    n := len(queries)
    ans := make([]int, n)
    stack := make([]int, 0)
    j := n-1
    for i := len(heights)-1; i >= 0; i-- {
        for len(stack) > 0 && heights[stack[len(stack)-1]] <= heights[i] {
            stack = stack[: len(stack)-1]
        }

        stack = append(stack, i)

        for j>=0 && queries[j][1] == i {
            if queries[j][1] == queries[j][0] {
                ans[queries[j][2]] = i
                j--
                continue
            }

            val := heights[queries[j][1]]
            if val <= heights[queries[j][0]] {
                val = heights[queries[j][0]] + 1
            }

            low := 0
            high := len(stack)-1
            idx := -1
            for low <= high {
                md := (low + high)/2

                if heights[stack[md]] >= val {
                    idx = stack[md]
                    low = md +1
                } else {
                    high = md - 1
                }
            }
            
            ans[queries[j][2]] = idx
            j--
        }
    }

    return ans
}
