func isCircularSentence(sentence string) bool {
    n := len(sentence)
    if sentence[0] != sentence[n-1] {
        return false
    }

    for i := 1; i<n-1; i++ {
        if sentence[i] == ' '  {
            if sentence[i-1] != sentence[i+1] {
                return false
            }
        }
    }

    return true
}
