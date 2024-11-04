func compressedString(word string) string {
     var ch rune = ' '
     var cnt int
     var ans []rune
     for _, c  := range word {
        if ch == c && cnt < 9{
            cnt++
        } else {
            if cnt > 0 {
                ans = append(ans, rune('0' + cnt))
                ans = append(ans, ch)
            }
            ch = c
            cnt = 1
        }
     }

     if cnt >= 1 {
        ans = append(ans, rune('0' + cnt))
        ans = append(ans, ch)
     }
     return string(ans) 
}
