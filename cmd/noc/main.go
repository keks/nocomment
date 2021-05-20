// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"io"
	"os"

	"go.cryptoscope.co/nocomment"
)

func die(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func check(err error, f func(error)) {
	if err != nil {
		f(err)
	}
}

func main() {
	r := nocomment.NewReader(os.Stdin)
	_, err := io.Copy(os.Stdout, r)
	check(err, die)
}
