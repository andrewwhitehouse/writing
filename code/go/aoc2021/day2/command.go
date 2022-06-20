package day2

type Direction uint16

const (
	Forward Direction = 0
	Up      Direction = 1
	Down              = 2
)

type Command struct {
	direction Direction
	distance  uint16
}

func newCommand(direction Direction, distance uint16) *Command {
	c := Command{direction: direction, distance: distance}
	return &c
}
