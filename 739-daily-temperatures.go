// 739. Daily Temperatures
// Leetcode Medium done in Golang
// Runtime 500ms (faster than 99.5%)
// Memory usage: 6.6 MB (less than 100%)


func dailyTemperatures(T []int) []int {
    
	var lenT int = len(T)
	
	solution:= make([]int, lenT)
	
	for i := lenT-2; i >= 0; i-- {
			var numDays int = 0
			for _, v := range T[i+1:] {
					numDays += 1
					if v > T[i] {
							break
					} 
			}
			if numDays == lenT-1-i {
					if T[lenT-1] <= T[i] {
							numDays = 0
					}
			}
			solution[i] = numDays 
	}
	
	return solution   
}