package day1

import (
	"strconv"
	"strings"
)

func sliceAtoi(sa []string) ([]uint16, error) {
	si := make([]uint16, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, uint16(i))
	}
	return si, nil
}

func Parse(content string) ([]uint16, error) {
	sliceData := strings.Split(strings.TrimSpace(content), "\n")
	return sliceAtoi(sliceData)
}
