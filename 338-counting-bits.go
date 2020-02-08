// 338. Counting Bits
// Leetcode Medium done in Golang
// Runtime 4ms (faster than 100%)
// Memory usage: 6.1 MB (less than 100%)

func countBits(num int) []int {
    
    var bits []int
    
    if num >= 0 {
        bits = append(bits, 0)
    }
    
    if num >= 1 {
        bits = append(bits, 1)
    }
    
    var pow int = 1
    
    for i:= 2; i <= num; i++ {
        if i % pow == 0 {
            bits = append(bits, 1)
            pow *= 2
        } else {
            var pos int = i - pow
            bits = append(bits, bits[pos] + 1)
        }
    }
    
    return bits
    
}