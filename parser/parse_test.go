package parser

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/qiniu/qlang/v6/ast"
	"github.com/qiniu/qlang/v6/ast/asttest"
	"github.com/qiniu/qlang/v6/token"
	"github.com/qiniu/x/log"
)

func init() {
	log.SetFlags(log.Ldefault &^ log.LstdFlags)
	log.SetOutputLevel(log.Ldebug)
}

// -----------------------------------------------------------------------------

var fsTest1 = asttest.NewSingleFileFS("/foo", "bar.ql", `package bar; import "io"
func New() (*Bar, error) {
	return nil, io.EOF
}

bar, err := New()
if err != nil {
	log.Println(err)
}`)

func TestParseBarPackage(t *testing.T) {
	fset := token.NewFileSet()
	pkgs, err := ParseFSDir(fset, fsTest1, "/foo", nil, 0)
	if err != nil || len(pkgs) != 1 {
		t.Fatal("ParseFSDir failed:", err, len(pkgs))
	}
	bar := pkgs["bar"]
	file := bar.Files["/foo/bar.ql"]
	fmt.Println("Pkg:", file.Name)
	for _, decl := range file.Decls {
		fmt.Println("decl:", reflect.TypeOf(decl))
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch v := spec.(type) {
				case *ast.ImportSpec:
					fmt.Println(" - import:", v.Path.Value)
				}
			}
		case *ast.FuncDecl:
			fmt.Println(" - func:", d.Name.Name)
		}
	}
}

// -----------------------------------------------------------------------------

var fsTest2 = asttest.NewSingleFileFS("/foo", "bar.ql", `import "io"
func New() (*Bar, error) {
	return nil, io.EOF
}

bar, err := New()
if err != nil {
	log.Println(err)
}
`)

func TestParseNoPackageAndGlobalCode(t *testing.T) {
	fset := token.NewFileSet()
	pkgs, err := ParseFSDir(fset, fsTest2, "/foo", nil, 0)
	if err != nil || len(pkgs) != 1 {
		t.Fatal("ParseFSDir failed:", err, len(pkgs))
	}
	bar, isMain := pkgs["main"]
	if !isMain {
		t.Fatal("TestParseNoPackageAndGlobalCode failed: not main")
	}
	file := bar.Files["/foo/bar.ql"]
	fmt.Println("Pkg:", file.Name)
	for _, decl := range file.Decls {
		fmt.Println("decl:", reflect.TypeOf(decl))
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch v := spec.(type) {
				case *ast.ImportSpec:
					fmt.Println(" - import:", v.Path.Value)
				}
			}
		case *ast.FuncDecl:
			fmt.Println(" - func:", d.Name.Name)
		}
	}
}

// -----------------------------------------------------------------------------

var fsTestMapLit = asttest.NewSingleFileFS("/foo", "bar.ql", `
	x := {"Hello": 1, "xsw": 3.4}
	println("x:", x)
`)

func TestMapLit(t *testing.T) {
	fset := token.NewFileSet()
	pkgs, err := ParseFSDir(fset, fsTestMapLit, "/foo", nil, Trace)
	if err != nil || len(pkgs) != 1 {
		t.Fatal("ParseFSDir failed:", err, len(pkgs))
	}
	bar, isMain := pkgs["main"]
	if !isMain {
		t.Fatal("TestMap failed: not main")
	}
	file := bar.Files["/foo/bar.ql"]
	fmt.Println("Pkg:", file.Name)
	for _, decl := range file.Decls {
		fmt.Println("decl:", reflect.TypeOf(decl))
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch v := spec.(type) {
				case *ast.ImportSpec:
					fmt.Println(" - import:", v.Path.Value)
				}
			}
		case *ast.FuncDecl:
			fmt.Println(" - func:", d.Name.Name)
		}
	}
}

// -----------------------------------------------------------------------------

var fsTestSliceLit = asttest.NewSingleFileFS("/foo", "bar.ql", `
	x := [1, 3.4, 5.6]
	println("x:", x)
`)

func TestSliceLit(t *testing.T) {
	fset := token.NewFileSet()
	pkgs, err := ParseFSDir(fset, fsTestSliceLit, "/foo", nil, Trace)
	if err != nil || len(pkgs) != 1 {
		t.Fatal("ParseFSDir failed:", err, len(pkgs))
	}
	bar, isMain := pkgs["main"]
	if !isMain {
		t.Fatal("TestMap failed: not main")
	}
	file := bar.Files["/foo/bar.ql"]
	fmt.Println("Pkg:", file.Name)
	for _, decl := range file.Decls {
		fmt.Println("decl:", reflect.TypeOf(decl))
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch v := spec.(type) {
				case *ast.ImportSpec:
					fmt.Println(" - import:", v.Path.Value)
				}
			}
		case *ast.FuncDecl:
			fmt.Println(" - func:", d.Name.Name)
		}
	}
}

// -----------------------------------------------------------------------------

var fsTestSliceLit2 = asttest.NewSingleFileFS("/foo", "bar.ql", `
	x := [5.6]
	println("x:", x)
`)

func TestSliceLit2(t *testing.T) {
	fset := token.NewFileSet()
	pkgs, err := ParseFSDir(fset, fsTestSliceLit2, "/foo", nil, Trace)
	if err != nil || len(pkgs) != 1 {
		t.Fatal("ParseFSDir failed:", err, len(pkgs))
	}
	bar, isMain := pkgs["main"]
	if !isMain {
		t.Fatal("TestMap failed: not main")
	}
	file := bar.Files["/foo/bar.ql"]
	fmt.Println("Pkg:", file.Name)
	for _, decl := range file.Decls {
		fmt.Println("decl:", reflect.TypeOf(decl))
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch v := spec.(type) {
				case *ast.ImportSpec:
					fmt.Println(" - import:", v.Path.Value)
				}
			}
		case *ast.FuncDecl:
			fmt.Println(" - func:", d.Name.Name)
		}
	}
}

// -----------------------------------------------------------------------------

var fsTestSliceLit3 = asttest.NewSingleFileFS("/foo", "bar.ql", `
	x := []
	println("x:", x)
`)

func TestSliceLit3(t *testing.T) {
	fset := token.NewFileSet()
	pkgs, err := ParseFSDir(fset, fsTestSliceLit3, "/foo", nil, Trace)
	if err != nil || len(pkgs) != 1 {
		t.Fatal("ParseFSDir failed:", err, len(pkgs))
	}
	bar, isMain := pkgs["main"]
	if !isMain {
		t.Fatal("TestMap failed: not main")
	}
	file := bar.Files["/foo/bar.ql"]
	fmt.Println("Pkg:", file.Name)
	for _, decl := range file.Decls {
		fmt.Println("decl:", reflect.TypeOf(decl))
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch v := spec.(type) {
				case *ast.ImportSpec:
					fmt.Println(" - import:", v.Path.Value)
				}
			}
		case *ast.FuncDecl:
			fmt.Println(" - func:", d.Name.Name)
		}
	}
}

// -----------------------------------------------------------------------------

var fsTestListComprehension = asttest.NewSingleFileFS("/foo", "bar.ql", `
	x := [x*x for x <- [1, 2, 3, 4]]
	println("x:", x)
`)

func TestListComprehension(t *testing.T) {
	fset := token.NewFileSet()
	pkgs, err := ParseFSDir(fset, fsTestListComprehension, "/foo", nil, Trace)
	if err != nil || len(pkgs) != 1 {
		t.Fatal("ParseFSDir failed:", err, len(pkgs))
	}
	bar, isMain := pkgs["main"]
	if !isMain {
		t.Fatal("TestMap failed: not main")
	}
	file := bar.Files["/foo/bar.ql"]
	fmt.Println("Pkg:", file.Name)
	for _, decl := range file.Decls {
		fmt.Println("decl:", reflect.TypeOf(decl))
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch v := spec.(type) {
				case *ast.ImportSpec:
					fmt.Println(" - import:", v.Path.Value)
				}
			}
		case *ast.FuncDecl:
			fmt.Println(" - func:", d.Name.Name)
		}
	}
}

// -----------------------------------------------------------------------------

var fsTestMapComprehension = asttest.NewSingleFileFS("/foo", "bar.ql", `
	x := {v: k*k for k, v <- [3, 5, 7, 11]}
	println("x:", x)
`)

func TestMapComprehension(t *testing.T) {
	fset := token.NewFileSet()
	pkgs, err := ParseFSDir(fset, fsTestMapComprehension, "/foo", nil, Trace)
	if err != nil || len(pkgs) != 1 {
		t.Fatal("ParseFSDir failed:", err, len(pkgs))
	}
	bar, isMain := pkgs["main"]
	if !isMain {
		t.Fatal("TestMap failed: not main")
	}
	file := bar.Files["/foo/bar.ql"]
	fmt.Println("Pkg:", file.Name)
	for _, decl := range file.Decls {
		fmt.Println("decl:", reflect.TypeOf(decl))
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch v := spec.(type) {
				case *ast.ImportSpec:
					fmt.Println(" - import:", v.Path.Value)
				}
			}
		case *ast.FuncDecl:
			fmt.Println(" - func:", d.Name.Name)
		}
	}
}

// -----------------------------------------------------------------------------
