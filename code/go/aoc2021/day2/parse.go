package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func parseLine(s string) (Command, error) {
	parts := strings.Split(strings.TrimSpace(s), " ")
	var direction Direction
	noCommand := newCommand(None, 0)
	switch {
	case parts[0] == "forward":
		direction = Forward
	case parts[0] == "up":
		direction = Up
	case parts[0] == "down":
		direction = Down
	default:
		return noCommand, fmt.Errorf("unrecognised direction %s", s)
	}
	distance, err := strconv.Atoi(parts[1])
	if err != nil {
		return noCommand, err
	}
	result := newCommand(direction, uint16(distance))
	return result, nil
}

func Parse(s string) ([]Command, error) {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	ret := make([]Command, len(lines))
	for i := 0; i < len(ret); i++ {
		command, err := parseLine(lines[i])
		if err != nil {
			return nil, err
		}
		ret[i] = command
	}
	return ret, nil
}
