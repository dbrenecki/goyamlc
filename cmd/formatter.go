package cmd

import (
	"fmt"
	"go/ast"
	"strings"
)

func (w formatWriter) formatCamelCase(s string) string {
	return strings.Repeat(" ", w.indentCount) + strings.ToLower(s[:1]) + s[1:]
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
