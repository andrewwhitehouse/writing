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

func Part1(content string) (uint16, error) {
	values, err := Parse(content)
	if err != nil {
		return 0, err
	}
	return CountIncreases(values), nil
}
