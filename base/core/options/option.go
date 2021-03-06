/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 1:16 下午
# @File : option.go
# @Description :
  默认名称为 第一位,然后最后一位为cmd 的描述符
# @Attention :
*/
package options

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Types of Command options
const (
	Invalid = reflect.Invalid
	Bool    = reflect.Bool
	Int     = reflect.Int
	Uint    = reflect.Uint
	Int64   = reflect.Int64
	Uint64  = reflect.Uint64
	Float   = reflect.Float64
	String  = reflect.String
	Strings = reflect.Array
)

type OptMap map[string]interface{}

// Option is used to specify a field that will be provided by a consumer
type Option interface {
	Name() string    // the main name of the option
	Names() []string // a list of unique names matched with user-provided flags

	Type() reflect.Kind  // value must be this type
	Description() string // a short string that describes this option

	WithDefault(interface{}) Option // sets the default value of the option
	Default() interface{}

	Parse(str string) (interface{}, error)

	Required() bool
	WithRequired(b bool)Option
}

type OptionWrapper struct {
	Option Option
	Value  interface{}
}

type option struct {
	names       []string
	kind        reflect.Kind
	description string
	defaultVal  interface{}
	required    bool
}

func (o *option) WithRequired(b bool) Option{
	o.required=b
	return o
}

func (o *option) Required() bool {
	return o.required
}

func (o *option) Name() string {
	return o.names[0]
}

func (o *option) Names() []string {
	return o.names
}

func (o *option) Type() reflect.Kind {
	return o.kind
}

func (o *option) Description() string {
	if len(o.description) == 0 {
		return ""
	}
	if !strings.HasSuffix(o.description, ".") {
		o.description += "."
	}
	if o.defaultVal != nil {
		if strings.Contains(o.description, "<<default>>") {
			return strings.Replace(o.description, "<<default>>",
				fmt.Sprintf("Default: %v.", o.defaultVal), -1)
		} else {
			return fmt.Sprintf("%s Default: %v.", o.description, o.defaultVal)
		}
	}
	return o.description
}

type converter func(string) (interface{}, error)

var converters = map[reflect.Kind]converter{
	Bool: func(v string) (interface{}, error) {
		if v == "" {
			return true, nil
		}
		v = strings.ToLower(v)

		return strconv.ParseBool(v)
	},
	Int: func(v string) (interface{}, error) {
		val, err := strconv.ParseInt(v, 0, 32)
		if err != nil {
			return nil, err
		}
		return int(val), err
	},
	Uint: func(v string) (interface{}, error) {
		val, err := strconv.ParseUint(v, 0, 32)
		if err != nil {
			return nil, err
		}
		return uint(val), err
	},
	Int64: func(v string) (interface{}, error) {
		val, err := strconv.ParseInt(v, 0, 64)
		if err != nil {
			return nil, err
		}
		return val, err
	},
	Uint64: func(v string) (interface{}, error) {
		val, err := strconv.ParseUint(v, 0, 64)
		if err != nil {
			return nil, err
		}
		return val, err
	},
	Float: func(v string) (interface{}, error) {
		return strconv.ParseFloat(v, 64)
	},
	String: func(v string) (interface{}, error) {
		return v, nil
	},
	Strings: func(v string) (interface{}, error) {
		return v, nil
	},
}

func (o *option) Parse(v string) (interface{}, error) {
	conv, ok := converters[o.Type()]
	if !ok {
		return nil, fmt.Errorf("option %q takes %s arguments, but was passed %q", o.Name(), o.Type(), v)
	}

	return conv(v)
}

// constructor helper functions
func NewOption(kind reflect.Kind, names ...string) Option {
	var desc string

	if len(names) >= 2 {
		desc = names[len(names)-1]
		names = names[:len(names)-1]
	}

	return &option{
		names:       names,
		kind:        kind,
		description: desc,
	}
}

func (o *option) WithDefault(v interface{}) Option {
	if v == nil {
		panic(fmt.Errorf("cannot use nil as a default"))
	}

	// if type of value does not match the option type
	if vKind, oKind := reflect.TypeOf(v).Kind(), o.Type(); vKind != oKind {
		// if the reason they do not match is not because of Slice vs Array equivalence
		// Note: Figuring out if the type of Slice/Array matches is not done in this function
		if !((vKind == reflect.Array || vKind == reflect.Slice) && (oKind == reflect.Array || oKind == reflect.Slice)) {
			panic(fmt.Errorf("invalid default for the given type, expected %s got %s", o.Type(), vKind))
		}
	}
	o.defaultVal = v
	return o
}

func (o *option) Default() interface{} {
	return o.defaultVal
}

// TODO handle description separately. this will take care of the panic case in
// NewOption

// For all func {Type}Option(...string) functions, the last variadic argument
// is treated as the description field.

func BoolOption(names ...string) Option {
	return NewOption(Bool, names...)
}
func IntOption(names ...string) Option {
	return NewOption(Int, names...)
}
func UintOption(names ...string) Option {
	return NewOption(Uint, names...)
}
func Int64Option(names ...string) Option {
	return NewOption(Int64, names...)
}
func Uint64Option(names ...string) Option {
	return NewOption(Uint64, names...)
}
func FloatOption(names ...string) Option {
	return NewOption(Float, names...)
}
func StringOption(names ...string) Option {
	return NewOption(String, names...)
}

// StringsOption is a command option that can handle a slice of strings
func StringsOption(names ...string) Option {
	return &stringsOption{
		Option:    NewOption(Strings, names...),
		delimiter: "",
	}
}

func DelimitedStringsOption(delimiter string, names ...string) Option {
	if delimiter == "" {
		panic("cannot create a DelimitedStringsOption with no delimiter")
	}
	return &stringsOption{
		Option:    NewOption(Strings, names...),
		delimiter: delimiter,
	}
}

type stringsOption struct {
	Option
	delimiter string
}

func (s *stringsOption) WithDefault(v interface{}) Option {
	if v == nil {
		return s.Option.WithDefault(v)
	}

	defVal := v.([]string)
	s.Option = s.Option.WithDefault(defVal)
	return s
}

func (s *stringsOption) Parse(v string) (interface{}, error) {
	if s.delimiter == "" {
		return []string{v}, nil
	}

	return strings.Split(v, s.delimiter), nil
}
