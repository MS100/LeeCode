//https://leetcode-cn.com/problems/trapping-rain-water/
func trap(height []int) int {
    sum := 0
    m,n := 0,len(height)-1
    h := 0
    for m < n {
        if height[m] <= height[n]{
            if h < height[m]{
                h = height[m]
            }else{
                sum += h - height[m]
            }
            
            m ++
        }else{
            
            if h < height[n]{
                h = height[n]
            }else{
                sum += h - height[n]
            }

            n--
        }
    }
    return sum
}
