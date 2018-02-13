package menu

import (
	"agamigo.io/gojs"
	"github.com/gopherjs/gopherjs/js"
)

// AfterStart implements the material.AfterStarter interface. It adds a missing
// getter to MDCMenu.quickOpen so we can work with that variable as expected in
// Go.
func (c *M) AfterStart() error {
	var err error
	gojs.CatchException(&err)
	o := c.GetObject()

	// Adds a getter for M.quickOpen.
	proto := js.Global.Get("Object").Call("getPrototypeOf", c)
	ogSetter := js.Global.Get("Object").Call("getOwnPropertyDescriptor",
		proto, "quickOpen").Get("set")
	js.Global.Get("Object").Call("defineProperty",
		c, "quickOpen",
		js.M{
			"set": ogSetter,
			"get": func() bool {
				return o.Get("foundation_").Get("quickOpen_").Bool()
			},
		},
	)

	return err
}
