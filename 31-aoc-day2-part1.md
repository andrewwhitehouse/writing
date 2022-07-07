**Advent of Code in Go, Day 1 Part 1**

[_Well, our Prime Minister has very reluctantly resigned but he's talking about staying around until October so he and his partner can use the official residence for their summer hols._]

This post has been a long time coming.

This is continued from [15 days ago](https://wc3.akimbo.com/t/andrew-whitehouse-dailies/61070/495?u=andrewwhitehouse).

There was something that was bothering me about the previous post; there are a number of points in the code where we encounter a possible error condition. 

The function signature looks like this:

```
func Parse(s string) ([]Command, error) {
```

which returns a tuple consisting of a Command array, and an error. If the function succeeds, the error is _nil_ which means "no value". If the function encounters an error, we want to return a non-value" for the command array. (The same issue applies for the individual commands, too).

Golang parameters are passed by value, which means the parameter value is copied. Changing the passed value inside the function doesn't change the value in the caller.

If we pass a pointer, we can use the nil value to represent "no value", like this `func Parse(s string) ([]*Command, error) {`

But this seems like an unnecessary use of points.

On further [reading](https://medium.com/@annapeterson89/whats-the-point-of-golang-pointers-everything-you-need-to-know-ac5e40581d4d) I concluded that using pointers isn't appropriate, so I needed to define another non-value.

I did it like this:

`command.go`

```
const (
	None    Direction = 0
	Forward Direction = 1
	Up      Direction = 2
	Down    Direction = 3
)
```

`parse.go`

```
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
```

I updated the tests too.

Having done this, I then wrote the Day2 (part 1) logic:

`day2.go`

```
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
```

The original uint16 type may not be too small to handle an arbitrarily long list of positions, and then multiply them at the end. uint32 _may_ be big enough but I opted to 64 bits instead to be safe. Frankly I'm guessing on how big the values should be without looking at the test data and considering what the eventual value should be. 32 bits probably would have been enough, actually.

Here's the function in `main.go` that brings it together for part 1:

```
func solveDay2() {
	content := loadContent("input/day2.txt")
	part1Result, _ := day2.Part1(content)
	fmt.Printf("Day 2 Part 1 %d\n", part1Result)
}
```

And the result ...

```
$ go run main.go
Day 1 Part 1 1342
Day 1 Part 2 1378
Day 2 Part 1 1636725
$ 
```

@beaver
