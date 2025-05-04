package cmd

import (
	"errors"
	"go/ast"

	"github.com/rs/zerolog/log"
)

type astObject struct {
	isRoot bool
	obj    *ast.StructType
}

// FindRootStructs identifies the root structs in the specified
// ast.File.
func FindRootStructs(f *ast.File) ([]string, error) {
	rootStructMap := make(map[string]astObject)

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
		if v.isRoot {
			rootStructs = append(rootStructs, k)
		}
	}

	if len(rootStructs) == 0 {
		return nil, errors.New("unable to determine root struct")
	}
	return rootStructs, nil
}

// getIdent casts and returns the matching ast Type.
func getIdent(field *ast.Field) *ast.Ident {
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

// walk recursively searches the ast tree to find all occurences of root structs
// and stores them in the map.
func walk(rootStructMap map[string]astObject, typeSpec *ast.TypeSpec) {
	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		log.Debug().Msgf(`"%s" *ast.TypeSpec.Type: %T is not of type "*ast.StructType" skipping`, typeSpec.Name.Name, typeSpec.Type)
		return
	}

	// determining which is the root struct
	_, ok = rootStructMap[typeSpec.Name.Name]
	if !ok {
		rootStructMap[typeSpec.Name.Name] = astObject{
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
		identType := getIdent(field)
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
