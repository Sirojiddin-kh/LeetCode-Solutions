func moveZeroes(nums []int)  {
    i := 0
    j := 0
    for i < len(nums){
        if nums[i] == 0{
            i += 1
        } else{
            temp := nums[j]
            nums[j] = nums[i]
            nums[i] = temp
            i += 1
            j += 1
        }
    }
    
}