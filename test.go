package main

import (
	"flag"

	mruby "webimizer.dev/go_mruby"
)

func main() {
	exefile := flag.String("mrb-file", "", "mrb file path")
	flag.Parse()
	mruby.Mruby_open()
	defer mruby.Mruby_close()
	mruby.Mruby_load_irep_file(*exefile)
}
