
# espeak - Activity

## Installation

```bash
flogo install github.com/ankaranika/flogo-extensions/activity/espeak
```
Link for flogo web:
```bash
https://github.com/ankaranika/flogo-extensions/activity/espeak
```

## Schema
Inputs and Outputs:

```json
{
"inputs":[
    {
      "name": "ip",
      "type": "string"
    },
    {
      "name": "req_id",
      "type": "string"
    }
  ]
}
```
## Inputs
| Input   | Description    |
|:----------|:---------------|
| ip | The ip of the request sender |
| req_id | The request id from the specified sender |