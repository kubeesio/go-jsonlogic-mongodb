# go-jsonlogic-mongodb

[![test](https://github.com/kubeesio/go-jsonlogic-mongodb/actions/workflows/test.yaml/badge.svg)](https://github.com/kubeesio/go-jsonlogic-mongodb/actions/workflows/test.yaml)
[![validate](https://github.com/kubeesio/go-jsonlogic-mongodb/actions/workflows/validate.yaml/badge.svg)](https://github.com/kubeesio/go-jsonlogic-mongodb/actions/workflows/validate.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kubeesio/go-jsonlogic-mongodb)](https://goreportcard.com/report/github.com/kubeesio/go-jsonlogic-mongodb)
[![Go Reference](https://pkg.go.dev/badge/github.com/kubeesio/go-jsonlogic-mongodb.svg)](https://pkg.go.dev/github.com/kubeesio/go-jsonlogic-mongodb)


Convert JsonLogic into a MongoDB aggregate query.

## How to use

Use the `Convert()` function with your jsonLogic as parameter and get your mongo pipeline generated.

```go
func Convert(rules io.Reader) (bson.D, error)
```

## Examples

### Equal operator

<ins>Json Logic input</ins>

```go
{"==": [1, 1]}
```

<ins>Mongo output</ins>

```go
bson.D{{
  Key: "$eq",
  Value: bson.A{
    Key: "1",
    Value: 1.0
  }
}}
```

### Not Equal operator

<ins>Json Logic input</ins>

```go
{"!=": ["Hello", "Bonjour"]}
```

<ins>Mongo output</ins>

```go
bson.D{{
  Key: "$ne", 
  Value: bson.A{"Hello", "Bonjour"}
}}
```

### And & Or operators

And & Or operators work the same way, just specify what operator is needed it will result in a mongo `$and` or `$or`.

<ins>Json Logic input</ins>

```go
{
  "and": [
    {
      "==": [
        1,
        1
      ]
    },
    {
      "!=": [
        "Hello",
        "Bonjour"
      ]
    }
  ]
}
```

<ins>Mongo output</ins>

```go
bson.D{{
  Key: "$and",
  Value: bson.A{
    bson.D{{
      Key: "$eq",
      Value: bson.A{
        1.0,
        1.0
      }
    }}, 
    bson.D{{
      Key: "$ne",
      Value: bson.A{
        "Hello", 
        "Bonjour"
      }
    }}
  }
}}
```
