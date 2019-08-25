package sortslice

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/types/typeutil"
)

// Analyzer performs analysis of calls to sort.Slice,
// reporting an error if the argument is not a slice.
var Analyzer = &analysis.Analyzer{
	Name:     "sortslice",
	Doc:      "Check arg type of sort.Slice",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspectNode := func(n ast.Node) {
		call := n.(*ast.CallExpr)
		fn, _ := typeutil.Callee(pass.TypesInfo, call).(*types.Func)
		if fn == nil {
			return
		}

		if fn.FullName() != "sort.Slice" {
			return
		}

		typ := pass.TypesInfo.Types[call.Args[0]].Type

		if _, ok := typ.(*types.Slice); ok {
			return
		}
		if _, ok := typ.Underlying().(*types.Slice); ok {
			return
		}
		if _, ok := typ.(*types.Interface); ok {
			return
		}
		if _, ok := typ.Underlying().(*types.Interface); ok {
			return
		}
		pass.Reportf(call.Pos(), "sort.Slice's argument must be a slice; is called with %s", typeName(typ))
	}
	inspect.Preorder(
		[]ast.Node{(*ast.CallExpr)(nil)},
		inspectNode)
	return nil, nil
}

func typeName(t types.Type) string {
	switch t := t.(type) {
	case *types.Named:
		return t.Obj().Pkg().Name() + "." + t.Obj().Name()
	case *types.Pointer:
		return "*" + typeName(t.Elem())
	}
	return fmt.Sprint(t)
}
