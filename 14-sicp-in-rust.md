[_Background: I'm exploring where my priorities are for this book. It should be able learning to code by solving problems. Getting code working early and then iterating on the solution. And writing tests to help the reader think about the domain, what are good test cases, and break down the problem/solution into small steps so that it doesn't become overwhemling. There is also an opportunity to discover trade offs with different data structures and algorithsm. Also I think that the combination of a "classic" text which avoids focusing on syntax through the use of LISP, with a contemporary language like Rust whose syntax can be quite intimidating but can be tackled bit by bit, could be interesting. I'm not abandoning the Python or Clojure approaches ... simply focusing on this for now. I'm planning on being around for a while!_]

Learning to code by solving problems is an enjoyable way to explore different aspects of a language.

Here is the problem statement from the [Structure and Interpretation of Computer Programmes](https://mitpress.mit.edu/sites/default/files/sicp/full-text/book/book.html):

> Exercise 1.3.  Define a procedure that takes three numbers as arguments and returns the sum of the squares of the two larger numbers.

A few initial thoughts come to mind.

Firstly, there are two steps to this, which might be expressed in functions like this:

`sum_of_squares(two_largest(a,b,c))`

Let's tackle the two_largest function first.

In Rust, we have a choice of data structure to return; we could use an array, a tuple, or a set.

What do these look like?

[The principle in this book is to get something working quickly and then iterate on it ... and using tests as guard rails along the way.]

`$ cargo new --lib ex1_3`
`$ cd ex1_3`

`lib.rs`

```
fn two_largest(_a: i32, _b: i32, _c: i32) -> [i32; 2] {
    [0, 0]
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn should_return_first_two_if_in_descending_order() {
        assert_eq!([3, 2], two_largest(3, 2, 1));
    }
}
```

We start with a failing test. [_Add explanation of Test-Driven Development cycle, if this is the first time._]

The Rust compiler complained about our not using the arguments to the function, so we prefix them with an underscore to indicate that we're ignoring them intentionally. The function uses 32-bit integers _i32_; the problem statement doesn't say how large the numbers can be, or whether we would allow float values. Let's assume we're OK with integers for now.

`$ cargo test`

```
$ cargo test --lib
    Finished test [unoptimized + debuginfo] target(s) in 0.00s
     Running unittests (/Users/andrewwhitehouse/code/writing/code/rust/ex1_3/target/debug/deps/ex1_3-dabf7430ba925b35)

running 1 test
test tests::should_return_first_two_if_in_descending_order ... FAILED

failures:

---- tests::should_return_first_two_if_in_descending_order stdout ----
thread 'tests::should_return_first_two_if_in_descending_order' panicked at 'assertion failed: `(left == right)`
  left: `[3, 2]`,
 right: `[0, 0]`', src/lib.rs:11:9
note: run with `RUST_BACKTRACE=1` environment variable to display a backtrace


failures:
    tests::should_return_first_two_if_in_descending_order

test result: FAILED. 0 passed; 1 failed; 0 ignored; 0 measured; 0 filtered out; finished in 0.00s

error: test failed, to rerun pass '--lib'
$
```

Starting with a test that fails deliberately helps us to see that changes we have made have fixed the code.

Let's commit our changes [_assume we've already set up git_]

`$ git add .`

`$ git commit -m "Failing test"`

Now we can fix the test in the simplest way possible, by returning the expected values:

```
fn two_largest(_a: i32, _b: i32, _c: i32) -> [i32; 2] {
    [3, 2]
}
```

```
$ cargo test --lib
   Compiling ex1_3 v0.1.0 (/Users/andrewwhitehouse/code/writing/code/rust/ex1_3)
    Finished test [unoptimized + debuginfo] target(s) in 1.53s
     Running unittests (target/debug/deps/ex1_3-dabf7430ba925b35)

running 1 test
test tests::should_return_first_two_if_in_descending_order ... ok

test result: ok. 1 passed; 0 failed; 0 ignored; 0 measured; 0 filtered out; finished in 0.00s

$
```

Commit the change (`git commit -am "First test passing.`)

Let's add another test case to make our function slightly more flexible:

```
#[test]
fn should_return_first_two_if_in_ascending_order() {
    assert_eq!([3, 2], two_largest(2, 3, 1));
}
```    

Write a failing test, and check that it fails (`cargo test --lib`).

Now add an if condition:

```
fn two_largest(a: i32, b: i32, _c: i32) -> [i32; 2] {
    if a > b { [a, b] } else { [b, a] }
}
```

Rust has optional parentheses around the condition `a > b`, but requires them around the result.

[_to be continued._]

> Other possible options to explore:
- sorting
- have the two_largest function return a set
- don't separate the functions, and observe opportunity for simplificaiton
- how does the solution change if we introduce floats
- note how types (e.g. i32, u32, u128, f64) affect the test cases
- using match in the solution
- using tuples in the solution

> Note to come back to ...

> We could use a `match` in Rust, which some other languages refer to as a switch statement. There are only three numbers so we could enumerate all the possible [permutations](https://en.wikipedia.org/wiki/Permutation) (of which there are 6, or 3 _factorial_, for 3 numbers). 

> Exercises:

> 1. If there were 5 numbers, what would the possible number of permutations be? Which of the solutions would become less appealing?

@beaver