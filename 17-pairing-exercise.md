**Pairing Exercise: The Berlin Clock.**

Here is the [description](https://en.wikipedia.org/wiki/Mengenlehreuhr) of the clock's behaviour. 

>**Example:** Two fields are lit in the first row (five hours multiplied by two, i.e. ten hours), but no fields are lit in the second row; therefore the hour value is 10. 
Six fields are lit in the third row (five minutes multiplied by six, i.e. thirty minutes), while the bottom row has one field on (plus one minute). Hence, the lights of the clock altogether tell the time as 10:31. (Source: Wikipedia: https://en.wikipedia.org/wiki/Mengenlehreuhr)
>
>**Task**: Write a function that takes in a particular time as 24h format ('hh:mm:ss') and outputs a string that reproduces the Berlin Clock. The parameters should be as follows:
>
>“O” = Light off
>“R” = Red light
>“Y” = Yellow light
>
>**Example Test Case** (this does not correspond to the image above)
Input String:  12:56:01
>
>Output String: O RROO RROO YYRYYRYYRYY YOOO

The interviewer didn't state what the process should be. But test-driven development is a good way to identify scenarios and then implement them progressively.

Let's create a project:

`lein new berlin-clock`

This is the project structure:

```
.
└── berlin-clock
    ├── CHANGELOG.md
    ├── LICENSE
    ├── README.md
    ├── doc
    │   └── intro.md
    ├── project.clj
    ├── resources
    ├── src
    │   └── berlin_clock
    │       └── core.clj
    └── test
        └── berlin_clock
            └── core_test.clj
```

Leiningen gives you one source file to start with (core.clj) and one test. That works for us, since we're required to implement a single function.

Based on the problem statement, our function takes an input string, and outputs a string too.

Let's creating a failing test based on the example:

`core.clj`

```
(defn berlin-clock [time] "")
```

`core_test.clj`

```
(deftest test-initial-case
  (testing "12:56:01"
    (is (= "O RROO RROO YYRYYRYYRYY YOOO" (berlin-clock "12:56:01")))))
```

`$ lein test`

```
$ lein test

lein test berlin-clock.core-test

lein test :only berlin-clock.core-test/test-initial-case

FAIL in (test-initial-case) (core_test.clj:7)
12:56:01
expected: (= "O RROO RROO YYRYYRYYRYY YOOO" (berlin-clock "12:56:01"))
  actual: (not (= "O RROO RROO YYRYYRYYRYY YOOO" ""))

Ran 1 tests containing 1 assertions.
1 failures, 0 errors.
Tests failed.
$ 
```
           
Now let's fix it.

I decided that it wouldn't be sensible to try an implement all the logic in the test case, as it could be challenging to accomplish that within the 45 minutes.

So I opted for a simpler test case.

`core_test.clj`

```
(deftest test-initial-case
  (comment testing "12:56:01"
    (is (= "O RROO RROO YYRYYRYYRYY YOOO" (berlin-clock "12:56:01"))))

  ( testing "midnight"
    (is (= "O OOOO OOOO OOOOOOOOOOO OOOO" (berlin-clock "00:00:00"))))
```

Lining up my new test case with the original I can see that I have the right number of characters. I did initially type zeroes instead of oh's, which one of my "buddies" helpfully pointed out. 

The simplest way to get this passing is to hard-code the return value.

```
(defn berlin-clock [time]
  "O OOOO OOOO OOOOOOOOOOO OOOO")
```

Test passes ...

![Screenshot 2022-06-04 at 08.53.10|690x392](upload://6u6YEekkgJgiHUSIUfnoIOzw6lF.png)


Right now we have a function that works, but only for a _very_ limited set of inputs; if our requirement was to implement a Berlin Clock that returns the correct value for midnight, then we would be done. But it would be correct only once per day (for a second).

(By the way, I commit the updates to version control after every cycle of add failing test / make test pass. Possibly after a refactor too so I can see if I've broken something (and perhaps missed a test case which didn't highlight the issue).

I added a comment at the top of the file to remind me which block is which:

```
;; O OOOO OOOO OOOOOOOOOOO OOOO
;; ^ seconds
;;   ^^^^ every 5 hours
;;        ^^^^ single hours
;;             ^^^^^^^^^^^ minutes
;; 
```

(actually the last line should say "every 5 minutes")

Initially I thought I could simply add an `if` to the implementation:

```
(defn berlin-clock [time]
  (if (= "00:00:00" time)
    "O OOOO OOOO OOOOOOOOOOO OOOO"
    nil))
```

to create a place for the alternate logic.

Here is the test:

```
( testing "5 hours blocks"
    (is (= "O ROOO OOOO OOOOOOOOOOO OOOO" (berlin-clock "05:00:00"))))
```

An a possible implementation:    

```
(defn time-to-clock [:keys [seconds ]]
  {})

(defn berlin-clock [time]
  (let [[hours mins seconds] (map #(Integer/parseInt %) (str/split time #":"))]
    (cond
      (and (zero? hours) (zero? mins) (zero? seconds))
        "O OOOO OOOO OOOOOOOOOOO OOOO"
      (= hours 5)
        "O ROOO OOOO OOOOOOOOOOO OOOO")))
```

Running the test showed me that I had a syntax error in the time-to-clock function, and I wasn't using it, so I deleted it.

```
(ns berlin-clock.core
  (:require [clojure.string :as str]))

(defn berlin-clock [time]
  (let [[hours mins seconds] (map #(Integer/parseInt %) (str/split time #":"))]
    (cond
      (and (zero? hours) (zero? mins) (zero? seconds))
      "O OOOO OOOO OOOOOOOOOOO OOOO"
      (= hours 5)
      "O ROOO OOOO OOOOOOOOOOO OOOO")))
```

I realised that I was going to need to do some conditional logic based on the elements of the time, and having them as numeric (integer) values would be most helpful.

So I imported the `clojure.string` library for the `split` funciton, using the `:` as the separator. 

_As this point, I was around 20 minutes into the pairing exercise, including a few minutes at the beginning to scan the problem statement. Having figured out a sensible structure for the implementation, and seeing that you having misunderstood the problem statement, it becomes easier to add more test cases._ 

**Explanation** This will return a vector of 3 elements

```
berlin-clock.core=> (require '[clojure.string :as str])
nil
berlin-clock.core=> (str/split "12:56:01" #":")
["12" "56" "01"]
berlin-clock.core=> 
```

which we can then map over to convert to integers.

```
berlin-clock.core=> (map #(Integer/parseInt %) (str/split "12:56:01" #":"))
(12 56 1)
berlin-clock.core=>
```

Lastly we can destructure the list elements returned by map:

```
berlin-clock.core=>   (def time "12:56:01")
#'berlin-clock.core/time
berlin-clock.core=>   (let [[hours mins seconds] (map #(Integer/parseInt %) (str/split time #":"))]
               #_=>     (println "hours" hours "minutes" mins "seconds" seconds))
hours 12 minutes 56 seconds 1
nil
berlin-clock.core=>
```

I keep this REPL code in a comment block in the implementation file as it's more convenient for editing.

```
(comment

  (def time "12:56:01")
  (let [[hours mins seconds] (map #(Integer/parseInt %) (str/split time #":"))]
    (println "hours" hours "minutes" mins "seconds" seconds))


  )
```

While coding this, I updated the comment in the implementation file to correct it:

```
;; O OOOO OOOO OOOOOOOOOOO OOOO
;; ^ seconds
;;   ^^^^ every 5 hours
;;        ^^^^ single hours
;;             ^^^^^^^^^^^ 5 minute blocks
;;                         ^^^^ single minutes
```

Now let's add the other 5-hour test cases:

```
( testing "5 hours blocks"
    (is (= "O ROOO OOOO OOOOOOOOOOO OOOO" (berlin-clock "05:00:00")))
    (is (= "O RROO OOOO OOOOOOOOOOO OOOO" (berlin-clock "10:00:00")))
    (is (= "O RRRO OOOO OOOOOOOOOOO OOOO" (berlin-clock "15:00:00")))
    (is (= "O RRRR OOOO OOOOOOOOOOO OOOO" (berlin-clock "20:00:00"))))
```

I recognises that it would be useful to have a function to convert the hour value to the possible 5-hour values: "OOOO", "ROOO", "RROO", "RRRO" and "RRRR".

Here was my first stab:

```
(defn to-clock-5-hour-segments [hour]
  (str
    (if (>= hour 5) "R" "O")
    (if (>= hour 10) "R" "O")
    (if (>= hour 15) "R" "O")
    (if (>= hour 10) "R" "O")))
```

One of my buddies pointed out the copy-and-paste error on the last line. I fixed the 15 case, but not the 20.

This is combining four single-letter strings whose values are based on whether we have crossed a particular threshold.

I knew this wasn't the "best" implementation, as it's fairly repetitive, but it is clear. And I could come back and refactor later (which I vocalized during the interview).

Here is the updated implementation which makes the tests pass:

```
(defn to-clock-5-hour-segments [hour]
  (str
    (if (>= hour 5) "R" "O")
    (if (>= hour 10) "R" "O")
    (if (>= hour 15) "R" "O")
    (if (>= hour 20) "R" "O")))

(defn berlin-clock [time]
  (let [[hours mins seconds] (map #(Integer/parseInt %) (str/split time #":"))]
    (cond
      (and (zero? hours) (zero? mins) (zero? seconds))
      "O OOOO OOOO OOOOOOOOOOO OOOO"
      :else
      (str "O" " "
           (to-clock-5-hour-segments hours) " "
           "OOOO OOOOOOOOOOO OOOO"))))
```               

The new function is generating part of the output string dynamically, and we are using `str` to combine it with the other hard-coded parts of the string.

_Roughly 30 minutes elapsed._

I'm taking a break now, and will add the rest in my next post.

@beaver