# goyamlc

go yaml marshaler with comments (WIP).

Example

```go
type Foo struct {
    // some field
    x string `yaml:"x"`
    // some other field
    y string `yaml:"y"`
}
```

Would generate:

```yaml
foo:
  # some field
  x: string
  # some other field
  y: string
```

More complex examples would include nested structs and multiple root structs.
