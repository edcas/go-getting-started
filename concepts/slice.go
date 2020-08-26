package slice

import (
	"fmt"
)

func main() {
	arr := [3]int{1, 2, 3}
	slice := arr[:]
	fmt.Println(arr, slice)

	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice)

	sliceOther := []int{1, 2, 3}
	fmt.Println(sliceOther)

	sliceOther = append(slice, 1, 1, 1)
	fmt.Println(sliceOther)

	s2 := sliceOther[1:]
	s3 := sliceOther[:2]
	s4 := sliceOther[1:2]
	fmt.Println(s2, s3, s4)
}
