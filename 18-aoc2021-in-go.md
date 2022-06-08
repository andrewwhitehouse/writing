**Advent of Code in Go**

I'm looking at Go for my next language. It seems to hit the sweet spot of performance and simplicity. Also I see more startups using it as a pragmatic way to get stuff delivered, and the contract market is reasonable.

My current go-to challenges are Advent of Code.

Let's start with Day 1, which is about a submarine, and counting increasing depths.

[If you don't already have go installed, follow the [online instructions](https://go.dev/doc/install). 

The setup is minimal: create a directory.

`$ mkdir aoc2021`

[_Note to self: be clear on typographical conventions so I don't have too much re-work later. e.g. the O'Reilly books tend to get this out of the way early on and I skip over it generally, but it's helpful to have clarity._]

I'm going to launch into the code, and avoid too much pre-amble.

Start with a failing test.

I'm assuming it's OK to organise the code as a package for each day.

The AOC page provides a sample set of input data, and the expected output. 

`day1_test.go`

```
package day1
  
import "testing"

func TestSum(t *testing.T) {
    depths := []int16{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
    expected := 7
    actual := CountIncreases(depths)
    if expected != actual {
       t.Errorf("CountIncreases was incorrect, got: %d, want: %d.", actual, expected)
    }
}
```

`day1.go`

```
package day1
  
func CountIncreases(depths []int16) int {
    return 0
}
```

[_Things that stang out: exported functions are written in _CapsCase_, and the type declaration comes after the variable or function definition, although array brackets come before. Using `go fmt` provides consistency of formatting, and seems to prefer using tabs for indentation. So set your tabstop to something sensible. I'm more accustomed to using spaces in my code, actually, and configure the editor to convert tabs to spaces._]

Related:

https://www.youtube.com/watch?v=SsoOG6ZeyUI

```
$ go test
--- FAIL: TestSum (0.00s)
    day1_test.go:10: CountIncreases was incorrect, got: 0, want: 7.
FAIL
exit status 1
FAIL	_/Users/andrewwhitehouse/code/active-projects/writing/code/go/aoc2021	0.005s
$
```

Implementation:

```
func CountIncreases(depths []int16) int {
    var increases int = 0
    for i := 1; i < len(depths); i++ {
        if depths[i] > depths[i-1] {
            increases++
        }
    }
    return increases
}
```

Run the test again, and you should see that it passes.

Array indices start at 0. We start at index 1 (the second element) and compare each element with the previous one; if it's greater we increment our `increases` counter.

**Exercises:**

1. Go has a form of [short variable declaration](https://go.dev/ref/spec#Short_variable_declarations). Replace the declaration starting with `var` with its short form equivalent.

2. What are range of values you can store in an int16?

This programme may help:

```
package main
  
import (
        "fmt"
        "math"
)

func main() {
        fmt.Println(math.MaxInt16)
        fmt.Println(math.MinInt16)
}
```

Run it with `go run <filename.go>`

3. Can you think of a better return type for `CountIncreases` assuming that the function will never receive an array with fewer than 10,000 elements? Hint: can the return value ever be negative? 

Change the implementation and re-run the test with `go test`.

[details="Answer"]

1. `increases := 0`

2. 32767, -32768

3.

```
func CountIncreases(depths []int16) uint16 {
    var increases uint16 = 0
    for i := 1; i < len(depths); i++ {
        if depths[i] > depths[i-1] {
            increases++
        }
    }
    return increases
}
```

```
func TestSum(t *testing.T) {
    depths := []int16{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
    var expected uint16 = 7
    actual := CountIncreases(depths)
    if expected != actual {
       t.Errorf("CountIncreases was incorrect, got: %d, want: %d.", actual, expected)
    }
}
```

[/details]

**Note**: I had some strange goings on while exploring this problem. In my normal terminal window running the tests works fine. But when I fire it up from a terminal window within *Golang* (Jetbrain's Go IDE) I see an error.

```
$ go test
go: cannot find main module, but found .git/config in /Users/andrewwhitehouse/code/active-projects/writing/code/go/aoc2021
        to create a module there, run:
        go mod init
$
```

It seems to be related to the environment variables:

In Goland:

```
$ which go
/usr/local/opt/go/libexec/bin/go
$ 
```

```
$ env | grep go
PATH=/usr/local/opt/go/libexec/bin:/Users/andrewwhitehouse/go/bin:...
. . .
GOROOT=/usr/local/opt/go/libexec
GOPATH=/Users/andrewwhitehouse/go
$ 
```

Outside of Goland:

```
$ which go
/usr/local/bin/go
$

I'm not yet sure why this is happening, or what the "correct" settings should be.

@beaver