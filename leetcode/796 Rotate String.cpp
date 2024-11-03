/*
|
bbbacddceeb
ceebbbbacdd
     
*/
func rotateString(s string, goal string) bool {
    n:= len(s)
    if len(goal) != n {
        return false
    }
    arr := make([]int, n)

    i := 0
    j := 1

    for i < n && j < n {
        if s[i] == s[j] {
            i++
            arr[j]=i
        } else if i > 0 {
            i = arr[i-1]
            continue
        } else {
            arr[j] = i
        }
        j++;
    }
    //fmt.Println(arr)
    i = 0
    j = 0
    cnt := 0
    for i < n && j < n {
        //fmt.Println(i, j)

        if s[i] == goal[j] {
            i++
            if cnt < i {
                cnt = i
            }
        } else if i > 0 {
            i = arr[i-1]
            //fmt.Println(i," ___")
            continue
        } 
        j++;
    }

    if i == n {
        return true
    }

    //fmt.Println(i)

    j = 0
    for i < n {
        if s[i] != goal[j] {
            return false
        }
        i++
        j++
    }
    return true
}
