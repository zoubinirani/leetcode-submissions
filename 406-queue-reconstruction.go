// 406. Queue Reconstruction by Height
// Leetcode Medium done in Golang
// Runtime 12ms (faster than 100%)
// Memory usage: 5.9 MB (less than 100%)
// Anonymous function used in sort.Slice taken from Leaf_Peng's solution (https://leetcode.com/problems/queue-reconstruction-by-height/discuss/427846/Go-golang-clean-solution)

import "fmt"

func reconstructQueue(people [][]int) [][]int {
    
    if len(people) == 0 {
        return people
    }
        
    solution:= make([][]int, len(people))
    
    // anonymous function taken from Leaf_Peng's solution
    // https://leetcode.com/problems/queue-reconstruction-by-height/discuss/427846/Go-golang-clean-solution
    
    sort.Slice(people, func(i, j int) bool { 
                   return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
               })
    
    for _, v := range people {
        copy(solution[v[1]+1:len(people)], solution[v[1]:len(people)-1])
        solution[v[1]] = v   
    }
       
    return solution
    
}