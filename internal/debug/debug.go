package debug

import (
	"os"

	"github.com/davecgh/go-spew/spew"
)

var Active bool
var Trace bool

func start() {
	Active = os.Getenv("debug") != ""
	if Active {
	}
}

func Dump(n ...interface{}) {
	if Active {
		spew.Dump(n)
	}
}