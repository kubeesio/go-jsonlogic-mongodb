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

### Not operator

<ins>Json Logic input</ins>

```go
{"!": "true"}
```

<ins>Mongo output</ins>

```go
bson.D{{
  Key: "$not", 
  Value: "true"
}}
```

### Var operator

<ins>Json Logic input</ins>

```go
{"==": ["kube-system", {"var": ".metadata.namespace"}]}
```

<ins>Mongo output</ins>

```go
bson.D{{
  Key: "$eq",
  Value: bson.A{
    "kube-system",
    "$metadata.namespace"
  }
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

### Custom operator

In order to add a custom operator, you need to use the `AddOperator` function.  
To make it work, it is mandatory to add your custom operator on the `jsonlogic` library side too, otherwise, `go-jsonlogic-mongodb` will not  validate it.

Example :

Create your custom function
```go
func isKeyValue(value interface{}) (primitive.D, error) {
	parsed, _ := value.([]interface{})

	key := parsed[1].(string)
	val := parsed[2].(string)

	firstArgument, internalError := InternalConvert(value.([]interface{})[0])
	if internalError != nil {
		return nil, internalError
	}

	return bson.D{{
		Key: "$match",
		Value: bson.D{{
			Key: "$expr", Value: bson.D{{
				Key: "$eq", Value: bson.A{
					bson.D{{
						Key: "$getField", Value: bson.D{
							{Key: "field", Value: bson.D{{Key: "$literal", Value: key}}},
							{Key: "input", Value: firstArgument},
						},
					}},
					val,
				},
			}},
		}},
	}}, nil
}
```

Add go-jsonlogic-mongodb custom operator
```go
AddOperator("is_key_value", isKeyValue)
```

Override the `jsonlogic` operator, if you are not interested in applying the jsonlogic with your custom operator in the future, you can override it with an empty function like following :
```go
	jsonlogic.AddOperator(name, func(values interface{}, data interface{}) (result interface{}) { return })

```

This way, you can achieve this :

```go
{
  "filter": [
    {
      "var": ".resources"
    },
    {
      "==": [
      {
        "var": ".metadata.namespace"
      },
      1
    ]}
  ]
}
```

<ins>Mongo output</ins>

```go
bson.D{{
  Key: "$addFields", Value: bson.D{{
    Key: "resources", Value: bson.D{{
      Key: "$filter", Value: bson.D{
        {
          Key: "input", Value: "$resources"
        },
        {
          Key: "cond", Value: bson.D{{
            Key: "$eq", Value: bson.A{
              "$$this.metadata.namespace",
              1.0
            }
          }}
        }
      }
    }}
  }}
}}
```