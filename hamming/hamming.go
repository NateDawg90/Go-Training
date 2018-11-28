package hamming

import (
	"fmt"
)

const testVersion = 5

func Distance(a, b string) (int, error) {
	// your code here!
	if len(a) != len(b) {
		return -1, fmt.Errorf("strands different lengths")
	}
	hamming := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hamming += 1
		}
	}
	return hamming, nil
}
