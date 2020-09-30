package main

import (
	"fmt"
	"os"
)

func main() { // HIntExclusion1
	a := os.Args[1:]
	s := "123456789"

	for range os.Args[1:] {
		for y := 1; y <= len(a); y++ {
			for x := 0; x <= 9; x++ {
				if s[x] == a[y] {
					s[x] = ""
					fmt.Println(s)
				}
			}
		}
	}
}
