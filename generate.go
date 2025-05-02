package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"github.com/rs/zerolog/log"
)

type AstObject struct {
	isRoot bool
	obj    *ast.StructType
}

func findRootStructs(f *ast.File) ([]string, error) {
	rootStructMap := make(map[string]AstObject)

	for _, decl := range f.Decls {
		var genDecl *ast.GenDecl
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			log.Debug().Msgf(`ast.Decl: "%T" is not of "*ast.GenDecl" type, skipping`, decl)
			_ = ok
			continue
		}
		for _, s := range genDecl.Specs {
			typeSpec, ok := s.(*ast.TypeSpec)
			if !ok {
				log.Debug().Msgf(`ast.Spec: "%T" is not of "*ast.TypeSpec" type, skipping`, decl)
				continue
			}
			walk(rootStructMap, typeSpec)
		}
	}

	var rootStructs []string
	for k, v := range rootStructMap {
		fmt.Println("k", k, "v", v)
		if v.isRoot {
			rootStructs = append(rootStructs, k)
		}
	}

	if len(rootStructs) == 0 {
		return nil, errors.New("unable to determine root struct")
	}
	return rootStructs, nil

}

func generate(typesPath string) error {
	fs := token.NewFileSet()

	f, err := parser.ParseFile(fs, typesPath, nil, parser.AllErrors|parser.ParseComments)
	if err != nil {
		return err
	}

	// TODO: this block only needed for tests
	out, err := os.Create("test/data/types.txt")
	if err != nil {
		return err
	}
	defer out.Close()

	log.Info().Msg("create ast obj tree")

	err = ast.Fprint(out, fs, f, nil)
	if err != nil {
		return err
	}
	//

	rootStructs, err := findRootStructs(f)
	if err != nil {
		return err
	}
	log.Info().Msgf(`root structs are %s`, rootStructs)

	createStructsForYaml(rootStructs, f)
	return nil
}

func getRecursiveIdent(field *ast.Field) *ast.Ident {
	switch t := field.Type.(type) {
	case *ast.ArrayType:
		ident, ok := t.Elt.(*ast.Ident)
		if !ok {
			log.Error().Msg("something went wrong with *ast.ArrayType casting")
			return nil
		}
		return ident
	case *ast.Ident:
		return t
	case *ast.MapType:
		ident, ok := t.Value.(*ast.Ident)
		if !ok {
			log.Error().Msg("something went wrong with *ast.MapType casting")
			return nil
		}
		return ident
	default:
		log.Debug().Msgf("not a recursive type %v", t)
		return nil
	}
}
func walk(rootStructMap map[string]AstObject, typeSpec *ast.TypeSpec) {
	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		log.Debug().Msgf(`"%s" *ast.TypeSpec.Type: %T is not of type "*ast.StructType" skipping`, typeSpec.Name.Name, typeSpec.Type)
		return
	}

	// determining which is the root struct
	_, ok = rootStructMap[typeSpec.Name.Name]
	if !ok {
		rootStructMap[typeSpec.Name.Name] = AstObject{
			isRoot: true,
			obj:    structType,
		}
	} else {
		astStruct := rootStructMap[typeSpec.Name.Name]
		astStruct.isRoot = false
		astStruct.obj = structType
		rootStructMap[typeSpec.Name.Name] = astStruct
	}
	for _, field := range structType.Fields.List {
		identType := getRecursiveIdent(field)
		if identType != nil {
			// recurse when there are nested structs
			if identType.Obj != nil {
				typeSpec2, ok := identType.Obj.Decl.(*ast.TypeSpec)
				if !ok {
					log.Debug().Msgf(`ident Obj Decl is not of type "*ast.TypeSpec", skipping`)
					continue
				}
				walk(rootStructMap, typeSpec2)
			}
		}
	}
}
