package service

import (
	"fmt"
	"strings"

	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

type TypeMap map[string]map[string]string

func ParseSchema(schemaStr string) (TypeMap, error) {
	doc, err := gqlparser.LoadSchema(&ast.Source{Input: schemaStr})
	if err != nil {
		return nil, err
	}
	types := make(TypeMap)
	for typeName, def := range doc.Types {
		if validateString(typeName) && len(def.Fields) > 0 {
			types[typeName] = make(map[string]string)
			fmt.Println(typeName)
			for _, filed := range def.Fields {
				fmt.Println(filed.Name)
				fmt.Println(filed.Type.String())
				if validateString(filed.Name) {
					types[typeName][filed.Name] = filed.Type.String()
				}
			}
		}

	}
	return types, nil
}

func validateString(str string) bool {
	return !strings.HasPrefix(str, "__")
}
