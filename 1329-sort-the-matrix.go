// 1329. Sort the Matrix Diagonally
// Leetcode Medium done in Golang
// Runtime: 8 ms (faster than 100% of previous submissions)
// Memory usage: 6 MB (less than 100% of previous submissions)

import "sort"
import "fmt"

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func diagonalSort(mat [][]int) [][]int {
    
    var nCol int = len(mat[0])
    var nRow int = len(mat)
    
    // fmt.Println(nCol, nRow)
    
    type position struct {
        x int
        y int
    }
    
    
    for i := nRow; i >= 0; i-- {
        sortArray := make([]int, 0)
        sortPositions := make([]position, 0)
        
        var jDist int = min(nRow - i, nCol)
        for j:= 0; j < jDist; j++ { 
            sortPositions = append(sortPositions, position{i + j, j})
        }
               
        for _, pos := range sortPositions {
            sortArray = append(sortArray, mat[pos.x][pos.y])
        }
        
        sort.Ints(sortArray)
        
        for i, pos := range sortPositions {
            mat[pos.x][pos.y] = sortArray[i]
        }
    }
    
    for i := 1; i < nCol; i++ {
        sortArray := make([]int, 0)
        sortPositions := make([]position, 0)
        
        var jDist int = min(nCol - i, nRow)
        for j:= 0; j < jDist; j++ { 
            sortPositions = append(sortPositions, position{j, i + j})
        }
               
        for _, pos := range sortPositions {
            sortArray = append(sortArray, mat[pos.x][pos.y])
        }
        
        sort.Ints(sortArray)
        
        for i, pos := range sortPositions {
            mat[pos.x][pos.y] = sortArray[i]
        }
    }    


    return mat   
}