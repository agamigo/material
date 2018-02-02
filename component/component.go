package component

import (
	"errors"

	"agamigo.io/gojs"
	"github.com/gopherjs/gopherjs/js"
)

type StatusType int

const (
	Uninitialized StatusType = iota
	Stopped
	Running
)

type C interface {
	GetObject() *js.Object
	Start() error
	StartWith(querySelector string) error
	StartWithElement(element *js.Object) error
	Stop() error
	String() string
}

type component struct {
	*js.Object
	name   Type
	status StatusType
}

func New(n Type) (mdcComponent C, err error) {
	defer gojs.CatchException(&err)

	c, err := NewWith(n, js.Global)
	return c, err
}

func NewWith(n Type, dom *js.Object) (mdcComponent C, err error) {
	defer gojs.CatchException(&err)

	c := &component{}
	c.name = n

	o, err := makeMDComponent(c, dom)
	if err != nil {
		return nil, err
	}

	c.setObject(o)
	c.setStatus(Stopped)
	return c, err
}

func (c *component) String() string {
	return "{\"component\":\"" + c.name.componentString() + "\"," +
		"\"status\":\"" + c.status.String() + "\"}"
}

func (c *component) setStatus(s StatusType) {
	c.status = s
}

func (c *component) GetObject() *js.Object {
	return c.Object
}

func (c *component) setObject(o *js.Object) error {
	var err error
	defer gojs.CatchException(&err)
	c.Object = o
	return err
}

func (s StatusType) String() string {
	switch s {
	case Stopped:
		return "stopped"
	case Running:
		return "running"
	}

	return "uninitialized"
}

func makeMDComponent(c *component, dom *js.Object) (*js.Object, error) {
	var err error
	defer gojs.CatchException(&err)

	mdcObject := dom.Get("mdc")

	switch c.name {
	case Animation:
		return mdcObject.Get("animation").Get(c.name.componentString()), err
	case Checkbox:
		return mdcObject.Get("checkbox").Get(c.name.componentString()), err
	case Dialog:
		return mdcObject.Get("dialog").Get(c.name.componentString()), err
	case PermanentDrawer:
		return mdcObject.Get("drawer").Get(c.name.componentString()), err
	case PersistentDrawer:
		return mdcObject.Get("drawer").Get(c.name.componentString()), err
	case SlidableDrawer:
		return mdcObject.Get("drawer").Get(c.name.componentString()), err
	case TemporaryDrawer:
		return mdcObject.Get("drawer").Get(c.name.componentString()), err
	case FormField:
		return mdcObject.Get("formField").Get(c.name.componentString()), err
	case GridList:
		return mdcObject.Get("gridList").Get(c.name.componentString()), err
	case IconToggle:
		return mdcObject.Get("iconToggle").Get(c.name.componentString()), err
	case LinearProgress:
		return mdcObject.Get("linearProgress").Get(c.name.componentString()), err
	case Menu:
		return mdcObject.Get("menu").Get(c.name.componentString()), err
	case Radio:
		return mdcObject.Get("radio").Get(c.name.componentString()), err
	case Ripple:
		return mdcObject.Get("ripple").Get(c.name.componentString()), err
	case Select:
		return mdcObject.Get("select").Get(c.name.componentString()), err
	// case SelectionControl:
	// 	return ""
	case Slider:
		return mdcObject.Get("slider").Get(c.name.componentString()), err
	case Snackbar:
		return mdcObject.Get("snackbar").Get(c.name.componentString()), err
	case Tab:
		return mdcObject.Get("tab").Get(c.name.componentString()), err
	case TabBar:
		return mdcObject.Get("tab").Get(c.name.componentString()), err
	case TabBarScroller:
		return mdcObject.Get("tab").Get(c.name.componentString()), err
	// case Textfield:
	// 	return ""
	case Toolbar:
		return mdcObject.Get("toolbar").Get(c.name.componentString()), err
	}
	return nil, err
}

func (c *component) Start() (err error) {
	switch c.name {
	case Checkbox:
		err = c.StartWith("div.mdc-" + string(c.name.classString()))
	}
	return err
}

func (c *component) StartWith(querySelector string) (err error) {
	defer gojs.CatchException(&err)

	e := js.Global.Get("window").Get("document").Call("querySelector",
		querySelector)

	return c.StartWithElement(e)
}

func (c *component) StartWithElement(e *js.Object) (err error) {
	defer gojs.CatchException(&err)

	if c.status == Running {
		return nil
	}

	if c.status != Stopped {
		return errors.New("Attempted to Start() an uninitialized component: " +
			c.String() + ". Use mdc.New()")
	}

	o := c.GetObject().New(e)
	err = c.setObject(o)
	c.setStatus(Running)

	return err
}

func (c *component) Stop() (err error) {
	defer gojs.CatchException(&err)

	if c.status == Stopped {
		return errors.New("Cannot Stop() already stopped component: " +
			c.String())
	}

	if c.status != Running {
		return errors.New("Cannot Stop() an uninitialized component: " +
			c.String() + ". Use mdc.New()")
	}

	c.GetObject().Call("destroy")

	return err
}