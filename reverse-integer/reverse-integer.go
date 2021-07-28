func reverse(x int) int {
    var result, num int
    num = x
    
    if x == 0  { 
        
        return 0
    }
    
    if x < 0 {
        num = -x 
    } 
    
    for num > 0 {
        
        reminder := num % 10
        result = (result * 10) + reminder
        num /= 10
    }
    
    if result > 2147483647 || result < - 2147483648 {
        
        return 0
    }
    
    if x < 0 {
        
        return -result
    } else {
         return result
    }
    
   
    
}