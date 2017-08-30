package main

import "fmt"

func main() {
	secondLast := 1
	last := 2
	sum := last //only the last (2) is even

	for last < 4e6 {
		last, secondLast = secondLast+last, last
		if last%2 == 0 {
			sum += last
		}
	}

	fmt.Println(sum)
}
