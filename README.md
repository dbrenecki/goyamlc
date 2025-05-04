# goyamlc

goyaml marshaler, but with comments!

goyamlc uses the Abstract Syntax Tree (AST) to maintain comments from structs to yaml.

This is useful for generating docs or examples in yaml whilst maintaining a single location of comments.


## How To use

```bash
go install github.com/dbrenecki/goyamlc@latest
```

Define your structs in one file which is your `--types-path`

```go
type Foo struct {
    // some field
    x string `yaml:"x"`
    // some other field
    y string `yaml:"y"`
}
```

Set the location to generate the yaml and run:
```bash
goyamlc  --types-path=<your-types-path> --gen-path=<your-yaml-path>
```

Which should generate e.g.

```yaml
foo:
  # some field
  x: string
  # some other field
  y: string
```

See [example](./test/data/generated.yaml) for an example generation.


## Upcoming Features

- Make gen order deterministic (bug)
- Ignore unexported fields
- Improve how map types are represented
