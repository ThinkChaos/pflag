package pflag

import (
	"fmt"
	"strconv"
)

// -- bool Value
type boolValue bool

func newBoolValue(val bool, p *bool) *boolValue {
	*p = val
	return (*boolValue)(p)
}

func (b *boolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = boolValue(v)
	return err
}

func (b *boolValue) String() string { return fmt.Sprintf("%v", *b) }

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string) {
	f.VarP(newBoolValue(value, p), name, "", usage)
}

// Like BoolVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BoolVarP(p *bool, name, shorthand string, value bool, usage string) {
	f.VarP(newBoolValue(value, p), name, shorthand, usage)
}

// Like BoolVar, but only accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BoolVarS(p *bool, shorthand string, value bool, usage string) {
	f.VarP(newBoolValue(value, p), "", shorthand, usage)
}

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func BoolVar(p *bool, name string, value bool, usage string) {
	CommandLine.VarP(newBoolValue(value, p), name, "", usage)
}

// Like BoolVar, but accepts a shorthand letter that can be used after a single dash.
func BoolVarP(p *bool, name, shorthand string, value bool, usage string) {
	CommandLine.VarP(newBoolValue(value, p), name, shorthand, usage)
}

// Like BoolVar, but only accepts a shorthand letter that can be used after a single dash.
func BoolVarS(p *bool, shorthand string, value bool, usage string) {
	CommandLine.VarP(newBoolValue(value, p), "", shorthand, usage)
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func (f *FlagSet) Bool(name string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolVarP(p, name, "", value, usage)
	return p
}

// Like Bool, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BoolP(name, shorthand string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolVarP(p, name, shorthand, value, usage)
	return p
}

// Like Bool, but only accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BoolS(shorthand string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolVarP(p, "", shorthand, value, usage)
	return p
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func Bool(name string, value bool, usage string) *bool {
	return CommandLine.BoolP(name, "", value, usage)
}

// Like Bool, but accepts a shorthand letter that can be used after a single dash.
func BoolP(name, shorthand string, value bool, usage string) *bool {
	return CommandLine.BoolP(name, shorthand, value, usage)
}

// Like Bool, but only accepts a shorthand letter that can be used after a single dash.
func BoolS(shorthand string, value bool, usage string) *bool {
	return CommandLine.BoolP("", shorthand, value, usage)
}
