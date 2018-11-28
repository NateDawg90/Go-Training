package pascal

const testVersion = 1

func Triangle(n int) [][]int {
	res := [][]int{}

	row := []int{}
	last := []int{}
	for i := 0; i < n; i++ {
		if len(res) == 0 {
			row = []int{1}
		} else if n == 1 {
			row = []int{1, 1}
		} else {
			row = []int{1}
			last = res[len(res)-1]
			var addition int
			for j := 0; j < len(last)-1; j++ {
				addition = last[j] + last[j+1]
				row = append(row, addition)
			}
			row = append(row, 1)
		}
		res = append(res, row)
	}
	return res
}
