We've introduced functions in Python and Go. Functions are a form of abstraction which allows us to group a set of operations together in a logical unit of code. Writing small well-named functions helps to improve the readability of our code.

In Go we wrote:

```
func square(x float64) float64 {
    return x*x
}
```

and in Python:

```
def square(x):
    return x*x
```

Let's compare and contrast the language features so far.

Both functions have a name, which is to be associated with the function definition in the environment.

Languages tend to standardise on a set of coding guidelines so that stylistic issues don't become a point of contention in a team. Python has its own [style guide](https://peps.python.org/pep-0008/#function-and-variable-names) while Go has it's own `go fmt <filename>` command to fix formatting issues, and there are various style guides that a team can adopt for code reviews ([example](https://github.com/golang/go/wiki/CodeReviewComments#gofmt)).

The function `parameters` are the names used within the body of the function to refer to the corresponding arguments of the function. The parameters in both Python and Go are grouped within parentheses and separated by commas.

In the simplest form, the `body` of the function declaration has a single `return` statement which concists of the keyword `return` followed by the `return` expression that will yield the value of the function application.

In Go, the function body is enclosed by braces `{` and `}`. 

Python ends the first line of a function definition with a colon `:` and subsequent lines are indented. A blank line separates the function body from the rest of the programme.

Go types are specific after the function parameters, and the function return type comes after the parameters. Specifying the parameter type _after_ the parameter names allows parameters of the same type to be grouped.

Both definitions are OK:

```
func multiply1(x, y float64) float64 {
        return x * y
}

func multiple2(x float64, y float64) float64 {
        return x * y
}
```

Having a single return point from a function can make it easier to reason about and debug, particularly if your language doesn't come with a debugger to step through the programme, and you debug your programmes by adding temporary print statements (as we did in the olden days of C programming before Test-Driven Development came along).

Python infers the types in your programme from the content of your programme. It's possible to call a function with an integer _or_ a float type without explicitly saying which it is, and as long as the operation is defined, it will run. Go will enforce the type constraints you define. So for example, running this:

```
import "fmt"

func identity(x int32) int32 {
        return x
}

func main() {
        var age int64 = 22
        fmt.Println(identity(age))
}
```

will produce

```
$ go run main.go
# command-line-arguments
./main.go:11:22: cannot use age (type int64) as type int32 in argument to identity
$ 
```

The `identity` function doesn't seem very useful since it simply returns the supplied parameter. However there are cases where we may want to perform a transformation on a value in some cases and not others, and we do that by passing a function as a function parameter. We'll cover this in [a future chapter].

even though the assigned value will happily fit in a 32-bit integer. (Intuitively, if age is representing years then an unsigned 8-bit integer is a better choice.)

In the next section we're going to cover combining functions ...

@Beaver 
