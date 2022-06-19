package day2

type Direction uint16

const (
	Forward Direction = 0
	Up      Direction = 1
	Down              = 2
)

type command struct {
	direction Direction
	distance  uint16
}

func newCommand(direction Direction, distance uint16) *command {
	c := command{direction: direction, distance: distance}
	return &c
}
