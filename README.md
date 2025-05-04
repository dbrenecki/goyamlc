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

See the [example](./test/data/generated.yaml) for a more complete example which includes nested types.


## Upcoming Features

- Make gen order deterministic
- Ignore unexported fields
- Improve how map types are represented
