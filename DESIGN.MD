# Design

Steps :
1. Validate JsonLogic input  
   We can use the jsonlogic package, it provides a `jsonlogic.IsValid()` function.
2. Convert it to a MongoDB aggregate query  
   Like jsonlogic package do, we can convert it to an `interface{}` and analyze each keyword thanks to the `reflect` package.
3. Validate MongoDB query  
   To validate the query, we can use the `mongodb.MarshalBSON()` function from the official package.

| JsonLogic Keyword | MongoDB aggregate equivalent |
| ----------------- | ---------------------------- |
| ==                | $match                       |
| !=                | $match + $not & $eq          |
| !                 | $not                         |
| or                | $or                          |
| and               | $and                         |
| filter            | $filter                      |
| var               | $match ?                     |

The library must support custom jsonlogic operator.

Example: 
```json
{
  "or": [
    {
      "is_key_value": [
        {
          "var": ".metadata.labels"
        },
        "app.kubernetes.io/component",
        "agent"
      ]
    },
    {
      "!": {
        "is_key_value": [
          {
            "var": ".metadata.labels"
          },
          "app.kubernetes.io/name",
          "kubees"
        ]
      }
    }
  ]
}
```

**The package must expose :**
- `Convert(jsonlogic interface{})`
- `AddOperator(jsonlogicKeyword string, func mongodbResult(data interface{}) bson.D{})` ? Not sure about this format, need to start development to see what happens.

## Tests

Multiple units tests must be written.