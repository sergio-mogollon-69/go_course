package main

import (
	"fmt"
)


func main() {
	miSlice := [] int {10,20,30,40,50}

	for i := range miSlice {
		fmt.Println(miSlice[i])
	}
}
