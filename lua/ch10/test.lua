---
--- Generated by Luanalysis
--- Created by Lenovo.
--- DateTime: 2022/12/10 21:58
---
function newCounter()
    local i = 0
    return function() -- 匿名函数
        i = i + 1
        return i
    end
end
c1 = newCounter()
print(c1()) --> 1
print(c1()) --> 2
c2 = newCounter()
print(c2()) --> 1
print(c1()) --> 3
print(c2()) --> 2