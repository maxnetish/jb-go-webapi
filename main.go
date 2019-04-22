package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello Go!")
	fmt.Println("Arguments:", os.Args)
	// fmt.Println("Environments:", os.Environ())
}
