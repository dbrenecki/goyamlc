# AST Structure

AST Objects have the following structure.

```
ast.File
  ast.GenDecl
    ast.CommentGroup
    ast.TypeSpec
      ast.Ident (the struct)
      ast.StructType
        ast.FieldList
          ast.Field
            (Primitive)
            ast.CommentGroup
            []ast.Ident
              ast.Ident
                ast.Object (nil)
            ast.Ident (Name e.g. int)
          ast.Field
            (Obj)
            ast.CommentGroup
            []ast.Ident
              ast.Ident
                ast.Object (Name)
              ast.Ident
                ast.Object
                  ast.TypeSpec
                    ast.Ident
                    ast.StructType
                      ast.FieldList
                        ast.Field
                          ast.CommentGroup
                          []ast.Ident
                          ...etc
```
