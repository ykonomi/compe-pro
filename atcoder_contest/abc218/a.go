package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	switch s[n-1] {
	case "o":
		fmt.Fprintf(writer, "%v\n", "Yes")
	case "x":
		fmt.Fprintf(writer, "%v\n", "No")
	}
}
