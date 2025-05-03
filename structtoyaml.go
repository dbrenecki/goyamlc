package main

import (
	"bufio"
	"errors"
	"fmt"
	"go/ast"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

func createStructsForYaml(rootStructs []string, f *ast.File) (err error) {
	file, err := os.Create("test/data/example.yaml")

	defer func() {
		if tempErr := file.Close(); tempErr != nil && err == nil {
			err = tempErr
		}
	}()

	w := bufio.NewWriter(file)

	_, err = w.WriteString("---\n")
	if err != nil {
		return err
	}

	for _, name := range rootStructs {
		if err := walkStructs(name, f, w); err != nil {
			return err
		}
	}

	if err := w.Flush(); err != nil {
		return err
	}
	return nil
}

func walkStructs(name string, f *ast.File, w *bufio.Writer) error {
	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			log.Debug().Msgf(`ast.Decl: "%T" is not of "*ast.Decl" type, skipping`, decl)
			_ = ok
			continue
		}

		if genDecl.Doc != nil {
			for _, c := range genDecl.Doc.List {
				if err := writeComment("", c, w); err != nil {
					return err
				}
			}
		}

		var typeSpec *ast.TypeSpec

		for _, s := range genDecl.Specs {
			typeSpec, ok = s.(*ast.TypeSpec)
			if !ok {
				log.Debug().Msgf(`ast.Spec type "%T" is not a "*ast.TypeSpec", skipping`, decl)
				continue
			}
			if typeSpec.Name.Name != name {
				continue
			}
		}

		if typeSpec == nil || typeSpec.Name.Name != name {
			continue
		}
		// wipe root name after its found.
		// TODO: position of root can probably be returned as part of generate.

		fmt.Println("Writing field", typeSpec.Name.Name)
		_, err := w.WriteString(strings.ToLower(typeSpec.Name.Name[:1]) + typeSpec.Name.Name[1:] + ":\n")
		if err != nil {
			return err
		}

		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			log.Debug().Msgf(`ast.TypeSpec: "%T" is not of "*ast.StrucType" type, skipping`, structType)
			return errors.New("somethign EWNWT WRONG")
		}

		fields := structType.Fields.List
		for _, field := range fields {
			if field.Doc != nil {
				indent := strings.Repeat(" ", 2)
				for _, c := range field.Doc.List {
					if err = writeComment(indent, c, w); err != nil {
						return err
					}
					// cmt := strings.Replace(c.Text, "//", "#", 1)

					// fmt.Println("Writing comment2")
					// _, err := w.WriteString(indent + cmt + "\n")
					// if err != nil {
					// 	return err
					// }
				}
			}

		}
		break

	}

	return nil
	// out, err := yaml.Marshal(specCfg)
	// if err != nil {
	// 	panic(err)
	// }
	// if err := os.WriteFile("test/data/example.yaml", out, 0644); err != nil {
	// 	panic(err)
	// }
}

func writeComment(indent string, astCmt *ast.Comment, w *bufio.Writer) error {
	cmt := strings.Replace(astCmt.Text, "//", "#", 1)

	fmt.Println("Writing comment2")
	_, err := w.WriteString(indent + cmt + "\n")
	if err != nil {
		return err
	}
	return nil
}
