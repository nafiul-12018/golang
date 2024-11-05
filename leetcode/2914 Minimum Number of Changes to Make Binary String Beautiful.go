/*
1001
1001001110

*/

func minChanges(s string) int {
    n := len(s)
    ans := 0
    for i := 1; i<n; i+=2 {
        if s[i] != s[i-1] {
            ans++
        }
    }

    return ans
}
