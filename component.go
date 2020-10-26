package goweb

import (
	"errors"
	"strconv"
)

type Attribute struct {
	Type  string `json:"attr-type"`
	Value string `json:"attr-value"`
}

func (a Attribute) String() string {
	return "value: " + a.Value + "\ttype: " + a.Type
}

type Component struct {
	Name string               `json:"comp-name"`
	Vals map[string]Attribute `json:"comp-values"`
}

func (c *Component) Mapped() (map[string]interface{}, error) {
	x := map[string]interface{}{}

	for key, val := range c.Vals {
		switch val.Type {
		case "number":
			f, err := strconv.ParseFloat(val.Value, 64)
			if err != nil {
				return nil, errors.New("error parsing component number value: " + err.Error())
			}

			x[key] = f
		case "bool":
			b, err := strconv.ParseBool(val.Value)
			if err != nil {
				return nil, errors.New("error parsing component boolean value: " + err.Error())
			}

			x[key] = b
		default:
			x[key] = val.Value
		}
	}

	return x, nil
}

func (c Component) String() string {
	s := "name: " + c.Name

	for i, v := range c.Vals {
		s = s + "\n\t\t" + i + ":\n\t\t" + v.String()
	}

	return s
}

func CreateComponentFromStringInterfaceMap(m map[string]interface{}) *Component {
	r := &Component{}

	// for k, v := range m {

	// }

	return r
}
