n := 2
a := make([]int, uint64(n))
a = append(a, 1, 2, 3)
println(a)

b := make([]int, 0, uint16(4))
c := [1, 2, 3]
b = append(b, c...)
println(b)
