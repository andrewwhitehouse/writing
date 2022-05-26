Square Roots using Newton's Method in Clojure

[_Still practising my Clojure chops as there is a possible interview looming._]

First of all, it's worth noting that Clojure already has a sqrt function available, which is available through Java or [JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math) interop. (We'll cover more in Java interop in a later chapter.)

The purpose of this chapter is to explore building up a solution by composing functions, and to introduce some mathematical [history](https://en.wikipedia.org/wiki/Newton%27s_method#History) through code.

```
$ lein repl
nREPL server started on port 55041 on host 127.0.0.1 - nrepl://127.0.0.1:55041
<snip>

user=> Math/PI
3.141592653589793
user=> (Math/sqrt 2)
1.4142135623730951
user=> ```
```

[_Note, would it be easier if I did these examples in repl.it? Let's try that._]

Go to https://replit.com/~ and sign up (you can login with a Google or other account).

Add image here..

Click Create, search for "clojure" in the search box, and "Create Repl".

You're given a main.clj to work with, which is good for us.

SICP describes it's approach top-down ... in other words by referring to functions that don't yet exist.

Let's do ours bottom up; define functions that we can interact with in the REPL.

Repl.it's approach, as an online tool, is actually a little different to how you would interact with a REPL on your local machine, but it works for our example.

Type in the following code:

```
(defn average [x y]
  (/ (+ x y) 2))

(println (average 3 4))
```

You're about to run this code, but before you do ask yourself what you would expect to see ... hopefully 3.5(?).

When you click the green Run button the REPL will take some time to download dependencies and start the runtime environment; then you should see:

```
> clojure -M main.clj
7/2
> 
```

Hmm. What's going on here?

Let's replace our print line:

```
(def result (average 3 4))

(println (type result) result)
```

Try it. What do you see? 

There are a couple of things going on here.

Clojure is running on top of the Java runtime.

If we were to calculate the average in Java it would look something like this:

```
class Main {
  public static void main(String[] args) {
    System.out.println((3 + 4) / 2);
  }
}
```

If you use integer values in a calculation, Java infers that you want an integer result, and so drops any fractional part. You can correct this by converting one of the values to a float, and then Java implicitly converts the expression to use floats.

```
class Main {
  public static void main(String[] args) {
    System.out.println((3 + 4) / 2.0);
  }
}
```

Clojure takes a different approach, and returns a [Ratio](https://clojure.org/reference/data_structures) of the integers, reduced to it's simpest form.

[_As an aside, note that we need the extra brackets in the calculation `(3+4)/2` because the `/` operator has higher precedence than '+'. Without the brackets, the calculation `3+4/2` would be intrepreted as '4/2' plus '3', or '5'_]

To obtain the result as a float, you convert one of the values to a float as in Java:

```
(defn average [x y]
  (/ (+ x y) 2.0))

(def result (average 3 4))

(println (type result) result)
```

In your Repl.it console you should see:

```
> clojure -M main.clj
java.lang.Double 3.5
> 
```

The `good-enough` function uses abs and square.

Let's define these and call them.

```
(defn average [x y]
  (/ (+ x y) 2.0))

(def result (average 3 4))

(println (Math/abs 2.3) (Math/abs -2))

(defn square [x] (* x x))

(println (square 2.2) (square 5))

(println (type result) result)
```

Output:

```
> clojure -M main.clj
2.3 2
4.840000000000001 25
java.lang.Double 3.5
> 
```

```
(defn good-enough? [guess x]
  (< (Math/abs (- (square guess) x)) 0.001))
```

[_Note that LISP functions allow the use of question marks, which are useful when defining functions that provide some sort of conditionla logic, known as a "predicate"_]

How can we call this function? We want to know if the guess is good enough as an approximation for the square root of x.

```
(defn average [x y]
  (/ (+ x y) 2.0))

(def result (average 3 4))

(comment println (Math/abs 2.3) (Math/abs -2))

(defn square [x] (* x x))

(comment println (square 2.2) (square 5))

(comment println (type result) result)

(defn good-enough? [guess x]
  (< (Math/abs (- (square guess) x)) 0.001))

(println (good-enough? 2.2 4.84))
```

I've commented out some of the unused code to reduce the noise in the output. You can delete it too. Repl.it supports Control-Z (or CMD-Z on a Mac) for undo, and Shift-Control-Z for redo.

Call your function with different parameters to make it return both true and false.











TODO:

- 
[_Update this example to demonstrate REPL-driven development._]