package iota

import (
	"fmt"
)

const (
	first  = iota + 6
	second = iota
)

const (
	third = iota
	fourth
)

func iota() {
	fmt.Println(first, second, third, fourth)
}
