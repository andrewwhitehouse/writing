**Sliding Windows**

For part 2 we are asked to adapt the solution to use a sliding window.

>Considering every single measurement isn't as useful as you expected: there's just too much noise in the data.
>
>Instead, consider sums of a three-measurement sliding window. Again considering the above example:
>
```
199  A      
200  A B    
208  A B C  
210    B C D
200  E   C D
207  E F   D
240  E F G  
269    F G H
260      G H
263        H
```
>Start by comparing the first and second three-measurement windows. The measurements in the first window are marked A (199, 200, 208); their sum is 199 + 200 + 208 = 607. The second window is marked B (200, 208, 210); its sum is 618. The sum of measurements in the second window is larger than the sum of the first, so this first comparison increased.
>
>Your goal now is to count the number of times the sum of measurements in this sliding window increases from the previous sum. So, compare A with B, then compare B with C, then C with D, and so on. Stop when there aren't enough measurements left to create a new three-measurement sum.

We have the opportunity to implement this by reusing code from part 1.

For part 1 we calculate the result with `CountIncreases(Parse(input))`.

We are using functions as a means of abstraction to break our solution down into well-defined pieces.

Part 2 introduces the concept of a sliding window. We can calculate this by progressively summing triplets of the original array.

So the calculation for part 2 will look like this: `CountIncreases(SlidingWindow(Parse(input))`.

Here is the code:

`day1_test.go`

```
func TestSlidingWindow(t *testing.T) {
	depths := []uint16{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := []uint16{607, 618, 618, 617, 647, 716, 769, 792}
	actual := SlidingWindow(depths)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("SlidingWindow was incorrect, got: %d, want: %d.", actual, expected)
	}
}
```

(again, I'm using the test case example from the problem statement.)

And the implementation:

```
func SlidingWindow(depths []uint16) []uint16 {
	ret := make([]uint16, len(depths)-2)
	for i := 2; i < len(depths); i++ {
		ret[i-2] = depths[i] + depths[i-1] + depths[i-2]
	}
	return ret
}
```

@beaver

