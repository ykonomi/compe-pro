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

	var y int
	fmt.Fscan(reader, &y)

	t, s := false, false
	if y%4 == 0 {
		t = true
		s = true
	}

	if y%100 == 0 {
		t = false
		s = true
	}

	if y%400 == 0 {
		t = true
		s = true
	}

	if !s {
		t = false
	}

	if t {
		fmt.Fprintf(writer, "%v\n", "YES")
	} else {
		fmt.Fprintf(writer, "%v\n", "NO")
	}

}
