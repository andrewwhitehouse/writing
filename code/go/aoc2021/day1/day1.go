package day1

/*
Note: type comes after variable.
Iterating over array. Could have used rang, but wanted to start from 1.
variable declaration and assignment.
Indentation: is it 4?
The function name looks like Pascal
*/

func CountIncreases(depths []int16) uint16 {
    var increases uint16 = 0
    for i := 1; i < len(depths); i++ {
        if depths[i] > depths[i-1] {
            increases++
        }
    }
    return increases
}
