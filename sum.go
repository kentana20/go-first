// 足し算をする
package main

import (
	"flag"
	"os"
	"strconv"
)

var omitNewline = flag.Bool("n", false, "don't print final newline")

func main() {
	flag.Parse()
	sum := 0
	var in int
	for i := 0; i < flag.NArg(); i++ {
		in, _ = strconv.Atoi(flag.Arg(i))
		sum += in
	}
	var out string = strconv.Itoa(sum)
	os.Stdout.WriteString(out)

}
