//不缺失的状态下求和，然后减去实际数组和
func missingNumber(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    n := int((1+len(nums)) * len(nums) / 2)
    
    sum := 0
    for i:=0;i<len(nums);i++{
        sum += nums[i]
    }
    
    return n - sum
}
