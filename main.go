package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ikawaha/jwtdec/decoder"
)

func main() {
	var r io.Reader
	r = os.Stdin
	if len(os.Args) == 2 {
		r = strings.NewReader(os.Args[1])
	} else if len(os.Args) > 2 {
		_, _ = fmt.Fprintf(os.Stderr, "usage: %s [jwt]\n", os.Args[0])
		os.Exit(1)
	}
	jwt, err := decoder.Decode(r)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(jwt.Header)
	fmt.Println(jwt.Payload)
	os.Exit(0)
}
