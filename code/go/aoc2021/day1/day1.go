package day1

func CountIncreases(depths []uint16) uint16 {
	var increases uint16 = 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			increases++
		}
	}
	return increases
}

func SlidingWindow(depths []uint16) []uint16 {
	ret := make([]uint16, len(depths)-2)
	for i := 2; i < len(depths); i++ {
		ret[i-2] = depths[i] + depths[i-1] + depths[i-2]
	}
	return ret
}

func Part2(content string) (uint16, error) {
	return 0, nil
}

func Part1(content string) (uint16, error) {
	values, err := Parse(content)
	if err != nil {
		return 0, err
	}
	return CountIncreases(values), nil
}
