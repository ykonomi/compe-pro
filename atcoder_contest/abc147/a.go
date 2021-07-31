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

	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)

	if a+b+c >= 22 {
		fmt.Fprintf(writer, "%v\n", "bust")
	} else {
		fmt.Fprintf(writer, "%v\n", "win")
	}

}
