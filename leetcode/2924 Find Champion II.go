

func findChampion(n int, edges [][]int) int {
    deg := make([]int, n)

    for _, item := range edges {
        deg[item[1]]++
    }

    root := -1
    for i:=0; i<n; i++ {
       if deg[i] == 0 {
            if root != -1 {
                return -1
            }
            root = i
       }
    }

    return root
}
