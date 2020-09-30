package main

import (
	"fmt"
)

var A = 4
var B = A - 1

func lol(A int) {
	for A > 0 && B > 0 {
		A = A * B
		B = B - 1
		fmt.Println(lol(A))
	}
}
