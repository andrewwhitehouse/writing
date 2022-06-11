package day1

func CountIncreases(depths []int16) uint16 {
	var increases uint16 = 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			increases++
		}
	}
	return increases
}
