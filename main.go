//go:build !js

package main

import "fmt"

func main() {
        var input string
	fmt.Print("Enter raw bytes: ")
	fmt.Scanln(&input)

	x := []byte(input)
	fmt.Printf("Raw bytes read: [% x]\n\n", x)

	for i, f := range getFrames(x) {
		fmt.Println("Frame at index:", i)
		fmt.Println(f)
	}
}
