package accumulate

const testVersion = 1

func Accumulate(s []string, f func(string) string) []string {

	var new_arr []string = make([]string, len(s))

	for i := 0; i < len(s); i++ {
		// new_arr = append(new_arr, f(s[i]))\
		newArr[i] = f(s[i])
	}
	return newArr
}
