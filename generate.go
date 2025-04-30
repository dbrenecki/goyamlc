package main

import (
	"go/parser"
	"go/token"
)

func generate(typesPath string) error {
	fs := token.NewFileSet()

	_, err := parser.ParseFile(fs, typesPath, nil, parser.AllErrors|parser.ParseComments)
	if err != nil {
		return err
	}
	return nil
}
