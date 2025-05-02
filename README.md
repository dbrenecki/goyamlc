# goyamlc

WIP

## Summary

Marshal go structs to yaml, but with comments!

`goyamlc` offers maintaining comments from your structs as part of your yaml.

This offers a single place to maintain comments whilst generating examples or documentation in yaml.

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
