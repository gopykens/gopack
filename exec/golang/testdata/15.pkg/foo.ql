package foo

func ReverseMap(m map[string]int) map[int]string {
    return {v: k for k, v <- m}
}
