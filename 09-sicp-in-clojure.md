I'm experimenting with another language to compare with the Python interpretation of SICP. (If you didn't read the previous post(s), SICP is a classic work of computer science called Structure and Interpretation of Computer Programs [sic] :) )

Also there is a Clojure contract role coming up, and I've been thinking why am I not making more of all the effort I've taken to learn Clojure.

And ... back to the code ...

---

We've previously seen a Python implementation of a square function. 

```
def square(x):
  return x*x
```

Thie code calls the multiply operator on x with the argument of x, again. When called with an integer, like `123`, the function will return an integer. When called with a float, like `123.4` the function will return a float.

Down at the hardware level data types are represented as 1s and 0s, and we can't represent all decimal values precisely. 

For example:

```
$ python
Python 3.6.10 (default, Feb  6 2020, 15:10:56) 
[GCC 4.2.1 Compatible Apple LLVM 10.0.0 (clang-1000.10.44.4)] on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>> 0.1 + 0.2
0.30000000000000004
>>> 
```

So when comparing float values we need to check that their difference is smaller than our allowed tolerance. In this case you could check that the difference is less than 0.000001. Actually you could do more, but do you really want to be managing constants in your code like 0.000000000000001. How many zeroes is that?

A language that has similar semantics to both Python and Scheme is Clojure.

The Clojure interpreter and compiler runs on top of either the Java Virtual Machine or the JavaScript V8 runtime used by Nodejs.

We're using the Java runtime.

Let's create a project:

`lein new sicp`

By default we get this directory structure:

```
.
└── sicp
    ├── CHANGELOG.md
    ├── LICENSE
    ├── README.md
    ├── doc
    │   └── intro.md
    ├── project.clj
    ├── resources
    ├── src
    │   └── sicp
    │       └── core.clj
    └── test
        └── sicp
            └── core_test.clj
```

Because we're organising our code in chapters I'll be renaming those core files.

The `lein` tool is actually called Leiningen. This is a literary reference to a Carl Stephenson work "Leiningen Versus the Ants". Ant is also a build tool that many of us has encountered in our careers.

`src/sicp/chapter1.clj`

```
(ns sicp.chapter1)
  
(defn square [x]
  (* x x))
```

This file defines a namespace within which any functions and other values are defined. We use `defn` to define a function, and function parameters are enclosed in square brackets.

I can load this file in a REPL by first tweaking my project.clj:

```
  :repl-options {:init-ns sicp.chapter1} 
```

```
$ lein repl
nREPL server started on port 56605 on host 127.0.0.1 - nrepl://127.0.0.1:56605
REPL-y 0.4.3, nREPL 0.6.0
Clojure 1.10.0
OpenJDK 64-Bit Server VM 11.0.7+10-LTS
    Docs: (doc function-name-here)
          (find-doc "part-of-name-here")
  Source: (source function-name-here)
 Javadoc: (javadoc java-object-or-class-here)
    Exit: Control+D or (exit) or (quit)
 Results: Stored in vars *1, *2, *3, an exception in *e

sicp.chapter1=> (square 2)
4
sicp.chapter1=> (square 2.2)
4.840000000000001
sicp.chapter1=> (type (square 2))
java.lang.Long
sicp.chapter1=> (type (square 2.2))
java.lang.Double
sicp.chapter1=> 
```  

To answer the question "is it (still) working" we can add a couple of tests.

`test/sicp/chapter1_test.clj`

```
(ns sicp.chapter1-test
  (:require [clojure.test :refer :all]
            [sicp.chapter1 :refer :all]))

(defn is-close [a b tolerance]
  (< (Math/abs (- a b)) tolerance))

(deftest test-square
  (testing "square integer."
    (is (= 9 (square 3))))
  (testing "square decimal"
    (is (is-close (square 2.2) 4.84 0.001))))
```

The namespace definition pulls in the libraries from the `clojure.test` namespace as well as the `sicp.chapter1` namespace where we've defined our square function. Because the functions are required with `:refer :all` the square function can be called directly. It also means we couldn't define a test with the same name (e.g. `deftest square`) because the names would clash.

I tend to prefix test names with `test-` which can end up being redundant.

The alternative is to change the way we reference the namespace being tested:

```
(ns sicp.chapter1-test
  (:require [clojure.test :refer :all]
            [sicp.chapter1 :as ch1]))

(defn is-close [a b tolerance]
  (< (Math/abs (- a b)) tolerance))

(deftest test-square
  (testing "square integer."
    (is (= 9 (ch1/square 3))))
  (testing "square decimal"
    (is (is-close (ch1/square 2.2) 4.84 0.001))))
```

I borrowed the `is-close` definition from [Python](https://stackoverflow.com/questions/558216/function-to-determine-if-two-numbers-are-nearly-equal-when-rounded-to-n-signific).

Note that Clojure tends to adopt lower case names for functions, and separates works with hyphens: this is called "kebab case".

Let's run the tests:

```
$ lein test

lein test sicp.chapter1-test

Ran 1 tests containing 2 assertions.
0 failures, 0 errors.
$
```        

Writing Clojure actually makes me happy; I can get a lot done with a small amount of code. It is elegant, expressive and concise. It's like a superpower. The language is beautifully designed. And Rich Hickey who designed it is a genius. It hasn't gained the widespread adoption that some other languages have; for large projects companies tend to prefer static typing, which makes refactoring easier. And because there are fewer people with the skill, fewer companies use it. 

You can mitigage the static typing issue with more tests. But that requires discipline, which requires experience. (See above point.)

I've heard it said that some of the community can be quite opinionated, to the point of being prickly, which has made it less accessible. 

Paul Graham, of Y Combinator, has written about [LISP](http://www.paulgraham.com/diff.html) and its [suitability for startups](http://www.paulgraham.com/avg.html).

@beaver