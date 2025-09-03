package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Print("Number: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Input Must Be Integer")
		return
	}

	// Convert Input Tobe Array String For Print Output
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = strconv.Itoa(lazyCharter(i))
	}

	fmt.Println("Output:", strings.Join(a, "-"))
}

// Calculate Input With A00124 Formula
func lazyCharter(n int) int {
	return (n*n + n + 2) / 2
}
