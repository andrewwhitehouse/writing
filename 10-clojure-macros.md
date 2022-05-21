The book has a [section](https://mitpress.mit.edu/sites/default/files/sicp/full-text/book/book-Z-H-10.html#%_sec_1.1.5) on applicative versus normal order evaluation.

Clojure takes the "applicative order" approach which is to evalute the arguments first and then apply.

```
(defn p [] (p))

(defn test [x y]
  (if (= x 0) 0 y))

(test 0 (p))
```

In this code, the `p` function calls itself unconditionally. And later we call the test function with its first argument as zero, and the second argument as a call the the p function.

Executing this code causes the second argument to test to be evaluated, which calls itself going into an infinite loop and eventually falling over with a stack overflow error, since the call stack has a finite capacity.

```
sicp.chapter1=> (test 0 (p))
Execution error (StackOverflowError) at sicp.chapter1/p (form-init3631130360703580959.clj:1).
null

sicp.chapter1=>
```

There are a couple of ways we can work around this issue in Clojure.

The first is with Clojure _macros_.

The Clojure interpreter involves a [Reader](https://clojure.org/reference/reader) which can apply code transformations through macros. These transformations allow us to define new features.

In the above example we can define test as a macro instead:

```
(defmacro test [x y]
   `(if (= ~x 0) 0 ~y))
```          

Macros highlight a significant benefit in Clojure of code and data having the same representation. We can use code (macros) to manipulate data through the Reader, to produce code which is then executed. 

The Reader applies this as a transformation  

```
sicp.chapter1=> (defmacro test [x y]
           #_=>    `(if (= ~x 0) 0 ~y))
#'sicp.chapter1/test
sicp.chapter1=> (macroexpand-1 '(test 0 (p)))
(if (clojure.core/= 0 0) 0 (p))
sicp.chapter1=>
```

The macros effectively applies a substitution on its arguments, and it is the result that is actually executed. So in this case the code no longer executes the `p` function, avoiding the infinite loop.

A second way we can defer execution is by passing the second argument as a function to be executed from `test` only when needed:

```
sicp.chapter1=> (defn p [] (p))
#'sicp.chapter1/p
sicp.chapter1=> (defn test [x y]
           #_=>   (if (= x 0) 0 (y)))
#'sicp.chapter1/test
sicp.chapter1=> (test 0 p)
0
sicp.chapter1=>
```

Passing around a function that loops infinitely is not going to be very useful in practice. However it is feasible that a function argument leads to an operation that has an overhead, and we want to evaluate only in some cases.

