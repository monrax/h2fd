//go:build !js

package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter raw bytes: ")
	in := bufio.NewReader(os.Stdin)

	line, err := in.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	x := make([]byte, 0)
	for i := 0; i < len(line)-2; {
		if (line[i] != 32) && (line[i] != 10) {
			x, err = hex.AppendDecode(x, line[i:i+2])
			if err != nil {
				fmt.Println(err)
				return
			}
			i += 2
		} else {
			i++
		}
	}

	fmt.Printf("Raw bytes read: [% x]\n\n", x)

	for i, f := range getFrames(x) {
		fmt.Println("Frame at index:", i)
		fmt.Println(f)
	}
}
