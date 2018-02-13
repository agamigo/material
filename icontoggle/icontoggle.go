// icontoggle implements a material icontoggle component.
//
// See: https://material.io/components/web/catalog/buttons/icon-toggle-buttons/
package icontoggle // import "agamigo.io/material/icontoggle"

import "agamigo.io/material"

// IT is a material icontoggle component.
type IT struct {
	*material.Component
	On       bool `js:"on"`
	Disabled bool `js:"disabled"`
}

// ComponentType implements the ComponentTyper interface.
func (c *IT) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCIconToggle",
		MDCCamelCaseName: "iconToggle",
	}
}

// SetComponent implements the Componenter interface and replaces the component's
// base Component with mdcC.
func (c *IT) SetComponent(mdcC *material.Component) {
	c.Component = mdcC
}

// String returns the component's "ComponentType: status" information.
func (c *IT) String() string {
	return c.ComponentType().String() + ": " + c.Component.String()
}

// TODO: Wrap refreshToggleData
// TODO: Handle MDCIconToggle:change events
