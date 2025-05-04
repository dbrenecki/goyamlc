package cmd

import (
	"fmt"
	"go/ast"
	"strings"
)

type arrayFormatter interface {
	format(typeName string, indentCount int) string
}

type objectArrayFormatter struct{}
type primitiveArrayFormatter struct{}

// type objectFormatter struct{}
// type primitiveFormatter struct{}

func (o objectArrayFormatter) format(typeName string, indentCount int) string {
	return strings.Repeat(" ", indentCount+2) + "- " + typeName
}

// func (o objectFormatter) format(typeName string, indentCount int) string {
// 	return strings.Repeat(" ", indentCount) + "- " + typeName
// }

func (p primitiveArrayFormatter) format(typeName string, indentCount int) string {
	return strings.Repeat(" ", indentCount+2) + "- " + typeName
}

// func (p primitiveFormatter) format(typeName string, indentCount int) string {
// 	// return strings.Repeat(" ", indentCount+2) + "- " + typeName
// }

func (w formatWriter) formatCamelCase(s string) string {
	camelWindow := 1
	indentCount := w.indentCount
	if len(s) > 0 && s[0] == '-' {
		camelWindow = 3
		indentCount += 2
	}
	return strings.Repeat(" ", indentCount) + strings.ToLower(s[:camelWindow]+s[camelWindow:])
}

// writeComment converts a golang comment to a yaml comment and writes to the buffer.
func (w formatWriter) writeComment(astCmt *ast.Comment) error {
	golangCmt := "//"
	if !strings.HasPrefix(astCmt.Text, golangCmt) {
		return fmt.Errorf("ast comment does not start with %#v", golangCmt)
	}
	cmt := strings.Replace(astCmt.Text, golangCmt, "#", 1)
	_, err := w.WriteString(strings.Repeat(" ", w.indentCount) + cmt + "\n")
	if err != nil {
		return err
	}
	return nil
}

func (w formatWriter) writeField(name string, typeName string, isObj bool, isArr *bool) error {
	var a arrayFormatter

	// TODO: this is a hacky incomplete way to detect primitive types and needs fixing.
	switch {
	case !isObj && *isArr:
		a = primitiveArrayFormatter{}
	// case !isObj && !*isArr:
	// 	a = primitiveFormatter{}
	case isObj && *isArr:
		a = objectArrayFormatter{}
		// case isObj && !*isArr:
		// 	a = objectFormatter{}
	}
	if *isArr && !isObj {
		name = w.formatCamelCase(name) + ":\n" + a.format(typeName, w.indentCount) + "\n"
	} else if !*isArr && isObj {
		name = w.formatCamelCase(name) + ": " + "\n"

	} else {
		name = w.formatCamelCase(name) + ": " + typeName + "\n"
	}

	_, err := w.WriteString(name)
	if err != nil {
		return err
	}
	return nil
}
