package matrix

type Pair struct {
	i, j int
}

func (m Matrix) Saddle() (pairs []Pair) {
	var maxPairs, minPairs []Pair
	lessFunc := func(i, j int) bool {
		return i <= j
	}
	greaterFunc := func(i, j int) bool {
		return i >= j
	}

	for i, row := range m.m {
		indexes := indexes(row, valueFunc(row, greaterFunc))
		for _, index := range indexes {
			pair := Pair{i, index}
			maxPairs = append(maxPairs, pair)
		}
	}
	for i, col := range m.Cols() {
		indexes := indexes(col, valueFunc(col, lessFunc))
		for _, index := range indexes {
			pair := Pair{index, i}
			minPairs = append(minPairs, pair)
		}
	}
	for _, p := range minPairs {
		for _, pp := range maxPairs {
			if p == pp {
				pairs = append(pairs, p)
			}
		}
	}
	return pairs
}
func indexes(arr []int, value int) (out []int) {
	for i, r := range arr {
		if r == value {
			out = append(out, i)
		}
	}
	return
}
func valueFunc(arr []int, f func(a, b int) bool) int {
	v := arr[0]
	for _, r := range arr {
		if f(r, v) {
			v = r
		}
	}
	return v
}
