**Exercise 1.3 in Clojure**  

> Define a procedure that takes three numbers as arguments and returns the sum of the squares of the two larger numbers.

I'm going to focus on the code to identify the two largest numbers.

Here was my first attempt, using named elements in a map.

```
;; First attempt
(defn two-largest [a b c]
  (let [first-pair (if (> a b) {:larger a :smaller b} {:larger b :smaller a})
        remaining-result (if (> c (:larger first-pair))
                           {:first c :second (:larger first-pair)}
                           (if (> c (:smaller first-pair) )
                             {:first (:larger first-pair) :second c}
                             {:first (:larger first-pair) :second (:smaller first-pair)}))]
    [(:first remaining-result) (:second remaining-result)]))
```

which is fairly gnarly and unreadable.

The key point is we need to identify which is the larger of the comparison in the first pair, and then compare that with the remaining value.

Oh, and I have written some tests:

```
(deftest test-two-largest
  (testing "all different"
    (is (= [3 2] (ch1/two-largest 1 2 3))))
  (testing "two equal"
    (is (= [2 1] (ch1/two-largest 1 2 1)))
    (is (= [2 1] (ch1/two-largest 1 1 2))))
  (testing "all equal"
    (is (= [3 3] (ch1/two-largest 3 3 3)))))
```

Let's try it using vectors.

```
;; Second attempt
(defn two-largest [a b c]
  (let [a-b-comparison (if (> a b) [a b] [b a])]
    (if (> c (first a-b-comparison))
      [c (first a-b-comparison)]
      (if (> c (second a-b-comparison))
        [(first a-b-comparison) c]
        a-b-comparison))))
```

I like this better, but there is still plenty of repetition that makes it harder to read.

Clojure has a `max` function which we can use:

```
;; Third ...
(defn two-largest [a b c]
   (let [max-a-b (max a b)
         min-a-b (min a b)]
     (if (> c max-a-b)
       [c max-a-b]
       [max-a-b (max min-a-b c)])))
```         

But I keep thinking we can improve the readability further by sorting the elements in descending order:

```
;; Fourth ...
(defn two-largest [& nums]
  (let [sorted (sort > nums)]
    [(first sorted) (second sorted)]))
```

There are a few things going on here. The ampersand in the arguments converts our parameters into a sequence.

```
sicp.chapter1=> (defn foo [& args] args)
#'sicp.chapter1/foo
sicp.chapter1=> (foo 1 2 3)
(1 2 3)
sicp.chapter1=>
```

We then sort that sequence and create a (scoped) let binding; the first element of the result is the largest and the second is the next largest. We ignore the other elements.

There are a couple of small improvements we could make. It's subjective how helpful these are:

```
;; Fifth
(defn two-largest [& nums]
  (let [first-two-values (fn [[a b]] [a b])]
    (first-two-values (sort > nums))))
```

This creates a local function that takes a collection and uses a technique called [destructuring](https://clojure.org/guides/destructuring) to bind the collection elements to specific variables. Note that the function arguments have an extra bracket.

```
sicp.chapter1=> (defn first-two-values [[a b]] [a b])
#'sicp.chapter1/first-two-values
sicp.chapter1=> (first-two-values [5 4 3 2 1])
[5 4]
sicp.chapter1=>
```

We can also bypass the local function by using a "thread-right" macro. This modifies the execution order of the expression.

The thread-right [macro](https://clojure.org/guides/threading_macros) takes the first item in the list and "threads" it as the last element of the next list (which it creates if needed), and then threads that result as the last element of the next list.

For example:

```
sicp.chapter1=> (macroexpand-1 '(->> '1 '2 '3))
(quote 3 (quote 2 (quote 1)))
sicp.chapter1=>
```

So the thread-right macro performs a code translation in the Reader step. Whether or not the result makes any sense when it comes to be executed is another matter.

Here's our version using the thread-right macro:

```
;; Sixth attempt
(defn two-largest [& nums]
  (->> (sort > nums)
       ((fn [[a b]] [a b]))))
```

This takes the result of sorting the numbers and then passes that as the argument to the function we've defined which returns a vector of the first two elements.

`macroexpand-1` is our friend again.

```
sicp.chapter1=> (macroexpand-1 '(->> (sort > nums) ((fn [[a b]] [a b]))))
((fn [[a b]] [a b]) (sort > nums))
sicp.chapter1=>
```     

The first part of the expression is an anonymous function which we call with the result of the sorting operation.

As I've been accumlating potential solutions I've been using a feature of Clojure: comment blocks.

```
(comment defn two-largest [& nums]
  (let [first-two-values (fn [[a b]] [a b])]
    (first-two-values (sort > nums))))

(defn two-largest [& nums]
  (->> (sort > nums)
       ((fn [[a b]] [a b]))))
```

The bracketed constructs are referred to as Clojure _forms_ and we can comment out an entire form by using the comment macro; this causes the entire form to evaluate to nil.

How do I know? Because Clojure lets us look at the source code.

```
user=> (source comment)
(defmacro comment
  "Ignores body, yields nil"
  {:added "1.0"}
  [& body])
nil
user=>
```

I think it's a close call as to which of the solutions is preferable. Some may feel that sorting the elements is over-engineering, but it's a standard operation and I believe it improves readability. I like the fourth and fifth because the intent is clear. And the last one wins for conciseness but it's easy to miss the extra define-and-call-function bracket, and the destructuring.

When I'm being paid to deliver functionality for a client I tend to stop at "it's readable, not over-engineered, and it works". These exercises allow us to go further and explore the trade offs.

@beaver

    
