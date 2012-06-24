// Permission to use, copy, modify, and/or distribute this software for
// any purpose is hereby granted, provided this notice appear in all copies.

/*
Goembed reads each file in sequence and writes on the standard output
Go source code that wraps the file's contents in a []byte. If no file
is given, goembed reads from the standard input.
*/
package main

import (
	"flag"
	"fmt"
)

func usage() {
	fmt.Fprint(os.Stderr, "usage: goembed [files]\n")
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()
}
