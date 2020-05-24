/*
 Copyright 2020 Qiniu Cloud (七牛云)

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package parser

import (
	"fmt"
	"go/token"
	"reflect"
	"testing"

	"github.com/qiniu/qlang/v6/ast"
	"github.com/qiniu/qlang/v6/ast/asttest"
)

// -----------------------------------------------------------------------------

var fsTestStd = asttest.NewSingleFileFS("/foo", "bar.ql", `package bar; import "io"
	// comment
	x := 0
	if t := false; t {
		x = 3
	} else {
		x = 5
	}
	println("x:", x)

	// comment 1
	// comment 2
	x = 0
	switch s := "Hello"; s {
	default:
		x = 7
	case "world", "hi":
		x = 5
	case "xsw":
		x = 3
	}
	println("x:", x)

	c := make(chan bool, 100)
	select {
	case c <- true:
	case v := <-c:
	default:
		panic("error")
	}
`)

func TestStd(t *testing.T) {
	fset := token.NewFileSet()
	pkgs, err := ParseFSDir(fset, fsTestStd, "/foo", nil, ParseComments)
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

var fsTestStd2 = asttest.NewSingleFileFS("/foo", "bar.ql", `package bar; import "io"
	x := []float64{1, 3.4, 5}
	y := map[string]float64{"Hello": 1, "xsw": 3.4}
	println("x:", x, "y:", y)

	a := [...]float64{1, 3.4, 5}
	b := [...]float64{1, 3: 3.4, 5}
	c := []float64{2: 1.2, 3, 6: 4.5}
	println("a:", a, "b:", b, "c:", c)
`)

func TestStd2(t *testing.T) {
	fset := token.NewFileSet()
	pkgs, err := ParseFSDir(fset, fsTestStd2, "/foo", nil, 0)
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

var fsTestStdFor = asttest.NewSingleFileFS("/foo", "bar.ql", `package bar; import "io"
	n := 0
	for range [1, 3, 5, 7, 11] {
		n++
	}
	println("n:", n)

	sum := 0
	for _, x := range [1, 3, 5, 7, 11] {
		if x > 3 {
			sum += x
		}
	}
	println("sum(1,3,5,7,11):", sum)

	sum = 0
	for i := 1; i < 100; i++ {
		sum += i
	}
	println("sum(1-100):", sum)
`)

func TestStdFor(t *testing.T) {
	fset := token.NewFileSet()
	pkgs, err := ParseFSDir(fset, fsTestStdFor, "/foo", nil, 0)
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
