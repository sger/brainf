package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCompilerInput(t *testing.T) {
	Convey("Compiler Input", t, func() {
		Convey("1 should print 1", func() {
			program := ",."
			input, output := compile(program)
			input <- 1
			result := <-output
			So(result, ShouldEqual, 1)
		})
	})
}
