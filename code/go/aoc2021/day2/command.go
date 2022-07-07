package day2

type Direction uint16

const (
	None    Direction = 0
	Forward Direction = 1
	Up      Direction = 2
	Down    Direction = 3
)

type Command struct {
	direction Direction
	distance  uint16
}

func newCommand(direction Direction, distance uint16) Command {
	return Command{direction: direction, distance: distance}
}
