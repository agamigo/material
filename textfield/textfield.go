// textfield implements a material textfield component.
//
// See: https://material.io/components/web/catalog/input-controls/text-field/
package textfield // import "agamigo.io/material/textfield"

import (
	"agamigo.io/gojs"
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherjs/js"
)

// TF is a material textfield component.
type TF struct {
	mdc *base.Component

	// The current value of the textfield. Changing this will update the
	// textfield’s value.
	Value string `js:"value"`

	// Whether or not the textfield is disabled.
	Disabled bool `js:"disabled"`

	// Valid and Required are updated according to HTML5 validation markup.
	Valid    bool `js:"valid"`
	Required bool `js:"required"`

	// HelperText provides supplemental information and/or validation
	// messages to users. It appears on input field focus and disappears on
	// input field blur by default, or it can be persistent.
	HelperText string `js:"helperText"`
}

// New returns a new component.
func New() *TF {
	c := &TF{}
	c.Component()
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *TF) Start(rootElem *js.Object) error {
	return base.Start(c, rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *TF) Stop() error {
	return base.Stop(c)
}

// Component returns the component's underlying base.Component.
func (c *TF) Component() *base.Component {
	switch {
	case c.mdc == nil:
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCTextField",
				MDCCamelCaseName: "textField",
			},
		}
		fallthrough
	case c.mdc.Object == nil:
		c.mdc.Component().SetState(c.StateMap())
	}
	return c.mdc.Component()
}

// StateMap implements the base.StateMapper interface.
func (c *TF) StateMap() base.StateMap {
	sm := base.StateMap{
		"value":      c.Value,
		"disabled":   c.Disabled,
		"valid":      c.Value,
		"required":   c.Required,
		"helperText": c.HelperText,
	}
	if c.Component().Object.Get("value").String() == "undefined" {
		sm["value"] = js.InternalObject(c).Get("Value")
	}
	if c.Component().Object.Get("helperText").String() == "undefined" {
		sm["helperText"] = js.InternalObject(c).Get("HelperText")
	}
	return sm
}

// Layout adjusts the dimensions and positions for all sub-elements.
func (tf *TF) Layout() error {
	var err error
	gojs.CatchException(&err)
	tf.mdc.Call("layout")
	return err
}
