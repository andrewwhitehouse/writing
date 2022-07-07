*Advent of Code Day 2**

Navigating the depths.

Day 2 involves processing a series of navigation for our fictitious submarine in order to determine its eventual location.

For example, after the submarine follows these instructions

```
forward 5
down 5
forward 8
up 3
down 8
forward 2
```

It has a horizontal position of 5+8+2=15, and a depth of 5-3+8=10.

The example raises a couple of initial questions:
- can we go backwards? (We'll assume now.)
- if the "up" command is given a value that is greater than our current depth, what should we do? (Again, we'll assume it can't happen until we see evidence to the contrary.)
- could we encounter any other commands in the full data set other than forward, down or up? (We'll code for the three we're given and adapt when we see the the real example.)

Let's lay the groundwork.

What would be good terminology to use for our solution?

Here's a clue:

>It seems like the submarine can take a series of commands like forward 1, down 2, or up 3:

So each of those lines in the example represents a "command", and we want a term to represent the associated value; "distance" seems good enough since that can represent horizontal or vertical movement.

`day2/command.go`

```
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
```

Write a test that checks the `newCommand` function.

`day2/command_test.go`

```
package day2

import (
	"testing"
)

func TestNewCommand(t *testing.T) {
	command := newCommand(Up, 20)
	actualDirection := command.direction
	expectedDirection := Up
	if actualDirection != expectedDirection {
		t.Errorf("Command direction was incorrect, got: %d, want: %d.", actualDirection, expectedDirection)
	}
	expectedDistance := uint16(20)
	actualDistance := command.distance
	if actualDistance != expectedDistance {
		t.Errorf("Command distance was incorrect, got: %d, want: %d.", actualDirection, expectedDirection)
	}
}
```

We also need to parse the input data into a form that is usable by our domain logic. We're choosing an array of Command objects (actually, pointers to them).

`day2/parse.go`

```
package day2

func Parse(s string) []Command {
	return make([]Command, 0)
}
```

`day2/parse_test.go`

```
package day2

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := "forward 5\ndown 5\nforward 8\nup 3\ndown 8"
	actual := Parse(input)
	expected := []*Command{
		newCommand(Forward, 5),
		newCommand(Down, 5),
		newCommand(Forward, 8),
		newCommand(Up, 3),
		newCommand(Down, 8),
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Parse was incorrect, got: %v, want: %v.", actual, expected)
	}
}
```

[_To do: Go by default is pass by value, which is why we pass around pointers -- I believe -- to prevent copying objets. I need to dig into this more._]

```
$ go test -test.v
=== RUN   TestNewCommand
--- PASS: TestNewCommand (0.00s)
=== RUN   TestParseInput
    parse_test.go:19: Parse was incorrect, got: [], want: [0xc0001220c8 0xc0001220cc 0xc0001220d0 0xc0001220d4 0xc0001220d8].
--- FAIL: TestParseInput (0.00s)
FAIL
exit status 1
FAIL    aoc2021/day2    0.007s
$ 
```

Now we need to make the parsing work.

