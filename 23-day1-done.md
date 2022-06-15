This part 1 seems to be taking a while.

We already have the `SlidingWindow` function so it's time to join the pieces together:

`day1.go`

```
func Part2(content string) (uint16, error) {
	return 0, nil
}
```

Make it fail first. I don't always do this, but it's useful to remember and the feeling of satisfaction going from failing test to passing test gives closure.

`day1_test.go`

Rename the first test, and add another based on the example given.

```
func TestPart1Integration(t *testing.T) {
	actual, _ := Part1("199\n200\n208\n210\n200\n207\n240\n269\n260\n263")
	expected := uint16(7)
	if expected != actual {
		t.Errorf("Part1 was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestPart2Integration(t *testing.T) {
	actual, _ := Part2("199\n200\n208\n210\n200\n207\n240\n269\n260\n263")
	expected := uint16(5)
	if expected != actual {
		t.Errorf("Part2 was incorrect, got: %d, want: %d.", actual, expected)
	}
}
```

```
$ go test -test.v
=== RUN   TestCountIncreases
--- PASS: TestCountIncreases (0.00s)
=== RUN   TestPart1Integration
--- PASS: TestPart1Integration (0.00s)
=== RUN   TestPart2Integration
    day1_test.go:29: Part2 was incorrect, got: 0, want: 5.
--- FAIL: TestPart2Integration (0.00s)
=== RUN   TestSlidingWindow
--- PASS: TestSlidingWindow (0.00s)
=== RUN   TestParseInput
--- PASS: TestParseInput (0.00s)
FAIL
exit status 1
FAIL    aoc2021/day1    0.009s
$ 
```

Fix the implementation

```
func Part2(content string) (uint16, error) {
	values, err := Parse(content)
	if err != nil {
		return 0, err
	}
	return CountIncreases(SlidingWindow(values)), nil
}
```

Update main.go

```
func solveDay1() {
	content := loadContent("input/day1.txt")
	part1Result, err := day1.Part1(content)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	part2Result, err := day1.Part2(content)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Day 1 Part 1 %d\n", part1Result)
	fmt.Printf("Day 1 Part 2 %d\n", part2Result)
}
```

For me, I see:

```
$ go run main.go
Day 1 Part 1 1342
Day 1 Part 2 1378
$
```

The repeated error checking feels like it could do with some improvement.

Here are a couple of articles I found to mull over:

* [Error Handling in Go that Every Beginner should Know](https://hussachai.medium.com/error-handling-in-go-a-quick-opinionated-guide-9199dd7c7f76)
* [Avoid checking if error is nil repetition?](https://stackoverflow.com/questions/18771569/avoid-checking-if-error-is-nil-repetition) (8 years old)

