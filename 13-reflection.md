**Reflections**

Prompted by @dougmay I have some further reflections on this exercise. 

First of all, how I came up with the scenarios isn't clear.

I'm trying to avoid too much up-front explanation. I find myself reading tech books and getting frustrated. It is the architecture, design and implementation that leads to creating a concrete set of code that delivers value. I understand that authors need to provide some background; I also believe that there is a tendency to pad books with unnecessary detail that wastes the reader's time in the name of adding to the page count.

So I'm figuring out what that balance is.

First of all, let's re-consider the problem statement:

> Define a procedure that takes three numbers as arguments and returns the sum of the squares of the two larger numbers.

So ... three numbers, not an arbitrary number. Also note that there are two main steps in this problem:
1. find two largest numbers
2. return the sum of the square of those numbers

Because the order of values does not matter in addition (they are commutative), it's actually possible to do a simplified solution: remove the smallest number, and sum the other two squared.

By focusing only on the first step, with 3 inputs and a result with a well-defined order (largest first) I've actually introduced more precision than is necessary.

In terms of breaking down the problem

Given a function `two_largest(a,b,c)` we want to return the two largest values like this [largest, second_largest].

To find the largest and second largest we can look at the pairs.

For example:

If a is larger than b, then the largest is either a or c. The second largest depends on the result of the first calculation.

For such a simple problem statement, the number of possible ways of tackling this is surprising.

As I mentioned, there is a solution that involves simply removing the smallest value, and accepting that the order of the two largest items may vary. In that case, a different representation may be needed (a set) I'll explore that in the next post.

I also realised yesterday that Clojure, with its prefix notation, has an elegant way of enumerating the possible solutions:

```
(defn two-largest [a b c]
  (cond
    (>= a b c) [a b]
    (>= b a c) [b a]
    (>= c b a) [a c]
    (>= a c b) [a c]
    (>= b c a) [b c]
    (>= c a b) [c a]))
```

It's easier to check this solution by inspection: are the parameters use in all different permutations, and are we always picking the first two values.

As I was saving this I see a copy-paste error in the third condition. My approaching to test maintenance for this example was less rigorous than the Python example. (I didn't save the tests.) We can guard against that in a long-lived software project by including "quality gates" to ensure that a sufficient portion of our code is actually being exercised by tests.

It's good practice, when you find a bug, to write a test for that code which fails, and then fix it.

Here is the corrected code:

```
(defn two-largest [a b c]
  (cond
    (>= a b c) [a b]
    (>= b a c) [b a]
    (>= c b a) [c b] ; << correction
    (>= a c b) [a c]
    (>= b c a) [b c]
    (>= c a b) [c a]))
```


There are 6 possible ways of organising the 3 variables; by looking at the inequality parameters we see each parameter appears in each "column" twice.

I'll come back to the unordered solution later.

@Beaver
