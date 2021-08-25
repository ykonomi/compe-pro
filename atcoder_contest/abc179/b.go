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

	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		fmt.Fscan(reader, &b[i])
	}

	cnt := 0
	for i := 0; i < n; i++ {
		if a[i] == b[i] {
			cnt++
		} else {
			cnt = 0
		}
		if cnt == 3 {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "No")
}
