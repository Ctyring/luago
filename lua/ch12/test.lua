---
--- Generated by Luanalysis
--- Created by Lenovo.
--- DateTime: 2022/12/11 23:11
---
t = {a = 1, b = 2, c = 3}
for k, v in pairs(t) do
    print(k, v)
end

t = {"a", "b", "c"}
for k, v in ipairs(t) do
    print(k, v)
end