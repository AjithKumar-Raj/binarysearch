package main

func main() {
	s0 := "ajith"
	s1 := "raj"
	var ans []byte
	for i := 0; i < len(s0) || i < len(s1); i++ {
		if i < len(s0) {
			ans = append(ans, s0[i])
		}
		if i < len(s1) {
			ans = append(ans, s1[i])
		}
	}
	return string(ans)
}
