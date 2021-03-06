package pub

type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}
func (s IntSlice) Less(i, j int) bool {
	return s[i] > s[j]
}
func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func insertSort(a IntSlice) IntSlice {
	var i, j int
	var n = len(a)
	if n == 0 || n == 1 {
		return a
	}
	for i = 1; i < n; i++ {
		for j = i - 1; j >= 0 && a.Less(j, j+1); j-- {
			a.Swap(j, j+1)
		}
	}
	return a
}
