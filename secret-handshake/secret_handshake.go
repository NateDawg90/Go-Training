package secret

const testVersion = 1

func Handshake(n uint) []string {

	code := map[uint]string{
		1: "wink",
		2: "double blink",
		4: "close your eyes",
		8: "jump",
	}

	var result []string

	var i uint
	for i = 0; i < 4; i++ {
		cursor := uint(1 << i)
		if cursor&n == cursor { // ????
			result = append(result, code[cursor])
		}
	}

	reverseFlag := uint(1 << 4)
	if reverseFlag&n == reverseFlag {
		return reverse(result)
	}
	return result
}

func reverse(ss []string) []string {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
	return ss
}
