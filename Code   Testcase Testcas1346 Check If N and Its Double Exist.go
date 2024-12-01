func checkIfExist(arr []int) bool {
    mp := make(map[int]bool, 0)
    for _, val := range arr {
        if val % 2 == 0 && mp[val/2] == true {
            return true
        } 

        if mp[val*2] == true {
            return true
        } 
        mp[val] = true
    } 
    return false
}
