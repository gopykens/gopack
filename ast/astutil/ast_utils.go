package astutil

import (
	"go/ast"
	"go/token"
	"reflect"
	"strconv"

	"github.com/qiniu/x/log"
)

// -----------------------------------------------------------------------------

// ToString converts a ast.BasicLit to string value.
func ToString(l *ast.BasicLit) string {
	if l.Kind == token.STRING {
		s, err := strconv.Unquote(l.Value)
		if err == nil {
			return s
		}
	}
	panic("ToString: convert ast.BasicLit to string failed")
}

// -----------------------------------------------------------------------------

// A ConstKind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
type ConstKind = reflect.Kind

const (
	// ConstBoundRune - bound type: rune
	ConstBoundRune = reflect.Int32
	// ConstBoundString - bound type: string
	ConstBoundString = reflect.String
	// ConstUnboundInt - unbound int type
	ConstUnboundInt = ConstKind(reflect.UnsafePointer + 3)
	// ConstUnboundFloat - unbound float type
	ConstUnboundFloat = ConstKind(reflect.UnsafePointer + 4)
	// ConstUnboundComplex - unbound complex type
	ConstUnboundComplex = ConstKind(reflect.UnsafePointer + 5)
)

// IsConstBound checks a const is bound or not.
func IsConstBound(kind ConstKind) bool {
	return kind <= reflect.UnsafePointer
}

// ToConst converts a ast.BasicLit to constant value.
func ToConst(v *ast.BasicLit) (ConstKind, interface{}) {
	switch v.Kind {
	case token.INT:
		n, err := strconv.ParseInt(v.Value, 0, 0)
		if err != nil {
			n2, err2 := strconv.ParseUint(v.Value, 0, 0)
			if err2 != nil {
				log.Fatalln("ToConst: strconv.ParseInt failed:", err2)
			}
			return ConstUnboundInt, n2
		}
		return ConstUnboundInt, n
	case token.CHAR, token.STRING:
		n, err := strconv.Unquote(v.Value)
		if err != nil {
			log.Fatalln("ToConst: strconv.Unquote failed:", err)
		}
		if v.Kind == token.CHAR {
			for _, c := range n {
				return ConstBoundRune, int64(c)
			}
			panic("not here")
		}
		return ConstBoundString, n
	case token.FLOAT:
		n, err := strconv.ParseFloat(v.Value, 64)
		if err != nil {
			log.Fatalln("ToConst: strconv.ParseFloat failed:", err)
		}
		return ConstUnboundFloat, n
	case token.IMAG: // 123.45i
		val := v.Value
		n, err := strconv.ParseFloat(val[:len(val)-1], 64)
		if err != nil {
			log.Fatalln("ToConst: strconv.ParseFloat failed:", err)
		}
		return ConstUnboundComplex, complex(0, n)
	}
	log.Fatalln("ToConst: unknown -", v)
	return 0, nil
}

// ConstBound bounds v into type `t`.
func ConstBound(v interface{}, t reflect.Type) (ret interface{}, ok bool) {
	nkind := t.Kind()
	nv := reflect.New(t).Elem()
	if nkind >= reflect.Int && nkind <= reflect.Int64 {
		switch ov := v.(type) {
		case int64:
			nv.SetInt(ov)
		case uint64:
			nv.SetInt(int64(ov))
		default:
			return nil, false
		}
	} else if nkind >= reflect.Uint && nkind <= reflect.Uintptr {
		switch ov := v.(type) {
		case int64:
			nv.SetUint(uint64(ov))
		case uint64:
			nv.SetUint(ov)
		default:
			return nil, false
		}
	} else if nkind == reflect.Float64 || nkind == reflect.Float32 {
		switch ov := v.(type) {
		case float64:
			nv.SetFloat(ov)
		case int64:
			nv.SetFloat(float64(ov))
		case uint64:
			nv.SetFloat(float64(ov))
		default:
			return nil, false
		}
	} else if nkind == reflect.Complex128 || nkind == reflect.Complex64 {
		switch ov := v.(type) {
		case complex128:
			nv.SetComplex(ov)
		case float64:
			nv.SetComplex(complex(float64(ov), 0))
		case int64:
			nv.SetComplex(complex(float64(ov), 0))
		case uint64:
			nv.SetComplex(complex(float64(ov), 0))
		default:
			return nil, false
		}
	} else {
		return nil, false
	}
	return nv.Interface(), true
}

// -----------------------------------------------------------------------------

// RecvInfo represents recv information of a method.
type RecvInfo struct {
	Name    string
	Type    string
	Pointer int
}

// ToRecv converts a ast.FieldList to recv information.
func ToRecv(recv *ast.FieldList) (ret RecvInfo) {
	fields := recv.List
	if len(fields) != 1 {
		panic("ToRecv: multi recv object?")
	}
	field := fields[0]
	if field.Names != nil {
		ret.Name = field.Names[0].Name
	}
	t := field.Type
retry:
	switch v := t.(type) {
	case *ast.Ident: // T
		ret.Type = v.Name
	case *ast.StarExpr: // *T
		ret.Pointer++
		t = v.X
		goto retry
	default:
		panic("ToRecv: recv can only be *T or T")
	}
	return
}

// -----------------------------------------------------------------------------
