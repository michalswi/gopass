package main

import (
	"fmt"

	"github.com/michalswi/color"
)

const (
	Version = "v0.1.0"
	banner  = `
┌─┐┌─┐┌─┐┌─┐┌─┐┌─┐
│ ┬│ │├─┘├─┤└─┐└─┐
└─┘└─┘┴  ┴ ┴└─┘└─┘
	` + Version + ` - @michalswi
`
)

func ShowBanner() {
	fmt.Printf("%s", color.Format(color.BLUE, banner))
}
