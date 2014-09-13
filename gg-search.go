// Webページを開いてみる
package main

import (
	"flag"
	"net/url"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	var w = flag.String("w", "一休", "search word")
	var t = flag.Int("t", 0, "search type 0:normal, 1:image")
	flag.Parse()

	var Url *url.URL
	Url, err := url.Parse("http://www.google.co.jp")
	if err != nil {
		panic("doooon!!!!")
	}

	Url.Path += "/search"
	params := url.Values{}
	params.Add("q", *w)
	if *t == 1 {
		params.Add("tbm", "isch")
	}
	Url.RawQuery = params.Encode()

	open.Run(Url.String())
}
