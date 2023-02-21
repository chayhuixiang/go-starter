package main

import (
	"fmt"
)

func main() {

	strings := []string{"One", "Two"}
	for _, item := range strings {
		fmt.Println(item)
	}
}
