**Numbers**

The first abstraction we're going to explore are numbers, one of the common data types.

That box on your desk, or in the cloud, deals in ones and zeros; base 2. 

The base of a number determines the number of possible values each element can contain: the base minus one.

So base 2 numbers can have the value 1 or 8.

Base 8 (also known as "octal") digits have the values 0 to 7.

Decimal numbers, containing digits from zero to nine, are base 10.

And hexadecimal values, contains characters that can have the value 0-9 and A-F, giving 16 possible values.

But under the covers it's stlil a combination of ones and zeros.

Other bases are possible, including Base 64; this is used for safely sending data over networks by encoding it first.

There are schemes called _big endian_ and _little endian_ which indicate the order in which bits are stored.

If the number 65 is stored as 00...001000001 this is _big endian_ because the higher value bits are stored first, and the lower value bits last.

This represents 65 because the different binary digits (bits) correspond to a particular power of two.

```
64 32 16  8  4  2  1
 1  0  0  0  0  0  1
```

So the value is 64 + 1 = 65.

Counting these ones and zeros can be challenging, but is easier if you treat them as groups of the same value.

If you need to store only positive numbers then the maximum value you can store in 4 bits, for example, is 1+2+4+8 = 15, which is 2^4-1. (2^4 is two the power 4, or 2*2*2*4 = 16).

So the maximum value you can store in 32 bits is 2^32-1.

As a useful shortcut, 2^10 is 1024 (or approximately 10^3). 

2^20 = 2^(10+10) = 2^10 * 2^10 which is approximately 1,000,000.

So every increase in power of 10 multiplies the result by 1,000. This can be useful to know when sanity checking your work.

If you want to store negative numbers, the maximum value you can store is roughly halved because one of the bits is used to indicate whether the number is positive or negative.

The [Fibonacci sequence](https://en.wikipedia.org/wiki/Fibonacci_number) is defined as:

```
F(0) = 0
F(1) = 1
F(n) = F(n-1) + F(n-2)
```

Some people incorrectly start the sequence from 1.

This number grows rather large rather quickly.

Here is the code:

```
package fib
  
func Fibonnacci(n uint16) uint64 {
        if n < 2 {
                return uint64(n)
        } else {
                var a, b uint64 = 0, 1
                for i := uint16(1); i < n; i++ {
                        b, a = a+b, b
                }
                return b
        }
}
```

and tests

```
package fib
  
import (
    "testing"
    "aoc2021/fib"
)

func TestCountIncreases(t *testing.T) {
    n := []uint16{0,1,2,3,4,5,6};
    expected := []uint64{0,1,1,2,3,5,8};
        for i := 0; i < len(n); i++ {
        actual := fib.Fibonnacci(n[i])
        if expected[i] != actual {
            t.Errorf("Fibonacci was incorrect, got: %d, want: %d.", actual, expected[i])
        }
        }       
}
```

Exercises:

What is the largest value you can store in a 64 bit unsigned integer?

And what is the value of n for which our function will correctly calculate the Fibonacci value?

What happens if you increase the value of n by 1 beyond this value

Can you add validation to the function so that it returns an error value if the parameter exceeds the maximum supported value?



