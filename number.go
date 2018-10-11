package jsonwat

import (
	"errors"
	"fmt"
	"strconv"
)

// Int handles both int and "integer string" in JSON documents.
type Int int

// UnmarshalJSON implements the json.Unmarshaler interface.
func (i *Int) UnmarshalJSON(data []byte) error {
	var s string

	switch c := data[0]; c {
	case '[': // array
		return errors.New("jsonwat: cannot assign array into Int")
	case '{': // object
		return errors.New("jsonwat: cannot assign object into Int")
	case 'n': // null
		return errors.New("jsonwat: cannot assign null value into Int")
	case 't', 'f': // true, false
		return errors.New("jsonwat: cannot assign boolean values into Int")
	case '"': // string
		var ok bool
		s, ok = unquote(data)
		if !ok {
			return errors.New("jsonwat: invalid integer string")
		}
	default: // number
		s = string(data)
	}

	j, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return fmt.Errorf("jsonwat: invalid integer string: %v", s)
	}
	*i = Int(j)
	return nil
}
