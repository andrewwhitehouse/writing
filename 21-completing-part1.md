**Parsing a file, and Completing Part 1**

The Advent of Code puzzles each follow a particular pattern:

- there is some domain logic to solve the puzzle based on a set of rules that have been provided for us
- input data is provided in a file 
- there are two parts to each day's puzzle; the second part isn't revealed until you have solved the first part and so it may involve some refactoring (that's one of the reasons why having tests is a good idea)

Currently we have a unit test which exercises our CountIncreases function.

We need to add some parsing logic to take the input and convert it to a form that is useful for our domain logic.

Add the following files:

`day1/parse.go`

```
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
```

`day1/parse_test.go`

```
package day1

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	actual, err := Parse("11\n13\n17\n19\n301")
	expected := []uint16{11, 13, 17, 19, 301}
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Parse was incorrect, got: %v, want: %v.", actual, expected)
	}
}
```

Run the tests with `go test`.

The parse function takes as its parameter the input, and returns an array of uint16 integers, as well as an error.

Create a function in `day1.go` that brings the pieces together:

`day1.go`

```
func Part1(content string) (uint16, error) {
	values, err := Parse(content)
	if err != nil {
		return 0, err
	}
	return CountIncreases(values), nil
}
```

`day1_test.go`

```
func TestIntegration(t *testing.T) {
	actual, _ := Part1("199\n200\n208\n210\n200\n207\n240\n269\n260\n263")
	expected := uint16(7)
	if expected != actual {
		t.Errorf("Part1 was incorrect, got: %d, want: %d.", actual, expected)
	}
}
```

Starting a function name with a capital letter exports it so that it can be imported by other packages.

This test checks that the two components -- parsing and domain logic -- work together correctly.

These tests actually overlap. To see this, temporarily change the result from `CountIncreases` to return an invalid value (for example: 0). How many tests fail?

Ideally we would want a test failure to help us identify where the issue is, but right now it could be in two places. (In this case, I would think about the dependencies -- what is calling what -- and start debugging the code that is common to both.)

There are other ways of writing tests that allow us to be more targeted, and we'll look at those in future chapters.

To finish part 1, we want to load the input data that has been provided for us.

The file contains 2000 lines, and starts like this:

```
176
184
196
199
204
206
```

Create a directory called `input` under the top level of your project directory, and save `day1.txt` into it.

Then add a helper function in `main.go` to load the file:

```
func loadContent(fileName string) string {
	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(fileBytes)
}
```

I also added a function to contain the logic for solving day 1

```
func solveDay1() {
	content := loadContent("input/day1.txt")
	part1Result, err := day1.Part1(content)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Day 1 Part 1 %d\n", part1Result)
}
```

It's possible that the day 1 processing could fail, if we accidentally corrupted the input data, so we should check the return value (and even add a test for it if this code is going to be long-lived or shared with others).

The imports need to be updated too:

```
import (
	"aoc2021/day1"
	"fmt"
	"io/ioutil"
	"os"
)
```

Finally update your main function in `main.go` to call `solveDay1()`.

In the top level directory of your project run the code:

```
$ go run main.go
Day 1 Part 1 <result>
$ 
```

Observations:

- Go returns success and error results explicitly from functions; you should generally check for errors especially if your input data is going to vary
- Exported functions begin with a capital letter
- Running `go test -test.v` prints the test names as they are run

@beaver