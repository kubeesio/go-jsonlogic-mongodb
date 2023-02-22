# go-jsonlogic-mongodb

Convert JsonLogic into a MongoDB aggregate query.

## Design

Steps :
1. Validate JsonLogic input
2. Convert it to a MongoDB aggregate query
3. Validate MongoDB query

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

## Tests

Multiple units tests must be written.