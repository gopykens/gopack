x := 0
if t := false; t {
    x = 3
} else {
    x = 5
}
println("x:", x)

x = 0
switch s := "Hello"; s {
default:
    x = 7
case "world", "hi":
    x = 5
case "xsw":
    x = 3
}
println("x:", x)

v := "Hello"
switch {
case v == "xsw":
    x = 3
case v == "Hello", v == "world":
    x = 9
default:
    x = 7
}
println("x:", x)

v = "Hello"
switch {
case v == "xsw":
    x = 3
case v == "hi", v == "world":
    x = 9
default:
    x = 11
}
println("x:", x)
