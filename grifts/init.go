package grifts

import (
	"github.com/danto9/moneybag/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
