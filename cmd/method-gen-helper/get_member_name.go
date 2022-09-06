package methodgenhelper

import "go/ast"

func GetMemberNames(node *ast.StructType) []string {
	retSlice := []string{}
	if node.Fields.List == nil {
		return retSlice
	}

	for _, field := range node.Fields.List {
		if len(field.Names) > 0 {
			retSlice = append(retSlice, field.Names[0].Name)
		} else {
			// embedded
			idx, ok := field.Type.(*ast.IndexExpr)
			if !ok {
				continue
			}
			ident, ok := idx.X.(*ast.Ident)
			if !ok {
				continue
			}
			retSlice = append(retSlice, ident.Name)
		}
	}
	return retSlice
}
