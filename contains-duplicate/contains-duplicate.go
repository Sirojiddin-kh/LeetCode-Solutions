func containsDuplicate(nums []int) bool {
    
    dict := make(map[int]int)
    
    for _, value := range nums {
        
        dict[value] += 1
    }
    
    for _, value := range dict {
            
        if value >= 2 {
            
            return true
        }
    }
    return false
}