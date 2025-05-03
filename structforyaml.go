package main

import (
	"fmt"
	"go/ast"

	"github.com/rs/zerolog/log"
)

type FieldInfo struct {
	description string                `yaml:"description"`
	properties  map[string]*FieldInfo `yaml:"properties,omitempty"`
	rawType     string                `yaml:"raw_type"`
	yamlType    string                `yaml:"yaml_type"`
}

// func createStructsForYaml(rootStructs []string, f *ast.File) {
// 	for _, name := range rootStructs {
// 		createStructForYaml(name, f)
// 	}
// }

// func createStructForYaml(name string, f *ast.File) {
// 	fieldInfos := make(map[string]*FieldInfo)
// 	for _, decl := range f.Decls {
// 		var genDecl *ast.GenDecl
// 		genDecl, ok := decl.(*ast.GenDecl)
// 		if !ok {
// 			log.Debug().Msgf(`ast.Decl: "%T" is not of "*ast.Decl" type, skipping`, decl)
// 			_ = ok
// 			continue
// 		}
// 		for _, s := range genDecl.Specs {
// 			typeSpec, ok := s.(*ast.TypeSpec)
// 			if !ok {
// 				log.Debug().Msgf(`ast.Spec: "%T" is not of "*ast.TypeSpec" type, skipping`, decl)
// 				continue
// 			}
// 			if typeSpec.Name.Name != name {
// 				continue
// 			}
// 			fieldInfos[name] = &FieldInfo{}
// 			constructFieldInfo(fieldInfos[name], typeSpec)
// 		}
// 	}

// 	out, err := yaml.Marshal(fieldInfos)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if err := os.WriteFile("test/data/example.yaml", out, 0644); err != nil {
// 		panic(err)
// 	}
// }

func constructFieldInfo(fieldInfo *FieldInfo, typeSpec *ast.TypeSpec) {
	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		log.Debug().Msgf(`"%s" *ast.TypeSpec.Type: %T is not of type "*ast.StructType" skipping`, typeSpec.Name.Name, typeSpec.Type)
		return
	}

	// TODO fill in kingdoms fields
	walk2(structType, fieldInfo)
}

// walk2 converts a type to a readable yaml type for simple representation
func walk2(structType *ast.StructType, fieldInfo *FieldInfo) {
	fieldInfo.properties = make(map[string]*FieldInfo)

	for _, field := range structType.Fields.List {
		fieldInfo.properties[field.Names[0].Name] = &FieldInfo{}
		fieldInfo.properties[field.Names[0].Name].description = field.Doc.Text()

		switch t := field.Type.(type) {
		case *ast.ArrayType:
			ident, ok := t.Elt.(*ast.Ident)
			if !ok {
				log.Error().Msg("error casting ArrayType")
			}

			fieldInfo.properties[field.Names[0].Name].rawType = fmt.Sprintf("[]%s", ident.Name)

			if ident.Obj == nil {
				fieldInfo.properties[field.Names[0].Name].yamlType = fmt.Sprintf("%s array", ident.Name)
				continue
			}
			fieldInfo.properties[field.Names[0].Name].yamlType = "object array"

			typeSpec, ok := ident.Obj.Decl.(*ast.TypeSpec)
			if !ok {
				fmt.Println("cast went wrong")
			}

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				fmt.Println("cast went wrong")
			}

			walk2(structType, fieldInfo.properties[field.Names[0].Name])

		case *ast.Ident:
			fieldInfo.properties[field.Names[0].Name].rawType = t.Name
			if t.Obj == nil {
				fieldInfo.properties[field.Names[0].Name].yamlType = t.Name
				continue
			}
			fieldInfo.properties[field.Names[0].Name].yamlType = "object"
			typeSpec, ok := t.Obj.Decl.(*ast.TypeSpec)
			if !ok {
				fmt.Println("cast went wrong")
			}

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				fmt.Println("cast went wrong")
			}

			walk2(structType, fieldInfo.properties[field.Names[0].Name])
		case *ast.MapType:
			keyIdent, ok := t.Key.(*ast.Ident)
			if !ok {
				panic("error in casting key *ast.MapType")
			}
			valueIdent, ok := t.Key.(*ast.Ident)
			if !ok {
				panic("error in casting value *ast.MapType")
			}

			fieldInfo.properties[field.Names[0].Name].rawType = fmt.Sprintf("map[%s]%s", keyIdent.Name, valueIdent.Name)
			if valueIdent.Obj == nil {
				fieldInfo.properties[field.Names[0].Name].yamlType = fmt.Sprintf("map[%s]%s", keyIdent.Name, valueIdent.Name)
				continue
			}
			fieldInfo.properties[field.Names[0].Name].yamlType = fmt.Sprintf("map[%s]object", keyIdent.Name)
			typeSpec, ok := keyIdent.Obj.Decl.(*ast.TypeSpec)
			if !ok {
				fmt.Println("ASD")
			}
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				fmt.Println("cast went wrong")
			}
			walk2(structType, fieldInfo.properties[field.Names[0].Name])
		case *ast.StarExpr:
			ident, ok := t.X.(*ast.Ident)
			if !ok {
				log.Error().Msgf("error in casting *ast.StarExpr")
			}

			fieldInfo.properties[field.Names[0].Name].rawType = ident.Name

			if ident.Obj == nil {
				fieldInfo.properties[field.Names[0].Name].yamlType = ident.Name
				continue
			}
			fieldInfo.properties[field.Names[0].Name].yamlType = "object"

			typeSpec, ok := ident.Obj.Decl.(*ast.TypeSpec)
			if !ok {
				fmt.Println("ASD")
			}
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				fmt.Println("cast went wrong")
			}

			walk2(structType, fieldInfo.properties[field.Names[0].Name])

		default:
			fmt.Println("something broke")
		}
	}
}
