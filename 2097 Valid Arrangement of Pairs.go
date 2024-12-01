func dfs(root int, graph map[int][]int, path *[]int) {
    for len(graph[root]) > 0 {
        v := graph[root][len(graph[root])-1]
        graph[root]= graph[root][:len(graph[root])-1]
        dfs(v, graph, path)
    }   

    *path = append(*path, root)
}

func validArrangement(pairs [][]int) [][]int {
    mp := make(map[int][]int, 0)
    cnt := make(map[int]int, 0)
    for _, data := range pairs {
        //mp[data[0]] = append(mp[data[0]], []int{idx,data[1]})
        if _, exists := mp[data[0]]; !exists {
            mp[data[0]] = make([]int, 0)
        }
        mp[data[0]] = append(mp[data[0]], data[1])
        
        cnt[data[1]]++
        cnt[data[0]]--
    }

    start := pairs[0][0]

    //fmt.Println(cnt)
    for k,v := range cnt {
        if v == -1 {
            start = k
            break
        }
    }
    

    data := make([]int, 0)
    //fmt.Println(mp)

    dfs(start, mp, &data)
    //fmt.Println("**",data)
    var ans [][]int

    for i := len(data)-1 ; i > 0; i-- {
        ans = append(ans, []int{data[i], data[i-1]})
    }
