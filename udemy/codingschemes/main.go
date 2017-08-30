package main

import "fmt"

func main() {
	fmt.Println(42)
	fmt.Printf("%d (decimal) = %b (binary) \n", 42, 42)
	fmt.Printf("%d (decimal) - %b (binary) - %#x (hexadecimal w/ notation) \n", 42, 42, 42)
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %#x \t %q \n", i, i, i, i)
	}
}
