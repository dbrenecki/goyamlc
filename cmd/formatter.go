package cmd

import (
	"fmt"
	"go/ast"
	"strings"
)

func (w formatWriter) formatCamelCase(s string) string {
	camelWindow := 1
	indentCount := w.indentCount
	if len(s) > 0 && false {
		fmt.Println("camel str", s)
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
