//https://leetcode-cn.com/problems/jump-game-ii/
func jump2(nums []int) int {
    l := len(nums)-1
    if l == 0 {
        return 0
    }
    
    if nums[0] >= l {
        return 1
    }
    
    x,max,n := 2,nums[0],nums[0]
    
    for i:= 1;i<l;i++{
        t := i+nums[i]
        if t > max {
            max = t
            if max >= l {
                return x
            }
        }
        
        if i == n {
            x++
            n = max
        }
    }
        
    
    return x
}

func jump(nums []int) int {
    l := len(nums)-1
    if l == 0 {
        return 0
    }
    
    if nums[0] >= l {
        return 1
    }
    
    x,max,m,n := 2,nums[0],1,nums[0]
    
    for {
        for i:= m;i<=n;i++{
            t := i+nums[i]
            if t > max {
                max = t
                if max >= l {
                    return x
                }
            }
        }
        
        m,n = n+1,max
        x++
    }
    
    return x
}
