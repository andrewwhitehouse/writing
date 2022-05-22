**Test-Driven Development in Python**

It's time to create a Python version of the exercise we wrote in Clojure.

First let's create the library with a failing test:

```
├── test_two_largest.py
└── two_largest
    ├── __init__.py
```

Unrelated files have been removed.

Here are the files:

```
import unittest
from two_largest import *

class TestTwoLargest(unittest.TestCase):

    def test_two_largest(self):
        self.assertEqual([3,2], two_largest(3,2,1))

if __name__ == '__main__':
    unittest.main()
```

```
def two_largest(*nums):
  return [0,0]
```

Running the tests:

```
Andrews-MacBook-Pro:python andrewwhitehouse$ python -m unittest test_two_largest.py 
F
======================================================================
FAIL: test_two_largest (test_two_largest.TestTwoLargest)
----------------------------------------------------------------------
Traceback (most recent call last):
  File "/Users/andrewwhitehouse/code/writing/code/python/test_two_largest.py", line 7, in test_two_largest
    self.assertEqual([3,2], two_largest(3,2,1))
AssertionError: Lists differ: [3, 2] != [0, 0]

```  
              
For this exercise we're taking a test-first approach, also known as _Test-Driven Development_.

We start with a deliberately failing test, first it, refactor if needed, and add another test scenario until we've covered sufficient scenarios.

First fix.

```
def two_largest(*nums):
  return [nums[0], nums[1]]
```

It's hard-coded and is only going to work for the specific test case.

What is the first two numbers are in the wrong order?

```
    def test_first_two(self):
        self.assertEqual([3,2], two_largest(3,2,1))

    def test_first_two_out_of_order(self):
        self.assertEqual([3,2], two_largest(2,3,1))
```

Re-run the tests 

`$ python -m unittest test_two_largest.py `


```
Andrews-MacBook-Pro:python andrewwhitehouse$ python -m unittest test_two_largest.py 
.F
======================================================================
FAIL: test_first_two_out_of_order (test_two_largest.TestTwoLargest)
----------------------------------------------------------------------
Traceback (most recent call last):
  File "/Users/andrewwhitehouse/code/writing/code/python/test_two_largest.py", line 10, in test_first_two_out_of_order
    self.assertEqual([3,2], two_largest(2,3,1))
AssertionError: Lists differ: [3, 2] != [2, 3]
```
  
Fix it

```
def two_largest(*nums):
  if nums[1] > nums[0]:
    return [nums[1], nums[0]]
  else:
    return [nums[0], nums[1]]
```

```
$ python -m unittest test_two_largest.py 
..
----------------------------------------------------------------------
Ran 2 tests in 0.000s

OK
$ 
```

The next test case involves the largest numbers being the second and third. Let's change the values used to show that we can handle values other than 1, 2 and 3.

```
    def test_last_two_largest_in_order(self):
        self.assertEqual([6,5], two_largest(4,6,5))
```
  
(Run tests again; newly added test fais)

```
def two_largest(*nums):
  if nums[1] > nums[0]:
    return [max(nums[1], nums[2]), max(nums[2], nums[0])]
  else:
    return [nums[0], nums[1]]
```

The nested 'ifs' are starting to become more difficult to reason about so it makes sense to use the max function to figure out the correct ordering.

I started to use the max function in both elements of the returned value. But on running the test I realised this wouldn't work. We're returning a two-element array and the first element is the maximum of nums[1] and nums[2]. If nums[2] is larger, then the second element should be nums[1]; if nums[1] is greater then the second element is the larger of nums[2] and nums[0]. As it's written, the code doesn't reflect that dependency between the elements.

New failing test ...

```
    def test_last_two_out_of_order(self):
        self.assertEqual([9,8], two_largest(7,8,9))
```

Make the change ...

```
def two_largest(*nums):
  if nums[1] > nums[0]:
    if nums[2] > nums[1]:
      return [nums[2], nums[1]]
    else:
      return [nums[1], max(nums[2], nums[0])]
  else:
    return [nums[0], nums[1]]
```

Next failing test ...

```
    def test_first_and_last_in_order(self):
        self.assertEqual([102, 101], two_largest(102, 100, 101))
```

First attempt failed an earlier test:

```
def two_largest(*nums):
  if nums[1] > nums[0]:
    if nums[2] > nums[1]:
      return [nums[2], nums[1]]
    else:
      return [nums[1], max(nums[2], nums[0])]
  else:
    if nums[0] > nums[2]:
      return [nums[0], nums[2]]
    else:
      return [nums[0], nums[1]]
```

```
Traceback (most recent call last):
  File "/Users/andrewwhitehouse/code/writing/code/python/test_two_largest.py", line 7, in test_first_two
    self.assertEqual([3,2], two_largest(3,2,1))
AssertionError: Lists differ: [3, 2] != [3, 1]
```

This is better:

```
def two_largest(*nums):
  if nums[1] > nums[0]:
    if nums[2] > nums[1]:
      return [nums[2], nums[1]]
    else:
      return [nums[1], max(nums[2], nums[0])]
  else:
    if nums[0] > nums[2]:
      return [nums[0], max(nums[2], nums[1])]
    else:
      return [nums[0], nums[1]]
```            
    
We have a couple of branches which are not doing a max in the second argument. Is that right? :thinking: 

On line 4, what test case could we select to make it fail?

num1 > num0, and num2 > num1, but the result should be num2, num0. Hmm. That looks like a well-defined ordering num2 > num1 > num0.

What about the last line:

num0 >= num1, and num2 >= num0

so num2 >= num0 >= num1.

So it _should_ be num2, num0.

Let's add a test case 501,500,502. This is the case where it's the first and third and they are out of order.

```
    def test_first_and_last_out_of_order(self):
        self.assertEqual([502, 501], two_largest(501, 500, 502))
```

This test fails (which is good):

```      
Traceback (most recent call last):
  File "/Users/andrewwhitehouse/code/writing/code/python/test_two_largest.py", line 22, in test_first_and_last_out_of_order
    self.assertEqual([502, 501], two_largest(501, 500, 502))
AssertionError: Lists differ: [502, 501] != [501, 500]
```

Correct the last line:

```
def two_largest(*nums):
  if nums[1] > nums[0]:
    if nums[2] > nums[1]:
      return [nums[2], nums[1]]
    else:
      return [nums[1], max(nums[2], nums[0])]
  else:
    if nums[0] > nums[2]:
      return [nums[0], max(nums[2], nums[1])]
    else:
      return [nums[2], nums[0]]
```

This solution has evolved differently to the Clojure solution where we didn't use test-driven development.

```
(defn two-largest [a b c]
  (let [max-a-b (max a b)]
     (if (> c max-a-b)
       [c max-a-b]
       [max-a-b (max (min a b) c)])))
```

In the Clojure example the use of the `max-a-b` changes how we would reason about the correctness of the code.

Let's finish this exercise by adding some more cases.

```
    def test_all_equal(self):
        self.assertEqual([123, 123], two_largest(123, 123, 123))

    def test_negative_numbers(self):
        self.assertEqual([-10, -20], two_largest(-10, -20, -30))
```

The test with negative numbers checks that we're not using the absolute values. (This is a regression test, really, as the code was already doing that. The regression test stops someone breaking it later.)

One more case:

```
    def test_decimals(self):
        self.assertEqual([2.001, 2.0003], two_largest(2.001, 2.0002, 2.0003))
```

Note that the first number is actually larger than the second.

@beaver