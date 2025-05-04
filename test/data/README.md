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


```


   379  .  .  .  .  .  .  .  .  0: *ast.Field {
   380  .  .  .  .  .  .  .  .  .  Doc: *ast.CommentGroup {
   381  .  .  .  .  .  .  .  .  .  .  List: []*ast.Comment (len = 1) {
   382  .  .  .  .  .  .  .  .  .  .  .  0: *ast.Comment {
   383  .  .  .  .  .  .  .  .  .  .  .  .  Slash: ../test/data/types.go:22:2
   384  .  .  .  .  .  .  .  .  .  .  .  .  Text: "// Some animal type."
   385  .  .  .  .  .  .  .  .  .  .  .  }
   386  .  .  .  .  .  .  .  .  .  .  }
   387  .  .  .  .  .  .  .  .  .  }
   388  .  .  .  .  .  .  .  .  .  Names: []*ast.Ident (len = 1) {
   389  .  .  .  .  .  .  .  .  .  .  0: *ast.Ident {
   390  .  .  .  .  .  .  .  .  .  .  .  NamePos: ../test/data/types.go:23:2
   391  .  .  .  .  .  .  .  .  .  .  .  Name: "Animal"
   392  .  .  .  .  .  .  .  .  .  .  .  Obj: *ast.Object {
   393  .  .  .  .  .  .  .  .  .  .  .  .  Kind: var
   394  .  .  .  .  .  .  .  .  .  .  .  .  Name: "Animal"
   395  .  .  .  .  .  .  .  .  .  .  .  .  Decl: *(obj @ 379)
   396  .  .  .  .  .  .  .  .  .  .  .  .  Data: nil
   397  .  .  .  .  .  .  .  .  .  .  .  .  Type: nil
   398  .  .  .  .  .  .  .  .  .  .  .  }
   399  .  .  .  .  .  .  .  .  .  .  }
   400  .  .  .  .  .  .  .  .  .  }
   401  .  .  .  .  .  .  .  .  .  Type: *ast.Ident {
   402  .  .  .  .  .  .  .  .  .  .  NamePos: ../test/data/types.go:23:9
   403  .  .  .  .  .  .  .  .  .  .  Name: "Animal"
   404  .  .  .  .  .  .  .  .  .  .  Obj: *(obj @ 47)
   405  .  .  .  .  .  .  .  .  .  }
   406  .  .  .  .  .  .  .  .  .  Tag: nil
   407  .  .  .  .  .  .  .  .  .  Comment: nil
   408  .  .  .  .  .  .  .  .  }
