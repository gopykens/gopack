header = `// DON'T EDIT!!! THIS FILE IS GENERATED BY %%qlang gentypes.ql%%
//
package builtin

import (
	"reflect"
)
`

template = `// -----------------------------------------------------------------------------

type ty{{Type}} int

func (p ty{{Type}}) GoType() reflect.Type {

	return goty{{TypeTitle}}
}

// NewInstance creates a new instance of a qlang type. required by %%qlang type%% spec.
//
func (p ty{{Type}}) NewInstance(args ...interface{}) interface{} {

	ret := new({{TypeLower}})
	if len(args) > 0 {
		*ret = {{Type}}(args[0])
	}
	return ret
}

func (p ty{{Type}}) Call(a interface{}) {{TypeLower}} {

	return {{Type}}(a)
}

// Ty{{Type}} represents the %%{{TypeLower}}%% type.
//
var Ty{{Type}} = ty{{Type}}(0)
`

footer = `// -----------------------------------------------------------------------------`

builtins = [
	"float32",
	"float64",
	"int",
	"int8",
	"int16",
	"int32",
	"int64",
	"uint",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"string",
]

f, err = os.create("types_builtin.go")
if err != nil {
	fprintln(os.stderr, err)
	return 1
}
defer f.close()

fprintln(f, strings.replace(header, "%%", "`", -1))

for _, typLower = range builtins {
	typ = strings.title(typLower)
	typTitle = strings.title(typLower)
	replacer = strings.replacer(
		"%%", "`", "{{Type}}", typ, "{{TypeLower}}", typLower, "{{TypeTitle}}", typTitle)
	fprintln(f, replacer.replace(template))
}

fprintln(f, footer)
