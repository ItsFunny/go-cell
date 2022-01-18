/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/18 8:29 下午
# @File : parser.go
# @Description :
# @Attention :
*/
package options

import (
	"fmt"
	"reflect"
	"strings"
)

type parseState struct {
	cmdline []string
	i       int
}

func Parse(cmdline []string, ops map[string]Option) (OptMap, error) {
	var (
		args = make([]string, 0, len(cmdline))
	)
	optDefs := ops
	ret := make(OptMap)
	st := &parseState{
		cmdline: cmdline,
		i:       0,
	}
L:
	for !st.done() {
		param := st.peek()
		switch {
		case param == "--":
			args = append(args, st.cmdline[st.i+1:]...)
			break L
		case strings.HasPrefix(param, "--"):
			// long option
			k, v, err := st.parseLongOpt(optDefs)
			if err != nil {
				return nil, err
			}
			kvType, err := getOptType(k, optDefs)
			if err != nil {
				return nil, err // shouldn't happen b/c k,v was parsed from optsDef
			}
			if err := setOpts(kv{Key: k, Value: v}, kvType, ret); err != nil {
				return nil, err
			}
		case strings.HasPrefix(param, "-") && param != "-":
			// short options
			kvs, err := st.parseShortOpts(optDefs)
			if err != nil {
				return nil, err
			}

			for _, kv := range kvs {
				kv.Key = optDefs[kv.Key].Names()[0]

				kvType, err := getOptType(kv.Key, optDefs)
				if err != nil {
					return nil, err // shouldn't happen b/c kvs was parsed from optsDef
				}
				if err := setOpts(kv, kvType, ret); err != nil {
					return nil, err
				}
			}
		}
		st.i++
	}
	return ret, nil
}

func (st *parseState) parseShortOpts(optDefs map[string]Option) ([]kv, error) {
	k, vStr, ok := splitkv(st.cmdline[st.i][1:])
	kvs := make([]kv, 0, len(k))

	if ok {
		// split at = successful
		k, v, err := parseOpt(k, vStr, optDefs)
		if err != nil {
			return nil, err
		}

		kvs = append(kvs, kv{Key: k, Value: v})

	} else {
	LOOP:
		for j := 0; j < len(k); {
			flag := k[j : j+1]
			od, ok := optDefs[flag]

			switch {
			case !ok:
				return nil, fmt.Errorf("unknown option %q", k)

			case od.Type() == Bool:
				// single char flags for bools
				kvs = append(kvs, kv{
					Key:   od.Name(),
					Value: true,
				})
				j++

			case j < len(k)-1:
				// single char flag for non-bools (use the rest of the flag as value)
				rest := k[j+1:]

				k, v, err := parseOpt(flag, rest, optDefs)
				if err != nil {
					return nil, err
				}

				kvs = append(kvs, kv{Key: k, Value: v})
				break LOOP

			case st.i < len(st.cmdline)-1:
				// single char flag for non-bools (use the next word as value)
				st.i++
				k, v, err := parseOpt(flag, st.cmdline[st.i], optDefs)
				if err != nil {
					return nil, err
				}

				kvs = append(kvs, kv{Key: k, Value: v})
				break LOOP

			default:
				return nil, fmt.Errorf("missing argument for option %q", k)
			}
		}
	}

	return kvs, nil
}

type kv struct {
	Key   string
	Value interface{}
}

func setOpts(kv kv, kvType reflect.Kind, opts OptMap) error {
	if kvType == Strings {
		res, _ := opts[kv.Key].([]string)
		opts[kv.Key] = append(res, kv.Value.([]string)...)
	} else if _, exists := opts[kv.Key]; !exists {
		opts[kv.Key] = kv.Value
	} else {
		return fmt.Errorf("multiple values for option %q", kv.Key)
	}
	return nil
}
func getOptType(k string, optDefs map[string]Option) (reflect.Kind, error) {
	if opt, ok := optDefs[k]; ok {
		return opt.Type(), nil
	}
	return reflect.Invalid, fmt.Errorf("unknown option %q", k)
}

func (st *parseState) parseLongOpt(optDefs map[string]Option) (string, interface{}, error) {
	k, v, ok := splitkv(st.peek()[2:])
	if !ok {
		optDef, ok := optDefs[k]
		if !ok {
			return "", nil, fmt.Errorf("unknown option %q", k)
		}
		if optDef.Type() == Bool {
			return k, true, nil
		} else if st.i < len(st.cmdline)-1 {
			st.i++
			v = st.peek()
		} else {
			return "", nil, fmt.Errorf("missing argument for option %q", k)
		}
	}

	k, optval, err := parseOpt(k, v, optDefs)
	return k, optval, err
}
func parseOpt(opt, value string, opts map[string]Option) (string, interface{}, error) {
	optDef, ok := opts[opt]
	if !ok {
		return "", nil, fmt.Errorf("unknown option %q", opt)
	}

	v, err := optDef.Parse(value)
	if err != nil {
		return "", nil, err
	}
	return optDef.Name(), v, nil
}
func splitkv(opt string) (k, v string, ok bool) {
	split := strings.SplitN(opt, "=", 2)
	if len(split) == 2 {
		return split[0], split[1], true
	} else {
		return opt, "", false
	}
}
func getArgDef(i int, argDefs []Argument) *Argument {
	if i < len(argDefs) {
		// get the argument definition (usually just argDefs[i])
		return &argDefs[i]

	} else if len(argDefs) > 0 {
		// but if i > len(argDefs) we use the last argument definition)
		return &argDefs[len(argDefs)-1]
	}

	// only happens if there aren't any definitions
	return nil
}
func (st *parseState) peek() string {
	return st.cmdline[st.i]
}
func (st *parseState) done() bool {
	return st.i >= len(st.cmdline)
}
