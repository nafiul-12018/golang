func countSquares(matrix [][]int) int {
    n := len(matrix)
    m := len(matrix[0])

    for i:=1; i<n; i++ {
        for j :=1; j<m; j++ {
            if matrix[i][j] == 0 {
                continue
            }

           matrix[i][j] = min(min(matrix[i-1][j-1], matrix[i-1][j]), matrix[i][j-1]) + 1
        } 
    }
    sum := 0
    for i:=0; i<n; i++ {
        for j :=0; j<m; j++ {
            if matrix[i][j] == 0 {
                continue
            }

           sum += matrix[i][j]
        } 
    }
    return sum
}

func min(x, y int ) int {
    if x < y {
        return x
    }
    return y
}
