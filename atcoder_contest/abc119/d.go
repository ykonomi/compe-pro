package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func scanInt() int       { sc.Scan(); return parseInt(sc.Bytes()) }
func scanString() string { sc.Scan(); return string(sc.Bytes()) }
func scanFloat() float64 { sc.Scan(); return parseFloat(sc.Bytes()) }
func scanInts(n int) (ints []int) {
	ints = make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = scanInt()
	}
	return
}
func scanStrings(n int) (st []string) {
	st = make([]string, n)
	for i := 0; i < n; i++ {
		st[i] = scanString()
	}
	return
}
func scanFloats(n int) (fs []float64) {
	fs = make([]float64, n)
	for i := 0; i < n; i++ {
		fs[i] = scanFloat()
	}
	return fs
}

func parseInt(b []byte) (n int) {
	for _, ch := range b {
		ch -= '0'
		if ch <= 9 {
			n = n*10 + int(ch)
		}
	}
	if b[0] == '-' {
		return -n
	}
	return
}

var float64pow10 = []float64{
	1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6,
	1e7, 1e8, 1e9, 1e10, 1e11, 1e12,
	1e13, 1e14, 1e15, 1e16, 1e17, 1e18,
	1e19, 1e20, 1e21, 1e22,
}

func parseFloat(b []byte) float64 {
	f, dot := 0.0, 0
	for i, ch := range b {
		if ch == '.' {
			dot = i + 1
			continue
		}
		if ch -= '0'; ch <= 9 {
			f = f*10 + float64(ch)
		}
	}
	if dot != 0 {
		f /= float64pow10[len(b)-dot]
	}
	if b[0] == '-' {
		return -f
	}
	return f
}

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func chmin(x *int, y int) {
	*x = min(*x, y)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1001001)
	a := scanInt()
	b := scanInt()
	q := scanInt()

	s := scanInts(a)
	t := scanInts(b)
	x := scanInts(q)

	inf := 1 << 60
	s = append([]int{-1 * inf}, s...)
	s = append(s, inf)
	t = append([]int{-1 * inf}, t...)
	t = append(t, inf)

	for i := range x {
		ans := 1 << 60
		j := sort.Search(len(s), func(j int) bool { return x[i] <= s[j] })
		k := sort.Search(len(t), func(j int) bool { return x[i] <= t[j] })

		sl, sr := abs(x[i]-s[j-1]), abs(x[i]-s[j])
		tl, tr := abs(x[i]-t[k-1]), abs(x[i]-t[k])

		chmin(&ans, max(sl, tl))
		chmin(&ans, max(sr, tr))

		sd := min(sl, sr)
		td := min(tl, tr)
		chmin(&ans, min(sd, td)*2+max(sd, td))
		fmt.Fprintf(wr, "%v\n", ans)
	}
}
