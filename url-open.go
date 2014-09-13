// Webページを開いてみる
package main

import (
	"flag"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	flag.Parse()
	var s string = "http://www.ikyu.com/"
	if flag.NArg() > 0 {
		s = flag.Arg(0)
	}

	open.Run(s)
}
