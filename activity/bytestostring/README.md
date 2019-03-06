
# Bytes to String - Activity

## Installation

```bash
flogo install github.com/ankaranika/flogo-extensions/activity/bytestostring
```
Link for flogo web:
```bash
https://github.com/ankaranika/flogo-extensions/activity/bytestostring
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "bytes",
      "type": "array"
    }
  ],
  "outputs": [
    {
      "name": "str",
      "type": "string"
    }
  ]
}
```
## Inputs
| Input   | Description    |
|:----------|:---------------|
| bytes | A byte array |

## Outputs
| Output   | Description    |
|:----------|:---------------|
| str | The stringified byte array |
