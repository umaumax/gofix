package main

import (
	"flag"
)

var (
	filetype string
)

func init() {
	flag.StringVar(&filetype, "filetype", "", "filetype e.g. *:all filetype")
}
