package main

import (
	"github.com/michalzoldak97/simple_concurrent_filepro/five"
	"github.com/michalzoldak97/simple_concurrent_filepro/four"
	"github.com/michalzoldak97/simple_concurrent_filepro/one"
	"github.com/michalzoldak97/simple_concurrent_filepro/three"
	"github.com/michalzoldak97/simple_concurrent_filepro/two"
)

func main() {
	one.Run()
	two.Run()
	three.Run()
	four.Run()
	five.Run()
}
