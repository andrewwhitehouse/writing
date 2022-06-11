**Project Oroganisation**

When I started solving the Advent of Code problems in Rust I organised the directory like this:

```
.
├── Cargo.toml
├── input
│   ├── day1.txt
│   └── day2.txt
└── src
    ├── day1.rs
    ├── day2.rs
    ├── day3.rs
    └── main.rs
```

I see there are being two sets of executable deliverables here: the daily implementation files (day1.rs, day2.rs, ...) each with its own tests, and the `main.rs` which brings all the implementations together and runs them on the actual problem data set.

It helps that the Rust tests and implementation are combined into the single file, as it simplifies code organisation.

[_I was reviewing the previous Rust code I wrote and I think it's harder to follow when showing only the code fragments; I need to remember to show the directory structure, and also (probably) to show the full set of code at the end._]

Go's organisation is different.

I battled with this for a while. Go enforces a directory structure that matches your package naming. 

To illustrate this let's setup a project, following [How to Write Go Code](https://go.dev/doc/code).

Go now uses the concept of modules for code organisation.

To start with you need your Go environment set up; there are three key variables that apply:

**GOROOT** is set to where Go is installed on your machine (mine is `/usr/local/go/`).

**GOPATH** is the locally of your locally built Go files. Actually, you don't have to set this; [since Go 1.11 this is defaulted](https://stackoverflow.com/questions/21001387/how-do-i-set-the-gopath-environment-variable-on-ubuntu-what-file-must-i-edit/53026674#53026674) to your home directory.

**Project Setup**

`$ mkdir aoc2021`

`$ cd aoc2021`

`$ go mod init aoc2021`

```
$ cat go.mod
module aoc2021

go 1.18
$
```

Some of the examples use project locations that reference github, like `github.com/andrewwhitehouse/aoc2021`; this seems excessive to me. I understand that it's helpful to have a uniform way of accessing all of your stuff, but I see this as a self-contained project which -- right now -- isn't likely to be imported into other projects. So I'm going to keep the names shorter and skip the `github.com/username` part.

Create a main "hello world" file:

`main.go`

```
package main

import "fmt"

func main() {
    fmt.Println("Hello, world.")
}
```

```
$ go run main.go
Hello, world.
$
```

For each daily challenge, I'm going to create a test file and an implementation file in Go.

To start with I tried putting these at the top-level of the project.

`day1.go`

```
package day1

func CountIncreases(depths []int16) uint16 {
    return 0 // Not yet implemented
}
```

`day1_test.go`

```
```

When I run my tests I see:

```
$ go test
found packages day1 (day1.go) and main (main.go) in . . ./code/go/aoc2021
$
```

Then I tried putting them in the `main` package.

That exercises the tests correctly

```
$ go test
--- FAIL: TestCountIncreases (0.00s)
    day1_test.go:10: CountIncreases was incorrect, got: 0, want: 7.
FAIL
exit status 1
FAIL	aoc2021	0.005s
$ 
```

If we update main.go to call the function:

```
package main
  
import "fmt"

func main() {
    fmt.Println(CountIncreases([]int16{1,2,3}))
}
```

This works (adding the directory name):

```
$ go run .
0
$
```

but this doesn't

```
$ go run main.go
# command-line-arguments
./main.go:6:17: undefined: CountIncreases
$
```

With all of the functions in the `main` package I can see that it could become hard to follow which functions are for which day (unless I include the day name in the function name, e.g. `Day1CountIncreases`. 

But I think using packages is better.

To fix the earlier issue we put the different packages in their own folder.

```
.
├── day1
│   ├── day1.go
│   └── day1_test.go
├── go.mod
└── main.go
```

`day1.go`

```
package day1
  
func CountIncreases(depths []int16) uint16 {
    return 0 // Not yet implemented
}
```

`day1_test.go`

```
package day1
  
import "testing"

func TestCountIncreases(t *testing.T) {
        depths := []int16{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
        var expected uint16 = 7
        actual := CountIncreases(depths)
        if expected != actual {
                t.Errorf("CountIncreases was incorrect, got: %d, want: %d.", actual, expected)
        }
}
```

`main.go`

```
package main
  
import (
        "fmt"
        "aoc2021/day1"
)

func main() {
    fmt.Println(day1.CountIncreases([]int16{1,2,3}))
}
```

To run the tests, you need to be in the sub-folder.

```
$ go test
?   	aoc2021	[no test files]
You have mail in /var/mail/andrewwhitehouse
Andrews-MacBook-Pro:tmp andrewwhitehouse$
```

```
$ cd day1
$ go test
--- FAIL: TestCountIncreases (0.00s)
    day1_test.go:10: CountIncreases was incorrect, got: 0, want: 7.
FAIL
exit status 1
FAIL	aoc2021/day1	0.005s
$ 
```

Here are the files now:

`main.go`

```
package main
  
import (
    "aoc2021/day1"
    "fmt"
)

func main() {
    fmt.Println(day1.CountIncreases([]int16{1, 2, 3}))
}
```

`day1/day1.go`

```
package day1
  
func CountIncreases(depths []int16) uint16 {
    return 0 // Not yet implemented
}
```

`day1/day1_test.go`

```
package day1
  
import "testing"

func TestCountIncreases(t *testing.T) {
    depths := []int16{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
    var expected uint16 = 7 
    actual := CountIncreases(depths)
    if expected != actual {
        t.Errorf("CountIncreases was incorrect, got: %d, want: %d.", actual, expected)
    }   
}
```

I'm running `go fmt` regularly on the files to make sure that the formatting is consitent. It uses tabs for indentation, and then I set my `tabstop` to 4 in the editor.

[_With this arrangement I want a way to run all the tests, and am still figuring ouw how to do this._]

```
$ go run .
0
$ go run main.go
0
$
```

Fix the implementation:

```
package day1
  
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

In the day1 folder ...

```
$ go test
PASS
ok  	aoc2021/day1	0.005s
$
```

In the top-level project folder ...

```
$ go run main.go
2
$
```
