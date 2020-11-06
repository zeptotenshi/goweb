package goweb

import (
	"errors"
	"syscall/js"
)

type WebPage struct {
	title   string
	version string

	document js.Value
	console  js.Value
}

func NewWebPage() *WebPage {
	doc := js.Global().Get("document")

	return &WebPage{
		title:    doc.Get("title").String(),
		document: doc,
		console:  js.Global().Get("console"),
		version:  "0.0.1",
	}
}

func (wp *WebPage) CreateElementWithTag(_tag string) *Element {
	e := wp.document.Call("createElement", _tag)
	if e.IsUndefined() {
		panic("element: " + _tag + " : creation failed - undefined")
	} else if e.IsNull() {
		panic("element: " + _tag + " : creation failed - null")
	}

	return &Element{
		id:         "",
		tag:        _tag,
		el:         e,
		components: map[string]*Component{},
		wp:         wp,
	}
}

func (wp *WebPage) CreateElementFromValue(_v js.Value) *Element {
	e := &Element{
		el:         _v,
		components: map[string]*Component{},
		wp:         wp,
	}

	id := _v.Get("id")
	if id.IsNull() || id.IsUndefined() {
		e.id = ""
	} else {
		e.id = id.String()
	}

	tag := _v.Get("tagName")
	if tag.IsNull() || tag.IsUndefined() {
		e.tag = ""
	} else {
		e.tag = tag.String()
	}

	return e
}

func (wp *WebPage) GetElementByID(_id string) (*Element, error) {
	v := wp.document.Call("querySelector", "#"+_id)
	if v.IsUndefined() {
		return nil, errors.New(_id + " element undefined")
	} else if v.IsNull() {
		return nil, errors.New(_id + " element null")
	}

	return &Element{
		id:  _id,
		tag: v.Get("tagName").String(),
		el:  v,

		wp: wp,
	}, nil
}

func (wp *WebPage) SetCookie(_name, _val string) {
	wp.document.Set("cookie", _name+"="+_val)
}

// GetElementByTag returns the first element in the document with the given tag (e.g. <div>, <a-entity>, etc.)
func (wp *WebPage) GetElementByTag(_tag string) (*Element, error) {
	v := wp.document.Call("querySelector", _tag)
	if v.IsUndefined() {
		return nil, errors.New("element of type [" + _tag + "] undefined")
	} else if v.IsNull() {
		return nil, errors.New("element of type [" + _tag + "] null")
	}

	el := &Element{
		tag:        _tag,
		el:         v,
		components: map[string]*Component{},
		wp:         wp,
	}

	n := v.Get("id")
	if n.IsUndefined() || n.IsNull() {
		el.id = ""
	} else {
		el.id = n.String()
	}

	return el, nil
}

func (wp *WebPage) RemoveElementByID(_id string) {
	v := wp.document.Call("getElementById", _id)
	if v.IsNull() || v.IsUndefined() {
		return
	}

	v.Call("remove")
}

func (wp *WebPage) LogElement(e *Element) {
	wp.console.Call("log", e.el)
}

func (wp *WebPage) LogError(err error) {
	wp.LogMessage("{" + wp.title + "|ERROR}: " + err.Error())
}

func (wp *WebPage) LogMessage(msg string) {
	wp.console.Call("log", msg)
}

func (wp *WebPage) LogValue(_v js.Value) {
	wp.console.Call("log", _v)
}

func (wp *WebPage) PageLoaded() {
	wp.LogMessage(wp.title + " wasm initialized | ver " + wp.version)
}
