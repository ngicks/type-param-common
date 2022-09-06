package methodgenhelper

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"path/filepath"
	"strings"
)

type TypeInfo struct {
	ReceiverName        string
	TypeName            string
	TypeParams          string
	InnerMemberName     string
	LennerInterfaceName string
}

// ParseTypeInfo traverse through ast of a given file and retrieves TypeInfo that matches to targetTyName.
func ParseTypeInfo(
	fset *token.FileSet,
	direntName string,
	file io.Reader,
	targetTypeName string,
) (typeInfos []TypeInfo, packageName string, err error) {
	f, err := parser.ParseFile(fset, filepath.Base(direntName), file, 0)
	if err != nil {
		panic(err)
	}

	packageName = f.Name.Name

	ast.Inspect(f, func(n ast.Node) (goAhead bool) {
		// No need to stop stepping deeper.
		goAhead = true

		var ok bool
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok {
			return
		}
		structTy, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return
		}
		if structTy.Fields.List == nil {
			return
		}

		for _, field := range structTy.Fields.List {
			idx, ok := field.Type.(*ast.IndexExpr)
			if !ok {
				continue
			}
			ident, ok := idx.X.(*ast.Ident)
			if !ok {
				continue
			}

			if ident.Name == targetTypeName {
				var memberName string
				if len(field.Names) > 0 {
					memberName = field.Names[0].Name
				} else {
					// embedded
					memberName = targetTypeName
				}

				t := TypeInfo{
					ReceiverName:        "iter",
					TypeName:            typeSpec.Name.Name,
					TypeParams:          getTypeParam(typeSpec),
					InnerMemberName:     memberName,
					LennerInterfaceName: "Lenner",
				}
				typeInfos = append(typeInfos, t)
			}
		}
		return
	})
	return
}

// getTypeParam gets type parameters of typeSpec.
func getTypeParam(typeSpec *ast.TypeSpec) string {
	if typeSpec.TypeParams == nil || typeSpec.TypeParams.List == nil {
		return ""
	}

	tyParams := make([]string, 0)
	for _, v := range typeSpec.TypeParams.List {
		if v.Names != nil {
			for _, n := range v.Names {
				tyParams = append(tyParams, n.Name)
			}
		}
	}
	return strings.Join(tyParams, ",")
}
