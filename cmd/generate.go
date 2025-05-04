package cmd

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/rs/zerolog/log"
)

func Generate(typesPath string, testHook func(fs *token.FileSet, f *ast.File) error) error {
	fs := token.NewFileSet()

	f, err := parser.ParseFile(fs, typesPath, nil, parser.AllErrors|parser.ParseComments)
	if err != nil {
		return err
	}

	if testHook != nil {
		if err = testHook(fs, f); err != nil {
			return err
		}
	}

	rootStructs, err := FindRootStructs(f)
	if err != nil {
		return err
	}
	log.Info().Msgf(`root structs are %s`, rootStructs)

	if err = createStructsForYaml(rootStructs, f); err != nil {
		return err
	}
	return nil
}
