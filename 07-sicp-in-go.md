[_I've been mulling over what direction to take this in. Comparing two languages gives a useful concrete way to compare and contrast the design. So what would that look like? For blockchain it's Solidity vs Rust. For building an MVP application it's probably Clojure vs Kotlin. For moving from legacy to a newer stack it's Java vs Clojure or Kotlin. For centralised vs decentralied it could be serverless Rust vs blockchain Rust. And I looked at what I was planning to do ... start with Python and introduce elements of Rust (which I'm learning) and Clojure (which I know), or possibly Java (which I also know) and I thought _this is too much_. For me, and the person who's learning it. So let's focus on the audience, which is my daughter right now. And the skills that she'll need for a data engineering job. I don't know the specifics of the data engineering part yet, so I'll focus on the language.

I have some FOMO about leaving the other topics, and whether this is a market of one. But I'm OK with that. And it's doubling down on some of the non-technical issues: what are some of the things you would have likes to know when starting your career; what are the ways that men can help to make technology a good place to work; where are the boundaries of what makes a reasonable code review versus something else. I think I also need to be careful not to spoon feed. Because there is satisfaction in learning the meta tools, like figuring out an answer when you don't have someone standing behind you telling you what to do._]

We have explored writing a function in Python to calculate the square of a number. 

The other language we're going to use in this book, which is statically typed, meaning that we specify in our code what types we are using and -- depending on whether we compile the code first -- either the runtime, or the compiler, checks that our usage of those types makes sense.

Those types are like extra scaffolding, which invovle more up-front work, but they help you to see how the pieces fit together, like lego.

For a small project [[reference](https://www.reddit.com/r/golang/comments/qwyk3k/what_is_the_best_way_to_organize_code_within_a_go/)] we can create a `main` file in a sub-directory.

`mkdir -p go/chapter`
`cd go/chapter`

In code/go/chapter1

Create main.go with a "hello world" to get started.

```
package main
  
import "fmt"

func main() {
    fmt.Println("hello world")
}
```

We have various integer and decimal [numeric types](https://tip.golang.org/ref/spec) available.

We can interrogae the [math](https://pkg.go.dev/math) package to see what the maximum possible integer value is.

```
import "fmt"
import "math"

func main() {
    fmt.Println(math.MaxInt64);
}
```

(Note that it's "math" singular ... spelling of elements shipped with the language tend to follow the American form, like `color`. In the programmes we write, we can use whatever dialect works for those maintaining the code.)

```
$ go run main.go 
9223372036854775807
$
```

OK that's for a signed integer. 

```
fmt.Println(math.MaxUint64);
```    
    
gives

```
$ go run main.go 
# command-line-arguments
./main.go:7:16: constant 18446744073709551615 overflows int
$ 
```

Googling: <https://stackoverflow.com/questions/16474594/how-can-i-print-out-an-constant-uint64-in-go-using-fmt> it turns out the type used depends on the context. We need to explicitly convert the `MaxUint64` constant in our code.

```
fmt.Println(uint64(math.MaxUint64));
```
    
Multiping the maximum signed integer by itself is greater than the maximum (64-bit) integer. Is that an issue for us? Is it important to be able to work with large numbers? In this example we can make up our own requirements. For now, let's restrict the types of the square function so that we don't encounter these issues.

```
func square(x int32) int64 {
    return x*x
}

func main() {
    fmt.Println(square(math.MaxInt32));
}
```

```
$ go run main.go 
# command-line-arguments
./main.go:7:13: cannot use x * x (type int32) as type int64 in return argument
$
```

Still not happy.

```
func square(x int32) int64 {
    return int64(x) * int64(x)
}
```

```
$ go run main.go
4611686014132420609
$
```

Does that look right? Let's do a sanity check ... 2x10^9  times itself is 4 x 10^18.

So far we've limited ourselves to numbers that are integers. But we want to handle decimals too. Let's change the implementation.

```
func square(x float64) float64 {
    return x*x
}

func main() {
    fmt.Printf("Max float64 %f\n", math.MaxFloat64);
}
```
```
$ go run main.go
Max float64 179769313486231570814527423731704356798070567525844996598917476803157260780028538760589558632766878171540458953514382464234321326889464182768467546703537516986049910576551282076245490090389328944075868508455133942304583236903222948165808559332123348274797826204144723168738177180919299881250404026184124858368.000000
$

```


You can see that our statically typed language is requireing us to be more intentional about the types we choose to use in our code. 

Not all code involves mathematical calculations. But there is plenty of counting. 

We need someone to tell us when constraints we are working in (if they know) and otherwise make sensible assumptions. In a project team that person is often a Business Analysit or Product Owner or Product Manager

Let's suppose we were creating a business that sold triangular gardens with a right angle in Central London (based on a quirky medieval layout -- no tidy grids here). We want to calculate the length of the longest side, in millimetres.

Realistically, a 32-bit integer is going to suffice because that's around 4 billion millimeters, or 4 million metres. Quite a large garden.

```
package main
  
import "fmt"
import "math"

func longestWallLength(x, y uint32) uint32 {
  return uint32(math.Sqrt(float64(x*x + y*y)))
}

func main() {
    fmt.Printf("Longest wall %dmm\n", longestWallLength(10 * 1000, 4950));
}
```

```
$ go run main.go
Longest wall 11158mm
$
```

(We can debate whether it makes sense to measure inner city triangular gardens in millimeters.) 

You can see that the static typing in Go required us to be quite deliberate about the type conversions in the Pythagorean triangle calculation. Static typing doesn't prevent all bugs, but it helps to make assumptions explicit.

@beaver


