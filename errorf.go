package errorf

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"unicode"

	"github.com/Konstantin8105/errors"
)

// Test binary expression in file with filename and return error if
// first letter is not lower.
//
//	X: *ast.CallExpr {
//	.  Fun: *ast.SelectorExpr {
//	.  .  X: *ast.Ident {
//	.  .  .  Name: "fmt"
//	.  .  }
//	.  .  Sel: *ast.Ident {
//	.  .  .  Name: "Error"
//	.  .  }
//	.  }
//	.  Args: []ast.Expr (len = 1) {
//	.  .  0: *ast.BasicLit {
//	.  .  .  Kind: STRING
//	.  .  .  Value: "\"Hello, Golang\\n\""
//	.  .  }
//	.  }
//	}
func Test(filename string) error {
	// positions are relative to fset
	fset := token.NewFileSet()

	// parse file
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	// create new tree error
	et := errors.New(filename)

	funcname := []string{"Errorf"}

	// inspect the AST and print all identifiers and literals.
	ast.Inspect(f, func(n ast.Node) bool {
		// call
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		// selector
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		if id, ok := sel.X.(*ast.Ident); !ok || id.Name != "fmt" {
			return true
		}
		s := sel.Sel
		found := false
		for _, f := range funcname {
			if f == s.Name {
				found = true
			}
		}
		if !found {
			return true
		}

		// argument
		if 0 == len(call.Args) {
			return true
		}
		bl, ok := call.Args[0].(*ast.BasicLit)
		if !ok || bl.Kind != token.STRING {
			return true
		}
		str := []rune(bl.Value)
		if len(str) < 2 {
			return true
		}
		if unicode.IsUpper(str[1]) {
			et.Add(fmt.Errorf("%s:\tnot acceprable first letter: %s\n",
				fset.Position(n.Pos()), bl.Value))
		}
		return true
	})

	// error handling
	if et.IsError() {
		return et
	}
	return nil
}
