package mdctest

//go:generate yarn install

import (
	"agamigo.io/gojs"
	"agamigo.io/gojs/jsdom"
	"github.com/gopherjs/gopherjs/js"
)

const (
	MCW_NODE_MODULE = "material-components-web/dist/material-components-web"
)

// LoadMDCModule is a shortcut to
func LoadMDCModule() (err error) {
	gojs.CatchException(&err)
	js.Global.Set("mdc", js.Global.Call("require", MCW_NODE_MODULE))
	return err
}

// EmulateDOM sets up a fake DOM in Node for "gopherjs test" We emulate a
// browser dom since tests run in Node, and MDC components need a dom element to
// attach to.  This is not needed when running in a browser.
func EmulateDOM() (dom jsdom.JSDOM, err error) {
	dom, err = jsdom.New()
	if err != nil {
		return nil, err
	}
	dom.SetHTML(`<html><body></body></html>`)
	js.Global.Set("window", dom.Window())
	js.Global.Set("HTMLElement", dom.Window().Get("HTMLElement"))
	return dom, err
}
