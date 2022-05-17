Let's define a function to perform a simple calculation, and give it some values:

```
$ python
Python 3.6.10 (default, Feb  6 2020, 15:10:56) 
[GCC 4.2.1 Compatible Apple LLVM 10.0.0 (clang-1000.10.44.4)] on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>> def square(x):
  return x*x
... 
>>> result = square(2)
>>> result
4
>>> type(result)
<class 'int'>
>>> result2 = square(2.0)
>>> result2
4.0
>>> type(result2)
<class 'float'>
>>> 
```

Our function returns the square of its argument. if we pass it the value 2, which is an `integer` (a number with no fractional part), we receive an integer back.

In the second case, we call the function with a decimal number which happens to have the same value, but because we add the ".0" Python knows to treat it as a float, and so it returns the result of modifying two floats. Multiplication is defines for integers and floats.

What if we try a non-numeric value?

```
>>> square("hello")
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
  File "<stdin>", line 1, in square
TypeError: can't multiply sequence by non-int of type 'str'
>>> 
```

So Python will attempt to apply the operations in your function to the data types you give it. 

What if we change the definition slightly?

```
>>> def multiply(x,y):
...   return x*y
... 
>>> multiply(2,2)
4
>>> multiply(2.0, 2)
4.0
>>> multiply("hello", 2)
'hellohello'
>>> 
```

You can see that the result of multiplying a float and an integer is a float. Python [converts](https://docs.python.org/3/reference/expressions.html#arithmetic-conversions) the values to a common type. And the `*` operator works with a string and a number to repeat that string.
