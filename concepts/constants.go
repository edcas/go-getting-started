package constants

import (
	"fmt"
)

func constants() {
	const pi = 3.1415
	fmt.Println(pi)

	// pi = 1.2 cannot assign to pi

	const standard = 3
	fmt.Println(standard)
	fmt.Println(standard + 1.2) // Inference

	const standardInt int = 3
	fmt.Println(float32(standard) + 1.2)
	// fmt.Println(standard + 1.2) truncate to integer
}
