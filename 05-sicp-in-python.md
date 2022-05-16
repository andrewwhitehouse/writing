**SICP in Python**

In SICP one of the early [examples](https://mitpress.mit.edu/sites/default/files/sicp/full-text/book/book-Z-H-10.html#%_sec_1.1.1) given is calculating the circumference of a circle.

The approach taken (in Scheme) is an interactive one:

```
$ python
Python 3.6.10 (default, Feb  6 2020, 15:10:56) 
[GCC 4.2.1 Compatible Apple LLVM 10.0.0 (clang-1000.10.44.4)] on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>> radius = 10
>>> type(radius)
<class 'int'>
>>> import math
>>> circumference = 2 * math.pi * radius
>>> circumference
62.83185307179586
>>> 
```

Running Python without an argument opens it in interactive mode, which is known as the REPL. `REPL` is short for Read-Eval-Print-Loop; you are given a Python runtime where you can add definitions and try things out, for quick feedback. So it's a good way to explore the language.

However, once you close the REPL your code is gone. (Although you can cycle through the history and execute the commands again.) So it makes sense to save your code in more persistent form, such as a local file. (And keep it under version control like _git_, but that's a topic for a different post.)

And writing automated tests allow you to check quickly that your code is still working.

I followed [this example](https://realpython.com/python-testing/) for setting up the tests.

Here are the initial folder contents: 

```
.
├── chapter1
│   ├── __init__.py
│   └── chapter1.py
└── test_chapter1.py
```

`chapter1/__init__.py`

```
import math
  
def circumference(radius):
  return 2 * math.pi * radius
```

`test_chapter1.py`

```
import unittest
from chapter1 import circumference

class TestChapter1(unittest.TestCase):

    def test_circumference(self):
        self.assertEqual(circumference(10), 62.8318)

if __name__ == '__main__':
    unittest.main()
```

```
$ python test_chapter1.py 
F
======================================================================
FAIL: test_circumference (__main__.TestChapter1)
----------------------------------------------------------------------
Traceback (most recent call last):
  File "test_chapter1.py", line 7, in test_circumference
    self.assertEqual(circumference(10), 62.8318)
AssertionError: 62.83185307179586 != 62.8318

----------------------------------------------------------------------
Ran 1 test in 0.001s

FAILED (failures=1)
Andrews-MacBook-Pro:python andrewwhitehouse$ 
```

You can't compare decimals directly because their representations are not exact. 

So, change the test assertion:

```
self.assertAlmostEqual(circumference(10), 62.8318, places=3)
```

```
$ python test_chapter1.py 
.
----------------------------------------------------------------------
Ran 1 test in 0.000s

OK
$
```

:tada:
