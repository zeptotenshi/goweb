package goweb

import (
	"errors"
	"syscall/js"
)

type Element struct {
	id  string
	tag string

	el         js.Value
	components map[string]*Component

	wp *WebPage
}

func (e *Element) SetAttribute(_compName string, vals map[string]interface{}) {
	if len(vals) == 0 {
		e.el.Call("setAttribute", _compName, js.ValueOf(""))
	} else if v, ok := vals["var"]; ok && len(vals) == 1 {
		e.el.Call("setAttribute", _compName, js.ValueOf(v))
	} else {
		e.el.Call("setAttribute", _compName, js.ValueOf(vals))
	}

	e.components[_compName] = CreateComponentFromStringInterfaceMap(vals)
}

func (e *Element) SetAttributes(_comps []Component) error {
	for _, v := range _comps {
		m, err := v.Mapped()
		if err != nil {
			return errors.New("element id: " + e.id + "| error setting '" + v.Name + "' attribute| error generating map: " + err.Error())
		}

		e.SetAttribute(v.Name, m)
	}

	return nil
}

func (e *Element) GetProperty(_names ...string) (js.Value, error) {
	v := js.ValueOf(nil)

	for i, n := range _names {
		if i == 0 {
			v = e.el.Get(n)
			if v.IsUndefined() {
				return js.ValueOf(nil), errors.New(n + " property undefined")
			}
			if v.IsNull() {
				return js.ValueOf(nil), errors.New(n + " property null")
			}
		} else {
			v = v.Get(n)
			if v.IsUndefined() {
				return js.ValueOf(nil), errors.New(n + " property undefined")
			}
			if v.IsNull() {
				return js.ValueOf(nil), errors.New(n + " property null")
			}
		}
	}

	return v, nil
}

func (e *Element) GetAttribute(_name string) (js.Value, error) {
	a := e.el.Call("getAttribute", _name)
	if a.IsUndefined() {
		return js.ValueOf(nil), errors.New(_name + " attribute undefined")
	}
	if a.IsNull() {
		return js.ValueOf(nil), errors.New(_name + " attribute null")
	}

	return a, nil
}

// func (e *Element) GetComponent(_attName string) (Component, error) {
// 	if v, ok := e.components[_attName]; ok {
// 		return *v, nil
// 	}

// 	var bytes []byte
// 	var att map[string]interface{}

// 	js.CopyBytesToGo(&bytes, e.El.Call("getAttribute", _attName))
// 	err = json.Unmarshal(bytes, &att)
// 	if err != nil {
// 		return nil, errors.New("'" + _attName + "' attribute unmarshal error: " + err.Error())
// 	}

// 	c, err := mapToComponent(att)
// 	if err != nil {
// 		return nil, errors.New("error converting attribute '" + _attName + "' map to component: " + err.Error())
// 	}

// 	c.Name = _attName
// 	return *c, nil
// }

func (e *Element) AppendChild(_el Element) {

}

func (e *Element) Tag() string {
	return e.tag
}

func (e *Element) ID() string {
	return e.id
}

func (e *Element) Value() js.Value {
	return e.el
}

func (e *Element) String() string {
	return ""
}
