// go-lib is sourced via zed (into .vendor/.zed, wired in through the go.mod
// replace); everything else would come from the Go module proxy.
package main

import (
	"fmt"
	"os"
	"strings"

	golib "github.com/zed-pkg-test/go-lib"
)

func main() {
	msg := golib.Greet("go-app")
	fmt.Println(msg)
	if !strings.Contains(msg, "from zed-pkg-test/go-lib") {
		fmt.Fprintln(os.Stderr, "FAIL: zed-sourced module did not resolve")
		os.Exit(1)
	}
	fmt.Println("OK: zed-sourced module resolved via go.mod replace")
}
