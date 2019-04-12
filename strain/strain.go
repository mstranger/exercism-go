package strain

type Ints []int
type Lists [][]int
type Strings []string

func (arr Ints) Keep(f func(int) bool) Ints {
	if arr == nil {
		return nil
	}
	filtered := make([]int, 0)
	for _, v := range arr {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func (arr Lists) Keep(f func([]int) bool) Lists {
	if arr == nil {
		return nil
	}
	filtered := make([][]int, 0)
	for _, v := range arr {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func (arr Strings) Keep(f func(string) bool) Strings {
	if arr == nil {
		return nil
	}
	filtered := make([]string, 0)
	for _, v := range arr {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func (arr Ints) Discard(f func(int) bool) Ints {
	if arr == nil {
		return nil
	}
	filtered := make([]int, 0)
	for _, v := range arr {
		if !f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
