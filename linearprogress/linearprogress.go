// linearprogress implements a material linearprogress component.
//
// See: https://material.io/components/web/catalog/linear-progress/
package linearprogress // import "agamigo.io/material/linearprogress"

import (
	"agamigo.io/gojs"
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherjs/js"
)

// LP is a material libearprogress component.
type LP struct {
	mdc         *base.Component
	Determinate bool    `js:"determinate"`
	Reverse     bool    `js:"reverse"`
	Progress    float64 `js:"progress"`
	Buffer      float64 `js:"buffer"`
	bufferCache float64
}

// New returns a new component.
func New() *LP {
	c := &LP{}
	c.Component()
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *LP) Start(rootElem *js.Object) error {
	backup := c.StateMap()
	backup["buffer"] = c.Buffer
	err := base.Start(c, rootElem)
	if err != nil {
		return err
	}
	err = c.afterStart()
	if err != nil {
		// TODO: handle afterStart + stop error
		c.Stop()
		return err
	}
	c.Component().SetState(backup)
	return nil
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *LP) Stop() error {
	return base.Stop(c)
}

// Component returns the component's underlying base.Component.
func (c *LP) Component() *base.Component {
	switch {
	case c.mdc == nil:
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCLinearProgress",
				MDCCamelCaseName: "linearProgress",
			},
		}
		fallthrough
	case c.mdc.Object == nil:
		c.mdc.Component().SetState(c.StateMap())
	}
	return c.mdc.Component()
}

// StateMap implements the base.StateMapper interface.
func (c *LP) StateMap() base.StateMap {
	sm := base.StateMap{
		"determinate": c.Determinate,
		"reverse":     c.Reverse,
		"progress":    c.Progress,
		"buffer":      c.bufferCache,
	}
	if c.Component().Object.Get("progress").String() == "undefined" {
		sm["progress"] = js.InternalObject(c).Get("Progress")
	}
	return sm
}

// Open opens the linearProgress component.
func (lp *LP) Open() (err error) {
	gojs.CatchException(&err)
	lp.Component().Call("open")
	return err
}

// Close closes the linearProgress component.
func (lp *LP) Close() (err error) {
	gojs.CatchException(&err)
	lp.Component().Call("close")
	return err
}
