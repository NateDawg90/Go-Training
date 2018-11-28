package slice

const testVersion = 1

func All(n int, s string) []string {
	var res []string
	if n > len(s) {
		return []string{}
	}
	for i := 0; i < len(s)-n+1; i++ {
		res = append(res, s[i:i+n])
	}
	return res
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return ",", false
	}
	return s[0:n], true

}
