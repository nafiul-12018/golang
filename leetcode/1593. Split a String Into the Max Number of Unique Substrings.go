func dfs(pos int, s string, mp map[string]bool) int {
    if len(s) == pos {
        //fmt.Println(mp)
        return len(mp)
    }
    var ret int = 1
    for i := pos; i<len(s); i++ {
        tmp := s[pos: i+1]
        //fmt.Println(tmp)
        _, ok := mp[tmp]
        if !ok {
            mp[tmp] = true
            cur := dfs(i+1, s, mp)
            if cur > ret {
                ret = cur
            }
            delete(mp, tmp)
        } 
    }
    return ret
}
func maxUniqueSplit(s string) int {
    mp := map[string]bool{}
    return dfs(0, s, mp)
}
