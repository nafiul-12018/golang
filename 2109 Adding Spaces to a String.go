func addSpaces(s string, spaces []int) string {

    ans := make([]rune, 0)

    i := 0;
    for idx, ch := range s {
        if i < len(spaces) && idx == spaces[i] {
            ans = append(ans, ' ')
            i++
        }

        ans = append(ans, ch) 
    }

    return string(ans)
}
