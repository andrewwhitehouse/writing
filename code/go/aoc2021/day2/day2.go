package day2

type Position struct {
	distance uint16
	depth    uint16
}

func Navigate(commands []Command) Position {
	position := Position{0, 0}
	for _, command := range commands {
		switch command.direction {
		case Forward:
			position.distance += command.distance
		case Down:
			position.depth += command.distance
		case Up:
			position.depth -= command.distance
		}
	}
	return position
}

func Part1(content string) (uint64, error) {
	commands, err := Parse(content)
	if err != nil {
		return 0, err
	}
	position := Navigate(commands)
	return uint64(position.depth) * uint64(position.distance), nil
}
