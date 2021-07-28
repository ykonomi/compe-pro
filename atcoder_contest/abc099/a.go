package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	switch {
	case n < 1000:
		fmt.Fprintf(writer, "%v\n", "ABC")
	default:
		fmt.Fprintf(writer, "%v\n", "ABD")
	}

}
