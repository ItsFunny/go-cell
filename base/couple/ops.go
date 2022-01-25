package couple

import (
	"errors"
	"fmt"
	"github.com/itsfunny/go-cell/base/core/options"
	"reflect"
)

func CheckAndConvertOptions(request IServerRequest, opts []options.Option) (options.OptMap, error) {
	ret := make(options.OptMap)

	var (
		name     string
		required bool
		v        interface{}
	)
	for _, opt := range opts {
		name = opt.Name()
		required = opt.Required()
		p := request.GetParameter(name)
		if len(p) == 0 {
			v = opt.Default()
			// TODO len==0
			if v == nil {
				if !required {
					continue
				}
				return nil, errors.New("missing")
			}
		} else {
			v = p
		}

		kind := reflect.TypeOf(v).Kind()
		if kind != opt.Type() {
			if opt.Type() == options.Strings {
				if _, ok := v.([]string); !ok {
					return ret, fmt.Errorf("option %s should be type %q, but got type %q",
						opt.Name(), opt.Type().String(), kind.String())
				}
			} else {
				str, ok := v.(string)
				if !ok {
					return ret, fmt.Errorf("option %s should be type %q, but got type %q",
						name, opt.Type().String(), kind.String())
				}

				val, err := opt.Parse(str)
				if err != nil {
					value := fmt.Sprintf("value %q", v)
					if len(str) == 0 {
						value = "empty value"
					}
					return ret, fmt.Errorf("could not convert %s to type %q (for option %q)",
						value, opt.Type().String(), "-"+name)
				}
				ret[name] = val
			}
		} else {
			ret[name] = v
		}
	}

	return ret, nil
}
