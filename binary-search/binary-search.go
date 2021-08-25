func search(nums []int, target int) int {
    
    for key, value := range nums {
        
        if value == target {
            
            return key
            
        }
    }
    
    return -1
    
}