func movesToMakeZigzag(nums []int) int {
    a,b,c,d:=0,0,0,0
    l := len(nums)
    
    for i:=0;i<l;i++ {
        c = 0
        if i>0 && i+1 < l{
            if nums[i-1] > nums[i+1]{
                d = nums[i+1]
            }else{
                d = nums[i-1]
            }
           
        }else if i== 0 {
            d = nums[i+1]
        }else{
            d = nums[i-1]
        }
        
         
        
        if nums[i] >= d {
            c = nums[i]-d+1
        }
        
        
        if i % 2 == 0 {
            a += c
        }else{
            b += c
        }
    }
    
    if a > b {
        return b
    } else {
        return a
    }
}