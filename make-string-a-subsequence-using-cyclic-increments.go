func dfs(i, j int, str1, str2 string) bool {
    
    if j == len(str2) {
        return true
    }

    if i == len(str1) {
        return false
    }

    ret := false
    ch := rune(str1[i])
    nch := ch+1
    if ch == 'z' {
        nch = 'a'
    }
    //fmt.Println(string(ch), string(nch), rune('z'), string(123))
    if ch == rune(str2[j]) || nch == rune(str2[j]) {
        ret = ret || dfs(i+1, j+1, str1, str2)
    } else {
        ret = ret || dfs(i+1, j, str1, str2)
    }
    return ret
}   

func canMakeSubsequence(str1 string, str2 string) bool {
   return dfs(0, 0, str1, str2)
}
