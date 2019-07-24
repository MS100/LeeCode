func numberToWords(num int) string {
    if num == 0 {
        return "Zero"
    }
    qianfen := [4]string{"","Thousand ", "Million ", "Billion "}
    
    ge := [10]string{"","One ","Two ","Three ","Four ","Five ","Six ","Seven ","Eight ","Nine "}
    shiji := [10]string{"Ten ","Eleven ","Twelve ","Thirteen ","Fourteen ","Fifteen ","Sixteen ","Seventeen ","Eighteen ","Nineteen "}
    shi := [10]string{"","","Twenty ","Thirty ","Forty ","Fifty ","Sixty ","Seventy ","Eighty ","Ninety "}
    
    res := ""
    for i := 0; num > 0; i++ {
        a := num % 1000
        c := ""
        
        shigewei := a % 100
        if shigewei > 9 && shigewei < 20 {
            c = shiji[shigewei - 10] + c
        }else {
            gewei := shigewei % 10
            c = ge[gewei] + c
            shiwei := int(shigewei / 10)
            c = shi[shiwei] + c 
        }
        
        baiwei := int(a / 100)
        if baiwei > 0 {
            c = ge[baiwei] + "Hundred " + c
        }
        
        if c != "" {
            res = c + qianfen[i] + res        
        }
        num = num / 1000
    }
    
    return res[:len(res)-1]
}


